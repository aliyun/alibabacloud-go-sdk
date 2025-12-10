// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityGroupPermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSecurityGroupPermissionsResponseBody
	GetRequestId() *string
}

type DeleteSecurityGroupPermissionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AAE90880-4970-4D81-A534-A6C0F3631F74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSecurityGroupPermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityGroupPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupPermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSecurityGroupPermissionsResponseBody) SetRequestId(v string) *DeleteSecurityGroupPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecurityGroupPermissionsResponseBody) Validate() error {
	return dara.Validate(s)
}
