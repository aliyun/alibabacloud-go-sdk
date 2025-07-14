// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTicket4CopilotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTicket4CopilotResponseBody
	GetRequestId() *string
	SetResult(v string) *CreateTicket4CopilotResponseBody
	GetResult() *string
	SetSuccess(v bool) *CreateTicket4CopilotResponseBody
	GetSuccess() *bool
}

type CreateTicket4CopilotResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// D787************05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// ID of the Smart Q module to be embedded.
	//
	// example:
	//
	// f5eeb52e-d9c2-4a8b-80e3-47ab55c2****
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: The request succeeded
	//
	// - false: The request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTicket4CopilotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTicket4CopilotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTicket4CopilotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTicket4CopilotResponseBody) GetResult() *string {
	return s.Result
}

func (s *CreateTicket4CopilotResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTicket4CopilotResponseBody) SetRequestId(v string) *CreateTicket4CopilotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTicket4CopilotResponseBody) SetResult(v string) *CreateTicket4CopilotResponseBody {
	s.Result = &v
	return s
}

func (s *CreateTicket4CopilotResponseBody) SetSuccess(v bool) *CreateTicket4CopilotResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTicket4CopilotResponseBody) Validate() error {
	return dara.Validate(s)
}
