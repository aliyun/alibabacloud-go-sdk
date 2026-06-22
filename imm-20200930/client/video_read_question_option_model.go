// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoReadQuestionOption interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *VideoReadQuestionOption
	GetCount() *int32
	SetExtract(v bool) *VideoReadQuestionOption
	GetExtract() *bool
}

type VideoReadQuestionOption struct {
	// Specifies the maximum number of answers to return when `Extract` is `true`. If omitted, the service returns all detected answers.
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Specifies whether to extract answer segments from the video. When `true`, the service identifies and returns these segments. The default is `false`.
	Extract *bool `json:"Extract,omitempty" xml:"Extract,omitempty"`
}

func (s VideoReadQuestionOption) String() string {
	return dara.Prettify(s)
}

func (s VideoReadQuestionOption) GoString() string {
	return s.String()
}

func (s *VideoReadQuestionOption) GetCount() *int32 {
	return s.Count
}

func (s *VideoReadQuestionOption) GetExtract() *bool {
	return s.Extract
}

func (s *VideoReadQuestionOption) SetCount(v int32) *VideoReadQuestionOption {
	s.Count = &v
	return s
}

func (s *VideoReadQuestionOption) SetExtract(v bool) *VideoReadQuestionOption {
	s.Extract = &v
	return s
}

func (s *VideoReadQuestionOption) Validate() error {
	return dara.Validate(s)
}
