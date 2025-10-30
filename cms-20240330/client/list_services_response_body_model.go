// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServicesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListServicesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListServicesResponseBody
	GetRequestId() *string
	SetServices(v []*ListServicesResponseBodyServices) *ListServicesResponseBody
	GetServices() []*ListServicesResponseBodyServices
	SetTotalCount(v int32) *ListServicesResponseBody
	GetTotalCount() *int32
}

type ListServicesResponseBody struct {
	// Maximum number of results to return, with a maximum value of 200
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// Pagination token
	//
	// example:
	//
	// 2-ba4d-4b9f-aa24-dcb067a30f1c
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Request ID
	//
	// example:
	//
	// 0CEC5375-C554-562B-A65F-9A629907C1F0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// List of service information.
	Services []*ListServicesResponseBodyServices `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
	// Total count
	//
	// example:
	//
	// 66
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServicesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServicesResponseBody) GetServices() []*ListServicesResponseBodyServices {
	return s.Services
}

func (s *ListServicesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListServicesResponseBody) SetMaxResults(v int32) *ListServicesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServicesResponseBody) SetNextToken(v string) *ListServicesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServicesResponseBody) SetRequestId(v string) *ListServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServicesResponseBody) SetServices(v []*ListServicesResponseBodyServices) *ListServicesResponseBody {
	s.Services = v
	return s
}

func (s *ListServicesResponseBody) SetTotalCount(v int32) *ListServicesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListServicesResponseBody) Validate() error {
	if s.Services != nil {
		for _, item := range s.Services {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServicesResponseBodyServices struct {
	// Extended information.
	//
	// example:
	//
	// {"language":"java"}
	Attributes *string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// Creation time
	//
	// example:
	//
	// 2025-07-01T02:23:59Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Service description, valid only when serviceType=RUM.
	//
	// example:
	//
	// workspace api monitor test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Display name, valid only when serviceType=RUM.
	//
	// example:
	//
	// test
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// Historical compatible ARMS application ID
	//
	// example:
	//
	// kgcsf@192197e828d51aa
	Pid *string `json:"pid,omitempty" xml:"pid,omitempty"`
	// Service ID
	//
	// example:
	//
	// jm2pl0yoqf@d4905cb11a4f218dfb0a8
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// Service name
	//
	// example:
	//
	// demo-app
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// Service status, valid only when serviceType=RUM.
	//
	// example:
	//
	// Running
	ServiceStatus *string `json:"serviceStatus,omitempty" xml:"serviceStatus,omitempty"`
	// Service type
	//
	// example:
	//
	// TRACE
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	// Workspace.
	//
	// example:
	//
	// default-cms-1192928460540589-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListServicesResponseBodyServices) String() string {
	return dara.Prettify(s)
}

func (s ListServicesResponseBodyServices) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServices) GetAttributes() *string {
	return s.Attributes
}

func (s *ListServicesResponseBodyServices) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListServicesResponseBodyServices) GetDescription() *string {
	return s.Description
}

func (s *ListServicesResponseBodyServices) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListServicesResponseBodyServices) GetPid() *string {
	return s.Pid
}

func (s *ListServicesResponseBodyServices) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListServicesResponseBodyServices) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListServicesResponseBodyServices) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *ListServicesResponseBodyServices) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListServicesResponseBodyServices) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListServicesResponseBodyServices) SetAttributes(v string) *ListServicesResponseBodyServices {
	s.Attributes = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetCreateTime(v string) *ListServicesResponseBodyServices {
	s.CreateTime = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetDescription(v string) *ListServicesResponseBodyServices {
	s.Description = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetDisplayName(v string) *ListServicesResponseBodyServices {
	s.DisplayName = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetPid(v string) *ListServicesResponseBodyServices {
	s.Pid = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetServiceId(v string) *ListServicesResponseBodyServices {
	s.ServiceId = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetServiceName(v string) *ListServicesResponseBodyServices {
	s.ServiceName = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetServiceStatus(v string) *ListServicesResponseBodyServices {
	s.ServiceStatus = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetServiceType(v string) *ListServicesResponseBodyServices {
	s.ServiceType = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetWorkspace(v string) *ListServicesResponseBodyServices {
	s.Workspace = &v
	return s
}

func (s *ListServicesResponseBodyServices) Validate() error {
	return dara.Validate(s)
}
