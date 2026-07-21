// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetricKVPairDTO interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *MetricKVPairDTO
	GetKey() *string
	SetValue(v float32) *MetricKVPairDTO
	GetValue() *float32
}

type MetricKVPairDTO struct {
	// Metric Name.
	//
	// **Chat*	-
	//
	// - `total_calls`: Number Of Calls, integer, Count
	//
	// - `input_tokens`: Total input tokens, integer
	//
	// - `output_tokens`: Total output tokens, integer
	//
	// - `reasoning_tokens`: Reasoning tokens, integer
	//
	// - `cached_tokens`: Cached input tokens (hit), integer
	//
	// **Vision*	-
	//
	// - `total_calls`: Number Of Calls, integer, Count
	//
	// - `image_count`: Number of generated images, integer
	//
	// - `video_duration`: Generated video duration, rounded to 3 decimal places, seconds
	//
	// **Embedding*	-
	//
	// - `total_calls`: Number Of Calls, integer, Count
	//
	// - `embedding_output_tokens`: Embedding output tokens, integer
	//
	// - `billing_tokens`: Total billing tokens, integer
	//
	// - `image_tokens`: Image tokens (multimodal embedding), integer
	//
	// **Omni-modal (ChatFullmodal / ChatMultimodal)*	-
	//
	// - `total_calls`: Number Of Calls, integer, Count
	//
	// - `input_text_tokens`: Input text tokens, integer
	//
	// - `input_audio_tokens`: Input audio tokens, integer
	//
	// - `input_image_tokens`: Input image tokens, integer
	//
	// - `input_video_tokens`: Input video tokens, integer
	//
	// - `output_text_tokens`: Output text tokens, integer
	//
	// - `output_audio_tokens`: Output audio tokens, integer
	//
	// **Speech (TTS / ASR)*	-
	//
	// - `total_calls`: Number Of Calls, integer, Count
	//
	// - `characters`: Characters converted to speech, integer
	//
	// - `asr_duration`: Speech recognition duration, rounded to 3 decimal places, seconds
	//
	// example:
	//
	// total_calls
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// Metric value
	//
	// example:
	//
	// 100
	Value *float32 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s MetricKVPairDTO) String() string {
	return dara.Prettify(s)
}

func (s MetricKVPairDTO) GoString() string {
	return s.String()
}

func (s *MetricKVPairDTO) GetKey() *string {
	return s.Key
}

func (s *MetricKVPairDTO) GetValue() *float32 {
	return s.Value
}

func (s *MetricKVPairDTO) SetKey(v string) *MetricKVPairDTO {
	s.Key = &v
	return s
}

func (s *MetricKVPairDTO) SetValue(v float32) *MetricKVPairDTO {
	s.Value = &v
	return s
}

func (s *MetricKVPairDTO) Validate() error {
	return dara.Validate(s)
}
