// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateEipAddressWithRCInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnassociateEipAddressWithRCInstanceResponseBody
	GetRequestId() *string
}

type UnassociateEipAddressWithRCInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CDEAC7BF-D64B-54A1-9051-BE9AC0990E68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnassociateEipAddressWithRCInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnassociateEipAddressWithRCInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UnassociateEipAddressWithRCInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnassociateEipAddressWithRCInstanceResponseBody) SetRequestId(v string) *UnassociateEipAddressWithRCInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnassociateEipAddressWithRCInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
