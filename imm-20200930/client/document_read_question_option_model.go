// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocumentReadQuestionOption interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DocumentReadQuestionOption
	GetCount() *int32
	SetExtract(v bool) *DocumentReadQuestionOption
	GetExtract() *bool
}

type DocumentReadQuestionOption struct {
	Count   *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	Extract *bool  `json:"Extract,omitempty" xml:"Extract,omitempty"`
}

func (s DocumentReadQuestionOption) String() string {
	return dara.Prettify(s)
}

func (s DocumentReadQuestionOption) GoString() string {
	return s.String()
}

func (s *DocumentReadQuestionOption) GetCount() *int32 {
	return s.Count
}

func (s *DocumentReadQuestionOption) GetExtract() *bool {
	return s.Extract
}

func (s *DocumentReadQuestionOption) SetCount(v int32) *DocumentReadQuestionOption {
	s.Count = &v
	return s
}

func (s *DocumentReadQuestionOption) SetExtract(v bool) *DocumentReadQuestionOption {
	s.Extract = &v
	return s
}

func (s *DocumentReadQuestionOption) Validate() error {
	return dara.Validate(s)
}
