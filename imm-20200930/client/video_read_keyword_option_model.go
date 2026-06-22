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
	// Specifies the maximum number of keywords to return. The service may return fewer keywords than this limit. If omitted, the service returns all extracted keywords.
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Specifies whether to extract keywords from the video. Set to `true` to enable keyword extraction. The default is `false`.
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
