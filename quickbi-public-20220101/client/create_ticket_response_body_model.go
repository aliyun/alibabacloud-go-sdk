// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTicketResponseBody
	GetRequestId() *string
	SetResult(v string) *CreateTicketResponseBody
	GetResult() *string
	SetSuccess(v bool) *CreateTicketResponseBody
	GetSuccess() *bool
}

type CreateTicketResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The generated ticket value.
	//
	// example:
	//
	// ccd3428c-****-****-a608-26bae29dffee
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Value range:
	//
	// - true: Request succeeded
	//
	// - false: Request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTicketResponseBody) GetResult() *string {
	return s.Result
}

func (s *CreateTicketResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTicketResponseBody) SetRequestId(v string) *CreateTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTicketResponseBody) SetResult(v string) *CreateTicketResponseBody {
	s.Result = &v
	return s
}

func (s *CreateTicketResponseBody) SetSuccess(v bool) *CreateTicketResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTicketResponseBody) Validate() error {
	return dara.Validate(s)
}
