// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableBrandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableBrandResponseBody
	GetRequestId() *string
}

type DisableBrandResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableBrandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableBrandResponseBody) GoString() string {
	return s.String()
}

func (s *DisableBrandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableBrandResponseBody) SetRequestId(v string) *DisableBrandResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableBrandResponseBody) Validate() error {
	return dara.Validate(s)
}
