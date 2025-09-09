// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSharedAccountPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSharedAccountPermissionResponseBody
	GetRequestId() *string
}

type UpdateSharedAccountPermissionResponseBody struct {
	// RequestId
	//
	// example:
	//
	// C68B41B4-A646-5680-8A33-67884E3823A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSharedAccountPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSharedAccountPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSharedAccountPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSharedAccountPermissionResponseBody) SetRequestId(v string) *UpdateSharedAccountPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSharedAccountPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
