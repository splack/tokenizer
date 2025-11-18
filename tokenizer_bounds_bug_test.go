package tokenizer_test

import (
	"testing"

	"github.com/sugarme/tokenizer/pretrained"
)

func TestSliceBoundsOutOfRange(t *testing.T) {
	// This test demonstrates a slice bounds panic in normalizer/normalized.go:776
	// The bug occurs when TransformRange tries to access alignmentsOriginal[endRange[1]]
	// where endRange[1] exceeds the array length
	
	// Create a synthetic text that triggers the bug
	// The pattern involves nested braces with specific spacing/indentation
	text := `config = option {
            type = types.module {
              fields = {
                items = option {
                  type = types.list types.text;
                  default = [ "1.1.1.1" "8.8.8.8" ];
                  description = "List of addresses for checks";
                };
              };
            };
            default = {};
          };`

	// Load multilingual-e5-small tokenizer (or any tokenizer with Metaspace pretokenizer)
	tk, err := pretrained.FromFile("testdata/tokenizer.json")
	if err != nil {
		t.Skip("Tokenizer not available, skipping")
	}

	// This should not panic
	encoding, err := tk.EncodeSingle(text, true)
	if err != nil {
		t.Fatalf("EncodeSingle failed: %v", err)
	}

	t.Logf("Successfully encoded %d tokens", len(encoding.Ids))
}
