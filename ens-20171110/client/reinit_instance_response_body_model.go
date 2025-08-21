// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReinitInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReinitInstanceResponseBody
	GetRequestId() *string
}

type ReinitInstanceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C46C79B7-0C31-5947-9D86-82207661EADA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReinitInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReinitInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReinitInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReinitInstanceResponseBody) SetRequestId(v string) *ReinitInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReinitInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
