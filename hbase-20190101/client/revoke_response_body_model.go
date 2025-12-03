// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevokeResponseBody
	GetRequestId() *string
}

type RevokeResponseBody struct {
	// example:
	//
	// C9085433-A56A-4089-B49A-DF5A4E2B7B06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeResponseBody) SetRequestId(v string) *RevokeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeResponseBody) Validate() error {
	return dara.Validate(s)
}
