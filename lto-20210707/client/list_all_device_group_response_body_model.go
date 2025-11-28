// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllDeviceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAllDeviceGroupResponseBody
	GetCode() *string
	SetData(v []*ListAllDeviceGroupResponseBodyData) *ListAllDeviceGroupResponseBody
	GetData() []*ListAllDeviceGroupResponseBodyData
	SetHttpStatusCode(v int32) *ListAllDeviceGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAllDeviceGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAllDeviceGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAllDeviceGroupResponseBody
	GetSuccess() *bool
}

type ListAllDeviceGroupResponseBody struct {
	Code           *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListAllDeviceGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAllDeviceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAllDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllDeviceGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAllDeviceGroupResponseBody) GetData() []*ListAllDeviceGroupResponseBodyData {
	return s.Data
}

func (s *ListAllDeviceGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAllDeviceGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAllDeviceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAllDeviceGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAllDeviceGroupResponseBody) SetCode(v string) *ListAllDeviceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListAllDeviceGroupResponseBody) SetData(v []*ListAllDeviceGroupResponseBodyData) *ListAllDeviceGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListAllDeviceGroupResponseBody) SetHttpStatusCode(v int32) *ListAllDeviceGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAllDeviceGroupResponseBody) SetMessage(v string) *ListAllDeviceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListAllDeviceGroupResponseBody) SetRequestId(v string) *ListAllDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAllDeviceGroupResponseBody) SetSuccess(v bool) *ListAllDeviceGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ListAllDeviceGroupResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAllDeviceGroupResponseBodyData struct {
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListAllDeviceGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAllDeviceGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAllDeviceGroupResponseBodyData) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *ListAllDeviceGroupResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListAllDeviceGroupResponseBodyData) SetDeviceGroupId(v string) *ListAllDeviceGroupResponseBodyData {
	s.DeviceGroupId = &v
	return s
}

func (s *ListAllDeviceGroupResponseBodyData) SetName(v string) *ListAllDeviceGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListAllDeviceGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
