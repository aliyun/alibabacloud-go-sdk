// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyLastUsedResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKey(v string) *GetAccessKeyLastUsedResourcesRequest
	GetAccessKey() *string
	SetNextToken(v string) *GetAccessKeyLastUsedResourcesRequest
	GetNextToken() *string
	SetPageSize(v string) *GetAccessKeyLastUsedResourcesRequest
	GetPageSize() *string
	SetServiceName(v string) *GetAccessKeyLastUsedResourcesRequest
	GetServiceName() *string
}

type GetAccessKeyLastUsedResourcesRequest struct {
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
	// > The request parameters must be the same as those of the last request.
	//
	// example:
	//
	// eyJhY2NvdW50IjoiMTQyNDM3OTU4NjM4NzE2MSIsImV2ZW50SWQiOiI3MkJDRTExRi02OTU3LTQ0NUItQjY0MC1CNEUyMkM4NUEwQzgiLCJsb2dJZCI6IjgyLTE0MjQzNzk1ODYzODcxNjEiLCJ0aW1lIjoxNjAyMzExNTQwMD****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries per page.
	//
	// 	- Valid values: 0 to 100.
	//
	// 	- Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The Alibaba Cloud service. For more information about the Alibaba Cloud services supported by ActionTrail, see [Supported Alibaba Cloud services](https://help.aliyun.com/document_detail/28829.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Ecs
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GetAccessKeyLastUsedResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyLastUsedResourcesRequest) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedResourcesRequest) GetAccessKey() *string {
	return s.AccessKey
}

func (s *GetAccessKeyLastUsedResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetAccessKeyLastUsedResourcesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *GetAccessKeyLastUsedResourcesRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetAccessKeyLastUsedResourcesRequest) SetAccessKey(v string) *GetAccessKeyLastUsedResourcesRequest {
	s.AccessKey = &v
	return s
}

func (s *GetAccessKeyLastUsedResourcesRequest) SetNextToken(v string) *GetAccessKeyLastUsedResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *GetAccessKeyLastUsedResourcesRequest) SetPageSize(v string) *GetAccessKeyLastUsedResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *GetAccessKeyLastUsedResourcesRequest) SetServiceName(v string) *GetAccessKeyLastUsedResourcesRequest {
	s.ServiceName = &v
	return s
}

func (s *GetAccessKeyLastUsedResourcesRequest) Validate() error {
	return dara.Validate(s)
}
