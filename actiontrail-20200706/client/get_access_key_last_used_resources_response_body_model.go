// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyLastUsedResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *GetAccessKeyLastUsedResourcesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *GetAccessKeyLastUsedResourcesResponseBody
	GetRequestId() *string
	SetResources(v []*GetAccessKeyLastUsedResourcesResponseBodyResources) *GetAccessKeyLastUsedResourcesResponseBody
	GetResources() []*GetAccessKeyLastUsedResourcesResponseBodyResources
}

type GetAccessKeyLastUsedResourcesResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJhY2NvdW50IjoiMTQyNDM3OTU4NjM4NzE2MSIsImV2ZW50SWQiOiI3MkJDRTExRi02OTU3LTQ0NUItQjY0MC1CNEUyMkM4NUEwQzgiLCJsb2dJZCI6IjgyLTE0MjQzNzk1ODYzODcxNjEiLCJ0aW1lIjoxNjAyMzExNTQwMD****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 145318BE-DEE1-4C57-AA7C-5BE7D34A6AE0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of returned resources.
	//
	// This parameter is required.
	Resources []*GetAccessKeyLastUsedResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s GetAccessKeyLastUsedResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyLastUsedResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetAccessKeyLastUsedResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccessKeyLastUsedResourcesResponseBody) GetResources() []*GetAccessKeyLastUsedResourcesResponseBodyResources {
	return s.Resources
}

func (s *GetAccessKeyLastUsedResourcesResponseBody) SetNextToken(v string) *GetAccessKeyLastUsedResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetAccessKeyLastUsedResourcesResponseBody) SetRequestId(v string) *GetAccessKeyLastUsedResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccessKeyLastUsedResourcesResponseBody) SetResources(v []*GetAccessKeyLastUsedResourcesResponseBodyResources) *GetAccessKeyLastUsedResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *GetAccessKeyLastUsedResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAccessKeyLastUsedResourcesResponseBodyResources struct {
	// The event details.
	//
	// example:
	//
	// {
	//
	//   "eventId": "239EB588-CD24-522E-B0B5-174A1A58****",
	//
	//   "eventVersion": 1,
	//
	//   "eventSource": "ecs.cn-hangzhou.aliyuncs.com",
	//
	//   "sourceIpAddress": "``10.10.**.**``",
	//
	//   "eventType": "ApiCall",
	//
	//   "userIdentity": {
	//
	//     "accountId": "104758519118****",
	//
	//     "principalId": "24549429003625****",
	//
	//     "type": "ram-user",
	//
	//     "userName": "alice"
	//
	//   },
	//
	//   "serviceName": "Ecs",
	//
	//   "apiVersion": "2016-01-20",
	//
	//   "requestId": "239EB588-CD24-522E-B0B5-174A1A588BE0",
	//
	//   "eventTime": "2021-08-05T09:21:32Z",
	//
	//   "isGlobal": false,
	//
	//   "acsRegion": "cn-hangzhou",
	//
	//   "eventName": "DescribeInstances"
	//
	// }
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The resource name.
	//
	// example:
	//
	// i-bp1ltva99x1a****
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The event source.
	//
	// Valid values:
	//
	// 	- Internal
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     other events
	//
	//     <!-- -->
	//
	// 	- ManagementEvent
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     management events
	//
	//     <!-- -->
	//
	// 	- DataEvent
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     data events
	//
	//     <!-- -->
	//
	// example:
	//
	// ManagementEvent
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The timestamp when the resource was used. Unit: millisecond.
	//
	// example:
	//
	// 1657247532000
	UsedTimestamp *int64 `json:"UsedTimestamp,omitempty" xml:"UsedTimestamp,omitempty"`
}

func (s GetAccessKeyLastUsedResourcesResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyLastUsedResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedResourcesResponseBodyResources) GetDetail() *string {
	return s.Detail
}

func (s *GetAccessKeyLastUsedResourcesResponseBodyResources) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetAccessKeyLastUsedResourcesResponseBodyResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetAccessKeyLastUsedResourcesResponseBodyResources) GetSource() *string {
	return s.Source
}

func (s *GetAccessKeyLastUsedResourcesResponseBodyResources) GetUsedTimestamp() *int64 {
	return s.UsedTimestamp
}

func (s *GetAccessKeyLastUsedResourcesResponseBodyResources) SetDetail(v string) *GetAccessKeyLastUsedResourcesResponseBodyResources {
	s.Detail = &v
	return s
}

func (s *GetAccessKeyLastUsedResourcesResponseBodyResources) SetResourceName(v string) *GetAccessKeyLastUsedResourcesResponseBodyResources {
	s.ResourceName = &v
	return s
}

func (s *GetAccessKeyLastUsedResourcesResponseBodyResources) SetResourceType(v string) *GetAccessKeyLastUsedResourcesResponseBodyResources {
	s.ResourceType = &v
	return s
}

func (s *GetAccessKeyLastUsedResourcesResponseBodyResources) SetSource(v string) *GetAccessKeyLastUsedResourcesResponseBodyResources {
	s.Source = &v
	return s
}

func (s *GetAccessKeyLastUsedResourcesResponseBodyResources) SetUsedTimestamp(v int64) *GetAccessKeyLastUsedResourcesResponseBodyResources {
	s.UsedTimestamp = &v
	return s
}

func (s *GetAccessKeyLastUsedResourcesResponseBodyResources) Validate() error {
	return dara.Validate(s)
}
