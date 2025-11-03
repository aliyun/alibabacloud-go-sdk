// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventBusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateEventBusResponseBody
	GetCode() *string
	SetData(v *CreateEventBusResponseBodyData) *CreateEventBusResponseBody
	GetData() *CreateEventBusResponseBodyData
	SetMessage(v string) *CreateEventBusResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateEventBusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateEventBusResponseBody
	GetSuccess() *bool
}

type CreateEventBusResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful. Other values indicate that the request failed. For more information about error codes, see Error codes.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateEventBusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	//
	// example:
	//
	// The event bus [xxxx] not existed!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A995F07C-E503-5881-9962-9CECA8566876
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. The value true indicates that the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateEventBusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEventBusResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEventBusResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateEventBusResponseBody) GetData() *CreateEventBusResponseBodyData {
	return s.Data
}

func (s *CreateEventBusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateEventBusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEventBusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateEventBusResponseBody) SetCode(v string) *CreateEventBusResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEventBusResponseBody) SetData(v *CreateEventBusResponseBodyData) *CreateEventBusResponseBody {
	s.Data = v
	return s
}

func (s *CreateEventBusResponseBody) SetMessage(v string) *CreateEventBusResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEventBusResponseBody) SetRequestId(v string) *CreateEventBusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEventBusResponseBody) SetSuccess(v bool) *CreateEventBusResponseBody {
	s.Success = &v
	return s
}

func (s *CreateEventBusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEventBusResponseBodyData struct {
	// The Alibaba Cloud Resource Name (ARN) of the event bus.
	//
	// example:
	//
	// acs:eventbridge:cn-hangzhou:123456789098****:eventbus/MyEventBus
	EventBusARN *string `json:"EventBusARN,omitempty" xml:"EventBusARN,omitempty"`
}

func (s CreateEventBusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateEventBusResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateEventBusResponseBodyData) GetEventBusARN() *string {
	return s.EventBusARN
}

func (s *CreateEventBusResponseBodyData) SetEventBusARN(v string) *CreateEventBusResponseBodyData {
	s.EventBusARN = &v
	return s
}

func (s *CreateEventBusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
