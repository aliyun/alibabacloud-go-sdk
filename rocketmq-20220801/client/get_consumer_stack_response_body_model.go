// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumerStackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetConsumerStackResponseBody
	GetCode() *string
	SetData(v *GetConsumerStackResponseBodyData) *GetConsumerStackResponseBody
	GetData() *GetConsumerStackResponseBodyData
	SetDynamicCode(v string) *GetConsumerStackResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetConsumerStackResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *GetConsumerStackResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetConsumerStackResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetConsumerStackResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetConsumerStackResponseBody
	GetSuccess() *bool
}

type GetConsumerStackResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data *GetConsumerStackResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30F2CBC7-F69D-5D78-9661-0254C9E1FBFA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetConsumerStackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerStackResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsumerStackResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetConsumerStackResponseBody) GetData() *GetConsumerStackResponseBodyData {
	return s.Data
}

func (s *GetConsumerStackResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetConsumerStackResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetConsumerStackResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetConsumerStackResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetConsumerStackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConsumerStackResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetConsumerStackResponseBody) SetCode(v string) *GetConsumerStackResponseBody {
	s.Code = &v
	return s
}

func (s *GetConsumerStackResponseBody) SetData(v *GetConsumerStackResponseBodyData) *GetConsumerStackResponseBody {
	s.Data = v
	return s
}

func (s *GetConsumerStackResponseBody) SetDynamicCode(v string) *GetConsumerStackResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetConsumerStackResponseBody) SetDynamicMessage(v string) *GetConsumerStackResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetConsumerStackResponseBody) SetHttpStatusCode(v int32) *GetConsumerStackResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetConsumerStackResponseBody) SetMessage(v string) *GetConsumerStackResponseBody {
	s.Message = &v
	return s
}

func (s *GetConsumerStackResponseBody) SetRequestId(v string) *GetConsumerStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConsumerStackResponseBody) SetSuccess(v bool) *GetConsumerStackResponseBody {
	s.Success = &v
	return s
}

func (s *GetConsumerStackResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetConsumerStackResponseBodyData struct {
	// The ID of the consumer group.
	//
	// example:
	//
	// CID-TEST
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Stack Information.
	Stacks []*GetConsumerStackResponseBodyDataStacks `json:"stacks,omitempty" xml:"stacks,omitempty" type:"Repeated"`
}

func (s GetConsumerStackResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerStackResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetConsumerStackResponseBodyData) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *GetConsumerStackResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetConsumerStackResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetConsumerStackResponseBodyData) GetStacks() []*GetConsumerStackResponseBodyDataStacks {
	return s.Stacks
}

func (s *GetConsumerStackResponseBodyData) SetConsumerGroupId(v string) *GetConsumerStackResponseBodyData {
	s.ConsumerGroupId = &v
	return s
}

func (s *GetConsumerStackResponseBodyData) SetInstanceId(v string) *GetConsumerStackResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetConsumerStackResponseBodyData) SetRegionId(v string) *GetConsumerStackResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetConsumerStackResponseBodyData) SetStacks(v []*GetConsumerStackResponseBodyDataStacks) *GetConsumerStackResponseBodyData {
	s.Stacks = v
	return s
}

func (s *GetConsumerStackResponseBodyData) Validate() error {
	if s.Stacks != nil {
		for _, item := range s.Stacks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetConsumerStackResponseBodyDataStacks struct {
	// Thread id.
	//
	// example:
	//
	// 123
	Thread *string `json:"thread,omitempty" xml:"thread,omitempty"`
	// Stack Information.
	Tracks []*string `json:"tracks,omitempty" xml:"tracks,omitempty" type:"Repeated"`
}

func (s GetConsumerStackResponseBodyDataStacks) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerStackResponseBodyDataStacks) GoString() string {
	return s.String()
}

func (s *GetConsumerStackResponseBodyDataStacks) GetThread() *string {
	return s.Thread
}

func (s *GetConsumerStackResponseBodyDataStacks) GetTracks() []*string {
	return s.Tracks
}

func (s *GetConsumerStackResponseBodyDataStacks) SetThread(v string) *GetConsumerStackResponseBodyDataStacks {
	s.Thread = &v
	return s
}

func (s *GetConsumerStackResponseBodyDataStacks) SetTracks(v []*string) *GetConsumerStackResponseBodyDataStacks {
	s.Tracks = v
	return s
}

func (s *GetConsumerStackResponseBodyDataStacks) Validate() error {
	return dara.Validate(s)
}
