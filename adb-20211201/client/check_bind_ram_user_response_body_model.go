// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckBindRamUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckBindRamUserResponseBody
	GetRequestId() *string
	SetResult(v bool) *CheckBindRamUserResponseBody
	GetResult() *bool
}

type CheckBindRamUserResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2FB9DCA3-DA56-5B43-A9A0-68E3D0E6AA84
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result returned by the node. Valid values:
	//
	// - **true**: The current Resource Access Management (RAM) user is attached to a database account.
	//
	// - **false**: The current Resource Access Management (RAM) user is not attached to a database account.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CheckBindRamUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckBindRamUserResponseBody) GoString() string {
	return s.String()
}

func (s *CheckBindRamUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckBindRamUserResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CheckBindRamUserResponseBody) SetRequestId(v string) *CheckBindRamUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckBindRamUserResponseBody) SetResult(v bool) *CheckBindRamUserResponseBody {
	s.Result = &v
	return s
}

func (s *CheckBindRamUserResponseBody) Validate() error {
	return dara.Validate(s)
}
