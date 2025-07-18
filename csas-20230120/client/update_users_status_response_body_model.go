// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUsersStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateUsersStatusResponseBody
	GetRequestId() *string
}

type UpdateUsersStatusResponseBody struct {
	// example:
	//
	// 47363C2B-1AAA-5954-8847-0E50FCC54117
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUsersStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUsersStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUsersStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUsersStatusResponseBody) SetRequestId(v string) *UpdateUsersStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUsersStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
