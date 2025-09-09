// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartServiceInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartServiceInstanceResponseBody
	GetRequestId() *string
}

type StartServiceInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2E91D771-0183-52CE-86CB-882D99B2CB77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartServiceInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartServiceInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartServiceInstanceResponseBody) SetRequestId(v string) *StartServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartServiceInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
