// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocumentReadKeywordOption interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DocumentReadKeywordOption
	GetCount() *int32
	SetExtract(v bool) *DocumentReadKeywordOption
	GetExtract() *bool
}

type DocumentReadKeywordOption struct {
	// The maximum number of keywords to return.
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Whether to extract keywords from the document.
	Extract *bool `json:"Extract,omitempty" xml:"Extract,omitempty"`
}

func (s DocumentReadKeywordOption) String() string {
	return dara.Prettify(s)
}

func (s DocumentReadKeywordOption) GoString() string {
	return s.String()
}

func (s *DocumentReadKeywordOption) GetCount() *int32 {
	return s.Count
}

func (s *DocumentReadKeywordOption) GetExtract() *bool {
	return s.Extract
}

func (s *DocumentReadKeywordOption) SetCount(v int32) *DocumentReadKeywordOption {
	s.Count = &v
	return s
}

func (s *DocumentReadKeywordOption) SetExtract(v bool) *DocumentReadKeywordOption {
	s.Extract = &v
	return s
}

func (s *DocumentReadKeywordOption) Validate() error {
	return dara.Validate(s)
}
