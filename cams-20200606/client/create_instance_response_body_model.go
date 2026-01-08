// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateInstanceResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateInstanceResponseBody
	GetCode() *string
	SetData(v *CreateInstanceResponseBodyData) *CreateInstanceResponseBody
	GetData() *CreateInstanceResponseBodyData
	SetMessage(v string) *CreateInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateInstanceResponseBody
	GetSuccess() *bool
}

type CreateInstanceResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// NULL
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2992939*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateInstanceResponseBody) GetData() *CreateInstanceResponseBodyData {
	return s.Data
}

func (s *CreateInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateInstanceResponseBody) SetAccessDeniedDetail(v string) *CreateInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateInstanceResponseBody) SetCode(v string) *CreateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceResponseBody) SetData(v *CreateInstanceResponseBodyData) *CreateInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateInstanceResponseBody) SetMessage(v string) *CreateInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetSuccess(v bool) *CreateInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateInstanceResponseBodyData struct {
	// example:
	//
	// 示例值
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// example:
	//
	// des
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// example:
	//
	// 2999292
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ins
	InstanceName *int64 `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 示例值示例值
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// example:
	//
	// 示例值
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
}

func (s CreateInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBodyData) GetChannelType() *string {
	return s.ChannelType
}

func (s *CreateInstanceResponseBodyData) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *CreateInstanceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstanceResponseBodyData) GetInstanceName() *int64 {
	return s.InstanceName
}

func (s *CreateInstanceResponseBodyData) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *CreateInstanceResponseBodyData) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *CreateInstanceResponseBodyData) SetChannelType(v string) *CreateInstanceResponseBodyData {
	s.ChannelType = &v
	return s
}

func (s *CreateInstanceResponseBodyData) SetInstanceDescription(v string) *CreateInstanceResponseBodyData {
	s.InstanceDescription = &v
	return s
}

func (s *CreateInstanceResponseBodyData) SetInstanceId(v string) *CreateInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBodyData) SetInstanceName(v int64) *CreateInstanceResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceResponseBodyData) SetResourceRegionId(v string) *CreateInstanceResponseBodyData {
	s.ResourceRegionId = &v
	return s
}

func (s *CreateInstanceResponseBodyData) SetSubmitTime(v string) *CreateInstanceResponseBodyData {
	s.SubmitTime = &v
	return s
}

func (s *CreateInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
