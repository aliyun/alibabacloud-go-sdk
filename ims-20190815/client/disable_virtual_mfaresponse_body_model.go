// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableVirtualMFAResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableVirtualMFAResponseBody
	GetRequestId() *string
}

type DisableVirtualMFAResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableVirtualMFAResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableVirtualMFAResponseBody) GoString() string {
	return s.String()
}

func (s *DisableVirtualMFAResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableVirtualMFAResponseBody) SetRequestId(v string) *DisableVirtualMFAResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableVirtualMFAResponseBody) Validate() error {
	return dara.Validate(s)
}
