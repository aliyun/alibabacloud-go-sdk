// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationCenterServiceInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListApplicationCenterServiceInstancesResponseBody
	GetCode() *string
	SetData(v *ListApplicationCenterServiceInstancesResponseBodyData) *ListApplicationCenterServiceInstancesResponseBody
	GetData() *ListApplicationCenterServiceInstancesResponseBodyData
	SetErrorCode(v string) *ListApplicationCenterServiceInstancesResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListApplicationCenterServiceInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListApplicationCenterServiceInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListApplicationCenterServiceInstancesResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ListApplicationCenterServiceInstancesResponseBody
	GetTraceId() *string
}

type ListApplicationCenterServiceInstancesResponseBody struct {
	// The API status code or POP error code. Valid values: 2xx: Success. 3xx: Redirection. 4xx: Client error. 5xx: Server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned results.
	Data *ListApplicationCenterServiceInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. If the request is successful, the ErrorCode parameter is not returned. If the request fails, the ErrorCode parameter is returned. For more information, see the Error codes section in this topic.
	//
	// example:
	//
	// Empty
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information. If the request is successful, success is returned. If the request fails, a specific error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID.
	//
	// example:
	//
	// 0a981dd515966966104121683d****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListApplicationCenterServiceInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationCenterServiceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationCenterServiceInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListApplicationCenterServiceInstancesResponseBody) GetData() *ListApplicationCenterServiceInstancesResponseBodyData {
	return s.Data
}

func (s *ListApplicationCenterServiceInstancesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListApplicationCenterServiceInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListApplicationCenterServiceInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationCenterServiceInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListApplicationCenterServiceInstancesResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ListApplicationCenterServiceInstancesResponseBody) SetCode(v string) *ListApplicationCenterServiceInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListApplicationCenterServiceInstancesResponseBody) SetData(v *ListApplicationCenterServiceInstancesResponseBodyData) *ListApplicationCenterServiceInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListApplicationCenterServiceInstancesResponseBody) SetErrorCode(v string) *ListApplicationCenterServiceInstancesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListApplicationCenterServiceInstancesResponseBody) SetMessage(v string) *ListApplicationCenterServiceInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListApplicationCenterServiceInstancesResponseBody) SetRequestId(v string) *ListApplicationCenterServiceInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationCenterServiceInstancesResponseBody) SetSuccess(v bool) *ListApplicationCenterServiceInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListApplicationCenterServiceInstancesResponseBody) SetTraceId(v string) *ListApplicationCenterServiceInstancesResponseBody {
	s.TraceId = &v
	return s
}

func (s *ListApplicationCenterServiceInstancesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApplicationCenterServiceInstancesResponseBodyData struct {
	// The list of service instances.
	ServiceInstances []*ListApplicationCenterServiceInstancesResponseBodyDataServiceInstances `json:"ServiceInstances,omitempty" xml:"ServiceInstances,omitempty" type:"Repeated"`
}

func (s ListApplicationCenterServiceInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationCenterServiceInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListApplicationCenterServiceInstancesResponseBodyData) GetServiceInstances() []*ListApplicationCenterServiceInstancesResponseBodyDataServiceInstances {
	return s.ServiceInstances
}

func (s *ListApplicationCenterServiceInstancesResponseBodyData) SetServiceInstances(v []*ListApplicationCenterServiceInstancesResponseBodyDataServiceInstances) *ListApplicationCenterServiceInstancesResponseBodyData {
	s.ServiceInstances = v
	return s
}

func (s *ListApplicationCenterServiceInstancesResponseBodyData) Validate() error {
	if s.ServiceInstances != nil {
		for _, item := range s.ServiceInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationCenterServiceInstancesResponseBodyDataServiceInstances struct {
	// The creation time.
	//
	// example:
	//
	// 2025-10-28T02:18:51Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance name.
	//
	// example:
	//
	// tftestacc54
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The service instance ID.
	//
	// example:
	//
	// si-2063aea4b23b4781a26b
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The service name.
	//
	// example:
	//
	// user-service
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The template name.
	//
	// example:
	//
	// Dify HA
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListApplicationCenterServiceInstancesResponseBodyDataServiceInstances) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationCenterServiceInstancesResponseBodyDataServiceInstances) GoString() string {
	return s.String()
}

func (s *ListApplicationCenterServiceInstancesResponseBodyDataServiceInstances) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListApplicationCenterServiceInstancesResponseBodyDataServiceInstances) GetName() *string {
	return s.Name
}

func (s *ListApplicationCenterServiceInstancesResponseBodyDataServiceInstances) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *ListApplicationCenterServiceInstancesResponseBodyDataServiceInstances) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListApplicationCenterServiceInstancesResponseBodyDataServiceInstances) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListApplicationCenterServiceInstancesResponseBodyDataServiceInstances) SetCreateTime(v string) *ListApplicationCenterServiceInstancesResponseBodyDataServiceInstances {
	s.CreateTime = &v
	return s
}

func (s *ListApplicationCenterServiceInstancesResponseBodyDataServiceInstances) SetName(v string) *ListApplicationCenterServiceInstancesResponseBodyDataServiceInstances {
	s.Name = &v
	return s
}

func (s *ListApplicationCenterServiceInstancesResponseBodyDataServiceInstances) SetServiceInstanceId(v string) *ListApplicationCenterServiceInstancesResponseBodyDataServiceInstances {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListApplicationCenterServiceInstancesResponseBodyDataServiceInstances) SetServiceName(v string) *ListApplicationCenterServiceInstancesResponseBodyDataServiceInstances {
	s.ServiceName = &v
	return s
}

func (s *ListApplicationCenterServiceInstancesResponseBodyDataServiceInstances) SetTemplateName(v string) *ListApplicationCenterServiceInstancesResponseBodyDataServiceInstances {
	s.TemplateName = &v
	return s
}

func (s *ListApplicationCenterServiceInstancesResponseBodyDataServiceInstances) Validate() error {
	return dara.Validate(s)
}
