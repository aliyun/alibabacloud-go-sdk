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
	// The request ID.
	//
	// example:
	//
	// 2FB9DCA3-DA56-5B43-A9A0-68E3D0E6AA84
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result of the request. Valid values:
	//
	// 	- **true**: the database account is associated with a RAM user.
	//
	// 	- **false**: the database account is not associated with a RAM user.
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
