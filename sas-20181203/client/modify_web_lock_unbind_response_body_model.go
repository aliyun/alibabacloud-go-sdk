// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebLockUnbindResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyWebLockUnbindResponseBody
	GetRequestId() *string
}

type ModifyWebLockUnbindResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// F0A8A039-930D-5EC1-97C8-43F05776188A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebLockUnbindResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebLockUnbindResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebLockUnbindResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWebLockUnbindResponseBody) SetRequestId(v string) *ModifyWebLockUnbindResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWebLockUnbindResponseBody) Validate() error {
	return dara.Validate(s)
}
