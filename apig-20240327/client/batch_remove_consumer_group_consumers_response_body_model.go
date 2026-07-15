// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRemoveConsumerGroupConsumersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchRemoveConsumerGroupConsumersResponseBody
	GetCode() *string
	SetData(v *BatchRemoveConsumerGroupConsumersResponseBodyData) *BatchRemoveConsumerGroupConsumersResponseBody
	GetData() *BatchRemoveConsumerGroupConsumersResponseBodyData
	SetMessage(v string) *BatchRemoveConsumerGroupConsumersResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchRemoveConsumerGroupConsumersResponseBody
	GetRequestId() *string
}

type BatchRemoveConsumerGroupConsumersResponseBody struct {
	// example:
	//
	// Ok
	Code *string                                            `json:"code,omitempty" xml:"code,omitempty"`
	Data *BatchRemoveConsumerGroupConsumersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s BatchRemoveConsumerGroupConsumersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchRemoveConsumerGroupConsumersResponseBody) GoString() string {
	return s.String()
}

func (s *BatchRemoveConsumerGroupConsumersResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchRemoveConsumerGroupConsumersResponseBody) GetData() *BatchRemoveConsumerGroupConsumersResponseBodyData {
	return s.Data
}

func (s *BatchRemoveConsumerGroupConsumersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchRemoveConsumerGroupConsumersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchRemoveConsumerGroupConsumersResponseBody) SetCode(v string) *BatchRemoveConsumerGroupConsumersResponseBody {
	s.Code = &v
	return s
}

func (s *BatchRemoveConsumerGroupConsumersResponseBody) SetData(v *BatchRemoveConsumerGroupConsumersResponseBodyData) *BatchRemoveConsumerGroupConsumersResponseBody {
	s.Data = v
	return s
}

func (s *BatchRemoveConsumerGroupConsumersResponseBody) SetMessage(v string) *BatchRemoveConsumerGroupConsumersResponseBody {
	s.Message = &v
	return s
}

func (s *BatchRemoveConsumerGroupConsumersResponseBody) SetRequestId(v string) *BatchRemoveConsumerGroupConsumersResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchRemoveConsumerGroupConsumersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchRemoveConsumerGroupConsumersResponseBodyData struct {
	FailedConsumerIds  []*string `json:"failedConsumerIds,omitempty" xml:"failedConsumerIds,omitempty" type:"Repeated"`
	SkippedConsumerIds []*string `json:"skippedConsumerIds,omitempty" xml:"skippedConsumerIds,omitempty" type:"Repeated"`
	SuccessConsumerIds []*string `json:"successConsumerIds,omitempty" xml:"successConsumerIds,omitempty" type:"Repeated"`
}

func (s BatchRemoveConsumerGroupConsumersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchRemoveConsumerGroupConsumersResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchRemoveConsumerGroupConsumersResponseBodyData) GetFailedConsumerIds() []*string {
	return s.FailedConsumerIds
}

func (s *BatchRemoveConsumerGroupConsumersResponseBodyData) GetSkippedConsumerIds() []*string {
	return s.SkippedConsumerIds
}

func (s *BatchRemoveConsumerGroupConsumersResponseBodyData) GetSuccessConsumerIds() []*string {
	return s.SuccessConsumerIds
}

func (s *BatchRemoveConsumerGroupConsumersResponseBodyData) SetFailedConsumerIds(v []*string) *BatchRemoveConsumerGroupConsumersResponseBodyData {
	s.FailedConsumerIds = v
	return s
}

func (s *BatchRemoveConsumerGroupConsumersResponseBodyData) SetSkippedConsumerIds(v []*string) *BatchRemoveConsumerGroupConsumersResponseBodyData {
	s.SkippedConsumerIds = v
	return s
}

func (s *BatchRemoveConsumerGroupConsumersResponseBodyData) SetSuccessConsumerIds(v []*string) *BatchRemoveConsumerGroupConsumersResponseBodyData {
	s.SuccessConsumerIds = v
	return s
}

func (s *BatchRemoveConsumerGroupConsumersResponseBodyData) Validate() error {
	return dara.Validate(s)
}
