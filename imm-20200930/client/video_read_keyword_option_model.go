// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoReadKeywordOption interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *VideoReadKeywordOption
	GetCount() *int32
	SetExtract(v bool) *VideoReadKeywordOption
	GetExtract() *bool
}

type VideoReadKeywordOption struct {
	// The number of keywords. Valid values: 0 to 10.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Specifies whether to fetch keywords.
	//
	// example:
	//
	// true
	Extract *bool `json:"Extract,omitempty" xml:"Extract,omitempty"`
}

func (s VideoReadKeywordOption) String() string {
	return dara.Prettify(s)
}

func (s VideoReadKeywordOption) GoString() string {
	return s.String()
}

func (s *VideoReadKeywordOption) GetCount() *int32 {
	return s.Count
}

func (s *VideoReadKeywordOption) GetExtract() *bool {
	return s.Extract
}

func (s *VideoReadKeywordOption) SetCount(v int32) *VideoReadKeywordOption {
	s.Count = &v
	return s
}

func (s *VideoReadKeywordOption) SetExtract(v bool) *VideoReadKeywordOption {
	s.Extract = &v
	return s
}

func (s *VideoReadKeywordOption) Validate() error {
	return dara.Validate(s)
}
