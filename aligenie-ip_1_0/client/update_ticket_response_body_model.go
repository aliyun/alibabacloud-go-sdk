// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpdateTicketResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateTicketResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateTicketResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *UpdateTicketResponseBody
	GetStatusCode() *int32
}

type UpdateTicketResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7***726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s UpdateTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTicketResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTicketResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTicketResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateTicketResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTicketResponseBody) SetMessage(v string) *UpdateTicketResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTicketResponseBody) SetRequestId(v string) *UpdateTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTicketResponseBody) SetResult(v bool) *UpdateTicketResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateTicketResponseBody) SetStatusCode(v int32) *UpdateTicketResponseBody {
	s.StatusCode = &v
	return s
}

func (s *UpdateTicketResponseBody) Validate() error {
	return dara.Validate(s)
}
