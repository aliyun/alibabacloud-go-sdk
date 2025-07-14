// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTicketNumResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTicketNumResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateTicketNumResponseBody
	GetResult() *bool
	SetSuccess(v bool) *UpdateTicketNumResponseBody
	GetSuccess() *bool
}

type UpdateTicketNumResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the update is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTicketNumResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTicketNumResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTicketNumResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTicketNumResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateTicketNumResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTicketNumResponseBody) SetRequestId(v string) *UpdateTicketNumResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTicketNumResponseBody) SetResult(v bool) *UpdateTicketNumResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateTicketNumResponseBody) SetSuccess(v bool) *UpdateTicketNumResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTicketNumResponseBody) Validate() error {
	return dara.Validate(s)
}
