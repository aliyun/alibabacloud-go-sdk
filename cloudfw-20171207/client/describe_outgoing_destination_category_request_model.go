// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingDestinationCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *DescribeOutgoingDestinationCategoryRequest
	GetCategoryId() *string
	SetDstType(v string) *DescribeOutgoingDestinationCategoryRequest
	GetDstType() *string
	SetEndTime(v string) *DescribeOutgoingDestinationCategoryRequest
	GetEndTime() *string
	SetLang(v string) *DescribeOutgoingDestinationCategoryRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeOutgoingDestinationCategoryRequest
	GetSourceIp() *string
	SetStartTime(v string) *DescribeOutgoingDestinationCategoryRequest
	GetStartTime() *string
	SetTypeId(v string) *DescribeOutgoingDestinationCategoryRequest
	GetTypeId() *string
}

type DescribeOutgoingDestinationCategoryRequest struct {
	// example:
	//
	// All
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// domain
	DstType *string `json:"DstType,omitempty" xml:"DstType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1750818370
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 59.82.45.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1749657600
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// All
	TypeId *string `json:"TypeId,omitempty" xml:"TypeId,omitempty"`
}

func (s DescribeOutgoingDestinationCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDestinationCategoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationCategoryRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *DescribeOutgoingDestinationCategoryRequest) GetDstType() *string {
	return s.DstType
}

func (s *DescribeOutgoingDestinationCategoryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeOutgoingDestinationCategoryRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeOutgoingDestinationCategoryRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeOutgoingDestinationCategoryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeOutgoingDestinationCategoryRequest) GetTypeId() *string {
	return s.TypeId
}

func (s *DescribeOutgoingDestinationCategoryRequest) SetCategoryId(v string) *DescribeOutgoingDestinationCategoryRequest {
	s.CategoryId = &v
	return s
}

func (s *DescribeOutgoingDestinationCategoryRequest) SetDstType(v string) *DescribeOutgoingDestinationCategoryRequest {
	s.DstType = &v
	return s
}

func (s *DescribeOutgoingDestinationCategoryRequest) SetEndTime(v string) *DescribeOutgoingDestinationCategoryRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOutgoingDestinationCategoryRequest) SetLang(v string) *DescribeOutgoingDestinationCategoryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOutgoingDestinationCategoryRequest) SetSourceIp(v string) *DescribeOutgoingDestinationCategoryRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOutgoingDestinationCategoryRequest) SetStartTime(v string) *DescribeOutgoingDestinationCategoryRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOutgoingDestinationCategoryRequest) SetTypeId(v string) *DescribeOutgoingDestinationCategoryRequest {
	s.TypeId = &v
	return s
}

func (s *DescribeOutgoingDestinationCategoryRequest) Validate() error {
	return dara.Validate(s)
}
