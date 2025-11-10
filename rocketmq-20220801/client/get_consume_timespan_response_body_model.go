// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumeTimespanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetConsumeTimespanResponseBody
	GetCode() *string
	SetData(v *GetConsumeTimespanResponseBodyData) *GetConsumeTimespanResponseBody
	GetData() *GetConsumeTimespanResponseBodyData
	SetDynamicCode(v string) *GetConsumeTimespanResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetConsumeTimespanResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *GetConsumeTimespanResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetConsumeTimespanResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetConsumeTimespanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetConsumeTimespanResponseBody
	GetSuccess() *bool
}

type GetConsumeTimespanResponseBody struct {
	// example:
	//
	// MissingInstanceId
	Code *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetConsumeTimespanResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// xxx
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// example:
	//
	// xxx
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// C115601B-8736-5BBF-AC99-7FEAE1245A80
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetConsumeTimespanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConsumeTimespanResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsumeTimespanResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetConsumeTimespanResponseBody) GetData() *GetConsumeTimespanResponseBodyData {
	return s.Data
}

func (s *GetConsumeTimespanResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetConsumeTimespanResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetConsumeTimespanResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetConsumeTimespanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetConsumeTimespanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConsumeTimespanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetConsumeTimespanResponseBody) SetCode(v string) *GetConsumeTimespanResponseBody {
	s.Code = &v
	return s
}

func (s *GetConsumeTimespanResponseBody) SetData(v *GetConsumeTimespanResponseBodyData) *GetConsumeTimespanResponseBody {
	s.Data = v
	return s
}

func (s *GetConsumeTimespanResponseBody) SetDynamicCode(v string) *GetConsumeTimespanResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetConsumeTimespanResponseBody) SetDynamicMessage(v string) *GetConsumeTimespanResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetConsumeTimespanResponseBody) SetHttpStatusCode(v int32) *GetConsumeTimespanResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetConsumeTimespanResponseBody) SetMessage(v string) *GetConsumeTimespanResponseBody {
	s.Message = &v
	return s
}

func (s *GetConsumeTimespanResponseBody) SetRequestId(v string) *GetConsumeTimespanResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConsumeTimespanResponseBody) SetSuccess(v bool) *GetConsumeTimespanResponseBody {
	s.Success = &v
	return s
}

func (s *GetConsumeTimespanResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetConsumeTimespanResponseBodyData struct {
	// example:
	//
	// 1570761026400
	ConsumeTimestamp *int64 `json:"consumeTimestamp,omitempty" xml:"consumeTimestamp,omitempty"`
	// example:
	//
	// CID-TEST
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 1570761026833
	MaxTimestamp *int64 `json:"maxTimestamp,omitempty" xml:"maxTimestamp,omitempty"`
	// example:
	//
	// 1570761026222
	MinTimestamp *int64 `json:"minTimestamp,omitempty" xml:"minTimestamp,omitempty"`
	// example:
	//
	// topic_test
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s GetConsumeTimespanResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetConsumeTimespanResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetConsumeTimespanResponseBodyData) GetConsumeTimestamp() *int64 {
	return s.ConsumeTimestamp
}

func (s *GetConsumeTimespanResponseBodyData) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *GetConsumeTimespanResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetConsumeTimespanResponseBodyData) GetMaxTimestamp() *int64 {
	return s.MaxTimestamp
}

func (s *GetConsumeTimespanResponseBodyData) GetMinTimestamp() *int64 {
	return s.MinTimestamp
}

func (s *GetConsumeTimespanResponseBodyData) GetTopicName() *string {
	return s.TopicName
}

func (s *GetConsumeTimespanResponseBodyData) SetConsumeTimestamp(v int64) *GetConsumeTimespanResponseBodyData {
	s.ConsumeTimestamp = &v
	return s
}

func (s *GetConsumeTimespanResponseBodyData) SetConsumerGroupId(v string) *GetConsumeTimespanResponseBodyData {
	s.ConsumerGroupId = &v
	return s
}

func (s *GetConsumeTimespanResponseBodyData) SetInstanceId(v string) *GetConsumeTimespanResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetConsumeTimespanResponseBodyData) SetMaxTimestamp(v int64) *GetConsumeTimespanResponseBodyData {
	s.MaxTimestamp = &v
	return s
}

func (s *GetConsumeTimespanResponseBodyData) SetMinTimestamp(v int64) *GetConsumeTimespanResponseBodyData {
	s.MinTimestamp = &v
	return s
}

func (s *GetConsumeTimespanResponseBodyData) SetTopicName(v string) *GetConsumeTimespanResponseBodyData {
	s.TopicName = &v
	return s
}

func (s *GetConsumeTimespanResponseBodyData) Validate() error {
	return dara.Validate(s)
}
