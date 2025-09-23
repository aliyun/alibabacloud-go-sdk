// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableLayer7CCResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableLayer7CCResponseBody
	GetRequestId() *string
}

type DisableLayer7CCResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableLayer7CCResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableLayer7CCResponseBody) GoString() string {
	return s.String()
}

func (s *DisableLayer7CCResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableLayer7CCResponseBody) SetRequestId(v string) *DisableLayer7CCResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableLayer7CCResponseBody) Validate() error {
	return dara.Validate(s)
}
