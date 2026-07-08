// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserAuthConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateUserAuthConfigResponseBody
	GetRequestId() *string
}

type UpdateUserAuthConfigResponseBody struct {
	// The unique request ID.
	//
	// example:
	//
	// 83C6468F-3D68-5791-860E-29AB8FCACC73
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserAuthConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserAuthConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserAuthConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserAuthConfigResponseBody) SetRequestId(v string) *UpdateUserAuthConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserAuthConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
