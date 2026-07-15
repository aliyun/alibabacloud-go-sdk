// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAddConsumerGroupConsumersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchAddConsumerGroupConsumersResponseBody
	GetCode() *string
	SetData(v *BatchAddConsumerGroupConsumersResponseBodyData) *BatchAddConsumerGroupConsumersResponseBody
	GetData() *BatchAddConsumerGroupConsumersResponseBodyData
	SetMessage(v string) *BatchAddConsumerGroupConsumersResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchAddConsumerGroupConsumersResponseBody
	GetRequestId() *string
}

type BatchAddConsumerGroupConsumersResponseBody struct {
	// example:
	//
	// Ok
	Code *string                                         `json:"code,omitempty" xml:"code,omitempty"`
	Data *BatchAddConsumerGroupConsumersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s BatchAddConsumerGroupConsumersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchAddConsumerGroupConsumersResponseBody) GoString() string {
	return s.String()
}

func (s *BatchAddConsumerGroupConsumersResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchAddConsumerGroupConsumersResponseBody) GetData() *BatchAddConsumerGroupConsumersResponseBodyData {
	return s.Data
}

func (s *BatchAddConsumerGroupConsumersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchAddConsumerGroupConsumersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchAddConsumerGroupConsumersResponseBody) SetCode(v string) *BatchAddConsumerGroupConsumersResponseBody {
	s.Code = &v
	return s
}

func (s *BatchAddConsumerGroupConsumersResponseBody) SetData(v *BatchAddConsumerGroupConsumersResponseBodyData) *BatchAddConsumerGroupConsumersResponseBody {
	s.Data = v
	return s
}

func (s *BatchAddConsumerGroupConsumersResponseBody) SetMessage(v string) *BatchAddConsumerGroupConsumersResponseBody {
	s.Message = &v
	return s
}

func (s *BatchAddConsumerGroupConsumersResponseBody) SetRequestId(v string) *BatchAddConsumerGroupConsumersResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchAddConsumerGroupConsumersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchAddConsumerGroupConsumersResponseBodyData struct {
	FailedConsumerIds  []*string `json:"failedConsumerIds,omitempty" xml:"failedConsumerIds,omitempty" type:"Repeated"`
	SkippedConsumerIds []*string `json:"skippedConsumerIds,omitempty" xml:"skippedConsumerIds,omitempty" type:"Repeated"`
	SuccessConsumerIds []*string `json:"successConsumerIds,omitempty" xml:"successConsumerIds,omitempty" type:"Repeated"`
}

func (s BatchAddConsumerGroupConsumersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchAddConsumerGroupConsumersResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchAddConsumerGroupConsumersResponseBodyData) GetFailedConsumerIds() []*string {
	return s.FailedConsumerIds
}

func (s *BatchAddConsumerGroupConsumersResponseBodyData) GetSkippedConsumerIds() []*string {
	return s.SkippedConsumerIds
}

func (s *BatchAddConsumerGroupConsumersResponseBodyData) GetSuccessConsumerIds() []*string {
	return s.SuccessConsumerIds
}

func (s *BatchAddConsumerGroupConsumersResponseBodyData) SetFailedConsumerIds(v []*string) *BatchAddConsumerGroupConsumersResponseBodyData {
	s.FailedConsumerIds = v
	return s
}

func (s *BatchAddConsumerGroupConsumersResponseBodyData) SetSkippedConsumerIds(v []*string) *BatchAddConsumerGroupConsumersResponseBodyData {
	s.SkippedConsumerIds = v
	return s
}

func (s *BatchAddConsumerGroupConsumersResponseBodyData) SetSuccessConsumerIds(v []*string) *BatchAddConsumerGroupConsumersResponseBodyData {
	s.SuccessConsumerIds = v
	return s
}

func (s *BatchAddConsumerGroupConsumersResponseBodyData) Validate() error {
	return dara.Validate(s)
}
