// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantOperatorPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GrantOperatorPermissionResponseBody
	GetRequestId() *string
}

type GrantOperatorPermissionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A438072A-E2E7-5509-9A3F-66293512A820
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantOperatorPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantOperatorPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *GrantOperatorPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantOperatorPermissionResponseBody) SetRequestId(v string) *GrantOperatorPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantOperatorPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
