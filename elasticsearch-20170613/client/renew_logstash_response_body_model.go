// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewLogstashResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RenewLogstashResponseBody
	GetRequestId() *string
	SetResult(v bool) *RenewLogstashResponseBody
	GetResult() *bool
}

type RenewLogstashResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result. Valid values:
	//
	// 	- true: The cluster is renewed.
	//
	// 	- false: The cluster fails to be renewed.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s RenewLogstashResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewLogstashResponseBody) GoString() string {
	return s.String()
}

func (s *RenewLogstashResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewLogstashResponseBody) GetResult() *bool {
	return s.Result
}

func (s *RenewLogstashResponseBody) SetRequestId(v string) *RenewLogstashResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewLogstashResponseBody) SetResult(v bool) *RenewLogstashResponseBody {
	s.Result = &v
	return s
}

func (s *RenewLogstashResponseBody) Validate() error {
	return dara.Validate(s)
}
