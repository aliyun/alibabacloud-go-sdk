// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantUserPermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GrantUserPermissionsResponseBody
	GetRequestId() *string
}

type GrantUserPermissionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5A7C9E37-C171-584F-9A99-869B48C4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantUserPermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantUserPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *GrantUserPermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantUserPermissionsResponseBody) SetRequestId(v string) *GrantUserPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantUserPermissionsResponseBody) Validate() error {
	return dara.Validate(s)
}
