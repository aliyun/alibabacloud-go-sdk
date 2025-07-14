// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelayTicketExpireTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DelayTicketExpireTimeResponseBody
	GetRequestId() *string
	SetResult(v bool) *DelayTicketExpireTimeResponseBody
	GetResult() *bool
	SetSuccess(v bool) *DelayTicketExpireTimeResponseBody
	GetSuccess() *bool
}

type DelayTicketExpireTimeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the extension is successful. Valid values:
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

func (s DelayTicketExpireTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DelayTicketExpireTimeResponseBody) GoString() string {
	return s.String()
}

func (s *DelayTicketExpireTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DelayTicketExpireTimeResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DelayTicketExpireTimeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DelayTicketExpireTimeResponseBody) SetRequestId(v string) *DelayTicketExpireTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DelayTicketExpireTimeResponseBody) SetResult(v bool) *DelayTicketExpireTimeResponseBody {
	s.Result = &v
	return s
}

func (s *DelayTicketExpireTimeResponseBody) SetSuccess(v bool) *DelayTicketExpireTimeResponseBody {
	s.Success = &v
	return s
}

func (s *DelayTicketExpireTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
