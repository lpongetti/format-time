package formattime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormat_DifferentDateFormats(t *testing.T) {
	input := "2022-06-08T09:04:05.832373232Z"
	parsedTime, err := time.Parse(time.RFC3339Nano, input)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	testCases := []struct {
		mask     string
		expected string
	}{
		{
			mask:     "YYYY-MM-DD HH:mm:ss:SSS",
			expected: "2022-06-08 09:04:05:000",
		},
		{
			mask:     "YY-M-D HH:m:s",
			expected: "22-6-8 09:4:5",
		},
		{
			mask:     "YYYY-MM-DD HH:mm:ss",
			expected: "2022-06-08 09:04:05",
		},
		{
			mask:     "YYYY-MMM-DD hh:mm:ss AA",
			expected: "2022-Jun-08 09:04:05 AM",
		},
		{
			mask:     "YYYY-MMMM-DD hh:mm:ss aa",
			expected: "2022-June-08 09:04:05 am",
		},
		{
			mask:     "YYYY-MM-DD Z",
			expected: "2022-06-08 +00:00",
		},
		{
			mask:     "DD/MM/YYYY ZZ",
			expected: "08/06/2022 +0000",
		},
		{
			mask:     "MMMM DD, YYYY",
			expected: "June 08, 2022",
		},
		{
			mask:     "MMM D, YYYY",
			expected: "Jun 8, 2022",
		},
	}

	for _, tc := range testCases {
		layout := tc.mask
		expected := tc.expected

		result := Format(layout, parsedTime)
		if result != expected {
			t.Errorf("Expected: %s, got: %s", expected, result)
		}
	}
}
