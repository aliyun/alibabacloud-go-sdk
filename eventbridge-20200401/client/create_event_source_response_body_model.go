// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateEventSourceResponseBody
	GetCode() *string
	SetData(v *CreateEventSourceResponseBodyData) *CreateEventSourceResponseBody
	GetData() *CreateEventSourceResponseBodyData
	SetMessage(v string) *CreateEventSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateEventSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateEventSourceResponseBody
	GetSuccess() *bool
}

type CreateEventSourceResponseBody struct {
	// The response code.
	//
	// - `Success`: The request was successful.
	//
	// - Other values indicate errors. For more information, see the "Error codes" section.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned by the request.
	Data *CreateEventSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned if the request is unsuccessful.
	//
	// example:
	//
	// Remote error. requestId: [A8EFABD2-95B9-1C46-9E01-xxxx], error code: [CreateRelatedResourceFailed], message: [Create related resource failed, EntityNotExist.Role : The role not exists: xxxx. \\r\\nRequestId : xxxx-168C-54ED-8FEB-BF11CB70AEB7]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2922208e-e1c6-43ee-bfd1-aca50263bc8a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. A value of `true` indicates that the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateEventSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEventSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEventSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateEventSourceResponseBody) GetData() *CreateEventSourceResponseBodyData {
	return s.Data
}

func (s *CreateEventSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateEventSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEventSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateEventSourceResponseBody) SetCode(v string) *CreateEventSourceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEventSourceResponseBody) SetData(v *CreateEventSourceResponseBodyData) *CreateEventSourceResponseBody {
	s.Data = v
	return s
}

func (s *CreateEventSourceResponseBody) SetMessage(v string) *CreateEventSourceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEventSourceResponseBody) SetRequestId(v string) *CreateEventSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEventSourceResponseBody) SetSuccess(v bool) *CreateEventSourceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateEventSourceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEventSourceResponseBodyData struct {
	// The Alibaba Cloud Resource Name (ARN) of the event source.
	//
	// example:
	//
	// acs:eventbridge:cn-hangzhou:164901546557****:eventbus/my-event-bus/eventsource/mymns.source
	EventSourceARN *string `json:"EventSourceARN,omitempty" xml:"EventSourceARN,omitempty"`
}

func (s CreateEventSourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateEventSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateEventSourceResponseBodyData) GetEventSourceARN() *string {
	return s.EventSourceARN
}

func (s *CreateEventSourceResponseBodyData) SetEventSourceARN(v string) *CreateEventSourceResponseBodyData {
	s.EventSourceARN = &v
	return s
}

func (s *CreateEventSourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
