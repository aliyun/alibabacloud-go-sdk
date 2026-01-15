// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchImageByTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDistinctProductId(v bool) *SearchImageByTextRequest
	GetDistinctProductId() *bool
	SetFilter(v string) *SearchImageByTextRequest
	GetFilter() *string
	SetInstanceName(v string) *SearchImageByTextRequest
	GetInstanceName() *string
	SetNum(v int32) *SearchImageByTextRequest
	GetNum() *int32
	SetScoreThreshold(v string) *SearchImageByTextRequest
	GetScoreThreshold() *string
	SetStart(v int32) *SearchImageByTextRequest
	GetStart() *int32
	SetText(v string) *SearchImageByTextRequest
	GetText() *string
}

type SearchImageByTextRequest struct {
	// example:
	//
	// false
	DistinctProductId *bool `json:"DistinctProductId,omitempty" xml:"DistinctProductId,omitempty"`
	// example:
	//
	// int_attr=1000 AND str_attr="value1"
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// demoinstance1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 10
	Num            *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	ScoreThreshold *string `json:"ScoreThreshold,omitempty" xml:"ScoreThreshold,omitempty"`
	// example:
	//
	// 0
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s SearchImageByTextRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByTextRequest) GoString() string {
	return s.String()
}

func (s *SearchImageByTextRequest) GetDistinctProductId() *bool {
	return s.DistinctProductId
}

func (s *SearchImageByTextRequest) GetFilter() *string {
	return s.Filter
}

func (s *SearchImageByTextRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *SearchImageByTextRequest) GetNum() *int32 {
	return s.Num
}

func (s *SearchImageByTextRequest) GetScoreThreshold() *string {
	return s.ScoreThreshold
}

func (s *SearchImageByTextRequest) GetStart() *int32 {
	return s.Start
}

func (s *SearchImageByTextRequest) GetText() *string {
	return s.Text
}

func (s *SearchImageByTextRequest) SetDistinctProductId(v bool) *SearchImageByTextRequest {
	s.DistinctProductId = &v
	return s
}

func (s *SearchImageByTextRequest) SetFilter(v string) *SearchImageByTextRequest {
	s.Filter = &v
	return s
}

func (s *SearchImageByTextRequest) SetInstanceName(v string) *SearchImageByTextRequest {
	s.InstanceName = &v
	return s
}

func (s *SearchImageByTextRequest) SetNum(v int32) *SearchImageByTextRequest {
	s.Num = &v
	return s
}

func (s *SearchImageByTextRequest) SetScoreThreshold(v string) *SearchImageByTextRequest {
	s.ScoreThreshold = &v
	return s
}

func (s *SearchImageByTextRequest) SetStart(v int32) *SearchImageByTextRequest {
	s.Start = &v
	return s
}

func (s *SearchImageByTextRequest) SetText(v string) *SearchImageByTextRequest {
	s.Text = &v
	return s
}

func (s *SearchImageByTextRequest) Validate() error {
	return dara.Validate(s)
}
