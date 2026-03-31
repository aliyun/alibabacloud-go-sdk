// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangePasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChangePasswordResponseBody
	GetRequestId() *string
}

type ChangePasswordResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangePasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangePasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ChangePasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangePasswordResponseBody) SetRequestId(v string) *ChangePasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangePasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
