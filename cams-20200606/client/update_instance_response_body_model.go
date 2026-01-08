// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateInstanceResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateInstanceResponseBody
	GetCode() *string
	SetData(v *UpdateInstanceResponseBodyData) *UpdateInstanceResponseBody
	GetData() *UpdateInstanceResponseBodyData
	SetMessage(v string) *UpdateInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateInstanceResponseBody
	GetSuccess() *bool
}

type UpdateInstanceResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// NULL
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateInstanceResponseBody) GetData() *UpdateInstanceResponseBodyData {
	return s.Data
}

func (s *UpdateInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateInstanceResponseBody) SetAccessDeniedDetail(v string) *UpdateInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetCode(v string) *UpdateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetData(v *UpdateInstanceResponseBodyData) *UpdateInstanceResponseBody {
	s.Data = v
	return s
}

func (s *UpdateInstanceResponseBody) SetMessage(v string) *UpdateInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetRequestId(v string) *UpdateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetSuccess(v bool) *UpdateInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateInstanceResponseBodyData struct {
	// example:
	//
	// VIBER
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// example:
	//
	// des
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// des
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// example:
	//
	// 293992992
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ins
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 示例值示例值
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// example:
	//
	// 2020-12-01 00:00:00
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s UpdateInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBodyData) GetChannelType() *string {
	return s.ChannelType
}

func (s *UpdateInstanceResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *UpdateInstanceResponseBodyData) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *UpdateInstanceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateInstanceResponseBodyData) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *UpdateInstanceResponseBodyData) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *UpdateInstanceResponseBodyData) SetChannelType(v string) *UpdateInstanceResponseBodyData {
	s.ChannelType = &v
	return s
}

func (s *UpdateInstanceResponseBodyData) SetDescription(v string) *UpdateInstanceResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateInstanceResponseBodyData) SetInstanceDescription(v string) *UpdateInstanceResponseBodyData {
	s.InstanceDescription = &v
	return s
}

func (s *UpdateInstanceResponseBodyData) SetInstanceId(v string) *UpdateInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceResponseBodyData) SetInstanceName(v string) *UpdateInstanceResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstanceResponseBodyData) SetResourceRegionId(v string) *UpdateInstanceResponseBodyData {
	s.ResourceRegionId = &v
	return s
}

func (s *UpdateInstanceResponseBodyData) SetSubmitTime(v string) *UpdateInstanceResponseBodyData {
	s.SubmitTime = &v
	return s
}

func (s *UpdateInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
