// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyLastUsedIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKey(v string) *GetAccessKeyLastUsedIpsRequest
	GetAccessKey() *string
	SetNextToken(v string) *GetAccessKeyLastUsedIpsRequest
	GetNextToken() *string
	SetPageSize(v string) *GetAccessKeyLastUsedIpsRequest
	GetPageSize() *string
	SetServiceName(v string) *GetAccessKeyLastUsedIpsRequest
	GetServiceName() *string
}

type GetAccessKeyLastUsedIpsRequest struct {
	// The AccessKey ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// LTAI****************
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// >  You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// eyJhY2NvdW50IjoiMTQyNDM3OTU4NjM4NzE2MSIsImV2ZW50SWQiOiI3MkJDRTExRi02OTU3LTQ0NUItQjY0MC1CNEUyMkM4NUEwQzgiLCJsb2dJZCI6IjgyLTE0MjQzNzk1ODYzODcxNjEiLCJ0aW1lIjoxNjAyMzExNTQwMD****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries per page. Valid values: 0 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The Alibaba Cloud service. For more information about the Alibaba Cloud services supported by ActionTrail, see [Services that work with ActionTrail](https://help.aliyun.com/document_detail/28829.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Ecs
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GetAccessKeyLastUsedIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyLastUsedIpsRequest) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedIpsRequest) GetAccessKey() *string {
	return s.AccessKey
}

func (s *GetAccessKeyLastUsedIpsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetAccessKeyLastUsedIpsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *GetAccessKeyLastUsedIpsRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetAccessKeyLastUsedIpsRequest) SetAccessKey(v string) *GetAccessKeyLastUsedIpsRequest {
	s.AccessKey = &v
	return s
}

func (s *GetAccessKeyLastUsedIpsRequest) SetNextToken(v string) *GetAccessKeyLastUsedIpsRequest {
	s.NextToken = &v
	return s
}

func (s *GetAccessKeyLastUsedIpsRequest) SetPageSize(v string) *GetAccessKeyLastUsedIpsRequest {
	s.PageSize = &v
	return s
}

func (s *GetAccessKeyLastUsedIpsRequest) SetServiceName(v string) *GetAccessKeyLastUsedIpsRequest {
	s.ServiceName = &v
	return s
}

func (s *GetAccessKeyLastUsedIpsRequest) Validate() error {
	return dara.Validate(s)
}
