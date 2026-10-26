// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTicket4AiProResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTicket4AiProResponseBody
	GetRequestId() *string
	SetResult(v string) *CreateTicket4AiProResponseBody
	GetResult() *string
	SetSuccess(v bool) *CreateTicket4AiProResponseBody
	GetSuccess() *bool
}

type CreateTicket4AiProResponseBody struct {
	// example:
	//
	// D787E1A**********DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTicket4AiProResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTicket4AiProResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTicket4AiProResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTicket4AiProResponseBody) GetResult() *string {
	return s.Result
}

func (s *CreateTicket4AiProResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTicket4AiProResponseBody) SetRequestId(v string) *CreateTicket4AiProResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTicket4AiProResponseBody) SetResult(v string) *CreateTicket4AiProResponseBody {
	s.Result = &v
	return s
}

func (s *CreateTicket4AiProResponseBody) SetSuccess(v bool) *CreateTicket4AiProResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTicket4AiProResponseBody) Validate() error {
	return dara.Validate(s)
}
