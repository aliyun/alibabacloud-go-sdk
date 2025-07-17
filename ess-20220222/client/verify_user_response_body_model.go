// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *VerifyUserResponseBody
	GetRequestId() *string
}

type VerifyUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyUserResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyUserResponseBody) SetRequestId(v string) *VerifyUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyUserResponseBody) Validate() error {
	return dara.Validate(s)
}
