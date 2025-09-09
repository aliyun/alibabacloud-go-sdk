// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserPublicKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUserPublicKeyResponseBody
	GetRequestId() *string
}

type DeleteUserPublicKeyResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserPublicKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserPublicKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserPublicKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserPublicKeyResponseBody) SetRequestId(v string) *DeleteUserPublicKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserPublicKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
