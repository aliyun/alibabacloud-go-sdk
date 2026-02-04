// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewFreeLicenseEndTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RenewFreeLicenseEndTimeResponseBody
	GetRequestId() *string
}

type RenewFreeLicenseEndTimeResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewFreeLicenseEndTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewFreeLicenseEndTimeResponseBody) GoString() string {
	return s.String()
}

func (s *RenewFreeLicenseEndTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewFreeLicenseEndTimeResponseBody) SetRequestId(v string) *RenewFreeLicenseEndTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewFreeLicenseEndTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
