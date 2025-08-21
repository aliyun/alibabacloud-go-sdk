// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoCcWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeAutoCcWhitelistRequest
	GetInstanceId() *string
	SetKeyWord(v string) *DescribeAutoCcWhitelistRequest
	GetKeyWord() *string
	SetPageNumber(v int32) *DescribeAutoCcWhitelistRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAutoCcWhitelistRequest
	GetPageSize() *int32
}

type DescribeAutoCcWhitelistRequest struct {
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
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAutoCcWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoCcWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoCcWhitelistRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAutoCcWhitelistRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *DescribeAutoCcWhitelistRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAutoCcWhitelistRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAutoCcWhitelistRequest) SetInstanceId(v string) *DescribeAutoCcWhitelistRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAutoCcWhitelistRequest) SetKeyWord(v string) *DescribeAutoCcWhitelistRequest {
	s.KeyWord = &v
	return s
}

func (s *DescribeAutoCcWhitelistRequest) SetPageNumber(v int32) *DescribeAutoCcWhitelistRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoCcWhitelistRequest) SetPageSize(v int32) *DescribeAutoCcWhitelistRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoCcWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
