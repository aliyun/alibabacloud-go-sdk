// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeChannelsRequest
	GetAppId() *string
	SetPageNo(v int32) *DescribeChannelsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeChannelsRequest
	GetPageSize() *int32
}

type DescribeChannelsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeChannelsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelsRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeChannelsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeChannelsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeChannelsRequest) SetAppId(v string) *DescribeChannelsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelsRequest) SetPageNo(v int32) *DescribeChannelsRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeChannelsRequest) SetPageSize(v int32) *DescribeChannelsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeChannelsRequest) Validate() error {
	return dara.Validate(s)
}
