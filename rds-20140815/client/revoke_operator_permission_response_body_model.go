// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeOperatorPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeOperatorPermissionResponseBody
	GetRequestId() *string
}

type RevokeOperatorPermissionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 842B73C8-5776-4BD9-9872-69C8C46DD7D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeOperatorPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeOperatorPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeOperatorPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeOperatorPermissionResponseBody) SetRequestId(v string) *RevokeOperatorPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeOperatorPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
