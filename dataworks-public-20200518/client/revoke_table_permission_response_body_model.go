// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeTablePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeTablePermissionResponseBody
	GetRequestId() *string
	SetRevokeSuccess(v bool) *RevokeTablePermissionResponseBody
	GetRevokeSuccess() *bool
}

type RevokeTablePermissionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0bc1ec92159376****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the permissions are revoked.
	//
	// example:
	//
	// true
	RevokeSuccess *bool `json:"RevokeSuccess,omitempty" xml:"RevokeSuccess,omitempty"`
}

func (s RevokeTablePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeTablePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeTablePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeTablePermissionResponseBody) GetRevokeSuccess() *bool {
	return s.RevokeSuccess
}

func (s *RevokeTablePermissionResponseBody) SetRequestId(v string) *RevokeTablePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeTablePermissionResponseBody) SetRevokeSuccess(v bool) *RevokeTablePermissionResponseBody {
	s.RevokeSuccess = &v
	return s
}

func (s *RevokeTablePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
