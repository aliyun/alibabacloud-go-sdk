// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTicketResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteTicketResponseBody
	GetResult() *bool
	SetSuccess(v bool) *DeleteTicketResponseBody
	GetSuccess() *bool
}

type DeleteTicketResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the deletion was successful. Possible values:
	//
	// - true: The request was successful
	//
	// - false: The request failed
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: The request was successful
	//
	// - false: The request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTicketResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTicketResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteTicketResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTicketResponseBody) SetRequestId(v string) *DeleteTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTicketResponseBody) SetResult(v bool) *DeleteTicketResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteTicketResponseBody) SetSuccess(v bool) *DeleteTicketResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTicketResponseBody) Validate() error {
	return dara.Validate(s)
}
