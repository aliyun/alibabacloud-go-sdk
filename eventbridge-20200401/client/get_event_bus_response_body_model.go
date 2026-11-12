// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventBusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetEventBusResponseBody
	GetCode() *string
	SetData(v *GetEventBusResponseBodyData) *GetEventBusResponseBody
	GetData() *GetEventBusResponseBodyData
	SetMessage(v string) *GetEventBusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetEventBusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetEventBusResponseBody
	GetSuccess() *bool
}

type GetEventBusResponseBody struct {
	// The response code. A value of Success indicates that the request is successful. Other values indicate error codes. For more information, see Error codes.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *GetEventBusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// EventBusNotExist
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// d5bfc188-4452-4ba7-b73a-a9005e522439
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEventBusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEventBusResponseBody) GoString() string {
	return s.String()
}

func (s *GetEventBusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetEventBusResponseBody) GetData() *GetEventBusResponseBodyData {
	return s.Data
}

func (s *GetEventBusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetEventBusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEventBusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetEventBusResponseBody) SetCode(v string) *GetEventBusResponseBody {
	s.Code = &v
	return s
}

func (s *GetEventBusResponseBody) SetData(v *GetEventBusResponseBodyData) *GetEventBusResponseBody {
	s.Data = v
	return s
}

func (s *GetEventBusResponseBody) SetMessage(v string) *GetEventBusResponseBody {
	s.Message = &v
	return s
}

func (s *GetEventBusResponseBody) SetRequestId(v string) *GetEventBusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEventBusResponseBody) SetSuccess(v bool) *GetEventBusResponseBody {
	s.Success = &v
	return s
}

func (s *GetEventBusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEventBusResponseBodyData struct {
	// The creation timestamp.
	//
	// example:
	//
	// 1641781825000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The description of the event bus.
	//
	// example:
	//
	// demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the event bus.
	//
	// example:
	//
	// acs:eventbridge:cn-hangzhou:123456789098****:eventbus/MyEventBus
	EventBusARN *string `json:"EventBusARN,omitempty" xml:"EventBusARN,omitempty"`
	// The name of the event bus.
	//
	// example:
	//
	// MyEventBus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The ID of the resource group to which the event bus belongs.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag information.
	Tags []*GetEventBusResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetEventBusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetEventBusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEventBusResponseBodyData) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *GetEventBusResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetEventBusResponseBodyData) GetEventBusARN() *string {
	return s.EventBusARN
}

func (s *GetEventBusResponseBodyData) GetEventBusName() *string {
	return s.EventBusName
}

func (s *GetEventBusResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetEventBusResponseBodyData) GetTags() []*GetEventBusResponseBodyDataTags {
	return s.Tags
}

func (s *GetEventBusResponseBodyData) SetCreateTimestamp(v int64) *GetEventBusResponseBodyData {
	s.CreateTimestamp = &v
	return s
}

func (s *GetEventBusResponseBodyData) SetDescription(v string) *GetEventBusResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetEventBusResponseBodyData) SetEventBusARN(v string) *GetEventBusResponseBodyData {
	s.EventBusARN = &v
	return s
}

func (s *GetEventBusResponseBodyData) SetEventBusName(v string) *GetEventBusResponseBodyData {
	s.EventBusName = &v
	return s
}

func (s *GetEventBusResponseBodyData) SetResourceGroupId(v string) *GetEventBusResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetEventBusResponseBodyData) SetTags(v []*GetEventBusResponseBodyDataTags) *GetEventBusResponseBodyData {
	s.Tags = v
	return s
}

func (s *GetEventBusResponseBodyData) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetEventBusResponseBodyDataTags struct {
	// The key of the first tag.
	//
	// example:
	//
	// department
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the first tag.
	//
	// example:
	//
	// R&D
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEventBusResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s GetEventBusResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *GetEventBusResponseBodyDataTags) GetKey() *string {
	return s.Key
}

func (s *GetEventBusResponseBodyDataTags) GetValue() *string {
	return s.Value
}

func (s *GetEventBusResponseBodyDataTags) SetKey(v string) *GetEventBusResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *GetEventBusResponseBodyDataTags) SetValue(v string) *GetEventBusResponseBodyDataTags {
	s.Value = &v
	return s
}

func (s *GetEventBusResponseBodyDataTags) Validate() error {
	return dara.Validate(s)
}
