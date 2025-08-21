// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoCcBlacklistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeAutoCcBlacklistRequest
	GetInstanceId() *string
	SetKeyWord(v string) *DescribeAutoCcBlacklistRequest
	GetKeyWord() *string
	SetPageNumber(v int32) *DescribeAutoCcBlacklistRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAutoCcBlacklistRequest
	GetPageSize() *int32
	SetQueryType(v string) *DescribeAutoCcBlacklistRequest
	GetQueryType() *string
}

type DescribeAutoCcBlacklistRequest struct {
	// The ID of the instance.
	//
	// > You can call the [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) operation to query the IDs of all instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The keyword for the query. This keyword is used to specify the prefix of the source IP address that you want to query.
	//
	// > The keyword must be greater than three characters in length.
	//
	// example:
	//
	// 138
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The number of the page to return. For example, to query the returned results on the first page, set the value to **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
}

func (s DescribeAutoCcBlacklistRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoCcBlacklistRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoCcBlacklistRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAutoCcBlacklistRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *DescribeAutoCcBlacklistRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAutoCcBlacklistRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAutoCcBlacklistRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *DescribeAutoCcBlacklistRequest) SetInstanceId(v string) *DescribeAutoCcBlacklistRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAutoCcBlacklistRequest) SetKeyWord(v string) *DescribeAutoCcBlacklistRequest {
	s.KeyWord = &v
	return s
}

func (s *DescribeAutoCcBlacklistRequest) SetPageNumber(v int32) *DescribeAutoCcBlacklistRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoCcBlacklistRequest) SetPageSize(v int32) *DescribeAutoCcBlacklistRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoCcBlacklistRequest) SetQueryType(v string) *DescribeAutoCcBlacklistRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeAutoCcBlacklistRequest) Validate() error {
	return dara.Validate(s)
}
