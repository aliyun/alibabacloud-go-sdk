// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocumentParseQuestionOption interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DocumentParseQuestionOption
	GetCount() *int32
	SetExtract(v bool) *DocumentParseQuestionOption
	GetExtract() *bool
}

type DocumentParseQuestionOption struct {
	// example:
	//
	// 1
	Count   *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	Extract *bool  `json:"Extract,omitempty" xml:"Extract,omitempty"`
}

func (s DocumentParseQuestionOption) String() string {
	return dara.Prettify(s)
}

func (s DocumentParseQuestionOption) GoString() string {
	return s.String()
}

func (s *DocumentParseQuestionOption) GetCount() *int32 {
	return s.Count
}

func (s *DocumentParseQuestionOption) GetExtract() *bool {
	return s.Extract
}

func (s *DocumentParseQuestionOption) SetCount(v int32) *DocumentParseQuestionOption {
	s.Count = &v
	return s
}

func (s *DocumentParseQuestionOption) SetExtract(v bool) *DocumentParseQuestionOption {
	s.Extract = &v
	return s
}

func (s *DocumentParseQuestionOption) Validate() error {
	return dara.Validate(s)
}
