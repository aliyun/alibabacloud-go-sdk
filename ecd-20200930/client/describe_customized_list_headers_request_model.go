// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizedListHeadersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLangType(v string) *DescribeCustomizedListHeadersRequest
	GetLangType() *string
	SetListType(v string) *DescribeCustomizedListHeadersRequest
	GetListType() *string
	SetRegionId(v string) *DescribeCustomizedListHeadersRequest
	GetRegionId() *string
}

type DescribeCustomizedListHeadersRequest struct {
	// example:
	//
	// zh-CN
	LangType *string `json:"LangType,omitempty" xml:"LangType,omitempty"`
	// example:
	//
	// desktop
	ListType *string `json:"ListType,omitempty" xml:"ListType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCustomizedListHeadersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizedListHeadersRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomizedListHeadersRequest) GetLangType() *string {
	return s.LangType
}

func (s *DescribeCustomizedListHeadersRequest) GetListType() *string {
	return s.ListType
}

func (s *DescribeCustomizedListHeadersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCustomizedListHeadersRequest) SetLangType(v string) *DescribeCustomizedListHeadersRequest {
	s.LangType = &v
	return s
}

func (s *DescribeCustomizedListHeadersRequest) SetListType(v string) *DescribeCustomizedListHeadersRequest {
	s.ListType = &v
	return s
}

func (s *DescribeCustomizedListHeadersRequest) SetRegionId(v string) *DescribeCustomizedListHeadersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCustomizedListHeadersRequest) Validate() error {
	return dara.Validate(s)
}
