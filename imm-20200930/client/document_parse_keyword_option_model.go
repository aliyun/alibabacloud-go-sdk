// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocumentParseKeywordOption interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DocumentParseKeywordOption
	GetCount() *int32
	SetExtract(v bool) *DocumentParseKeywordOption
	GetExtract() *bool
}

type DocumentParseKeywordOption struct {
	// example:
	//
	// 1
	Count   *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	Extract *bool  `json:"Extract,omitempty" xml:"Extract,omitempty"`
}

func (s DocumentParseKeywordOption) String() string {
	return dara.Prettify(s)
}

func (s DocumentParseKeywordOption) GoString() string {
	return s.String()
}

func (s *DocumentParseKeywordOption) GetCount() *int32 {
	return s.Count
}

func (s *DocumentParseKeywordOption) GetExtract() *bool {
	return s.Extract
}

func (s *DocumentParseKeywordOption) SetCount(v int32) *DocumentParseKeywordOption {
	s.Count = &v
	return s
}

func (s *DocumentParseKeywordOption) SetExtract(v bool) *DocumentParseKeywordOption {
	s.Extract = &v
	return s
}

func (s *DocumentParseKeywordOption) Validate() error {
	return dara.Validate(s)
}
