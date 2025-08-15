// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyLastUsedEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKey(v string) *GetAccessKeyLastUsedEventsRequest
	GetAccessKey() *string
	SetNextToken(v string) *GetAccessKeyLastUsedEventsRequest
	GetNextToken() *string
	SetPageSize(v string) *GetAccessKeyLastUsedEventsRequest
	GetPageSize() *string
	SetServiceName(v string) *GetAccessKeyLastUsedEventsRequest
	GetServiceName() *string
}

type GetAccessKeyLastUsedEventsRequest struct {
	// The AccessKey ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// LTAI****************
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The token that determines the start point of the query.
	//
	// > The request parameters must be the same as those of the last request.
	//
	// example:
	//
	// eyJhY2NvdW50IjoiMTQyNDM3OTU4NjM4NzE2MSIsImV2ZW50SWQiOiI3MkJDRTExRi02OTU3LTQ0NUItQjY0MC1CNEUyMkM4NUEwQzgiLCJsb2dJZCI6IjgyLTE0MjQzNzk1ODYzODcxNjEiLCJ0aW1lIjoxNjAyMzExNTQwMD****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 0 to 100.
	//
	// Default value: 20.
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

func (s GetAccessKeyLastUsedEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyLastUsedEventsRequest) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedEventsRequest) GetAccessKey() *string {
	return s.AccessKey
}

func (s *GetAccessKeyLastUsedEventsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetAccessKeyLastUsedEventsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *GetAccessKeyLastUsedEventsRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetAccessKeyLastUsedEventsRequest) SetAccessKey(v string) *GetAccessKeyLastUsedEventsRequest {
	s.AccessKey = &v
	return s
}

func (s *GetAccessKeyLastUsedEventsRequest) SetNextToken(v string) *GetAccessKeyLastUsedEventsRequest {
	s.NextToken = &v
	return s
}

func (s *GetAccessKeyLastUsedEventsRequest) SetPageSize(v string) *GetAccessKeyLastUsedEventsRequest {
	s.PageSize = &v
	return s
}

func (s *GetAccessKeyLastUsedEventsRequest) SetServiceName(v string) *GetAccessKeyLastUsedEventsRequest {
	s.ServiceName = &v
	return s
}

func (s *GetAccessKeyLastUsedEventsRequest) Validate() error {
	return dara.Validate(s)
}
