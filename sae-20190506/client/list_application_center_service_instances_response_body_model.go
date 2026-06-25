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
	Code      *string                                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListApplicationCenterServiceInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceId   *string                                                `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
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
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	ServiceName       *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	TemplateName      *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
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
