// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RenewInstanceResponseBody
	GetRequestId() *string
	SetResult(v bool) *RenewInstanceResponseBody
	GetResult() *bool
}

type RenewInstanceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return results:
	//
	// 	- true: renewal successfully
	//
	// 	- false: renewal failed
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s RenewInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewInstanceResponseBody) GetResult() *bool {
	return s.Result
}

func (s *RenewInstanceResponseBody) SetRequestId(v string) *RenewInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewInstanceResponseBody) SetResult(v bool) *RenewInstanceResponseBody {
	s.Result = &v
	return s
}

func (s *RenewInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
