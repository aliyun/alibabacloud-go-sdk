// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationProvisioningResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableApplicationProvisioningResponseBody
	GetRequestId() *string
}

type DisableApplicationProvisioningResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableApplicationProvisioningResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationProvisioningResponseBody) GoString() string {
	return s.String()
}

func (s *DisableApplicationProvisioningResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableApplicationProvisioningResponseBody) SetRequestId(v string) *DisableApplicationProvisioningResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableApplicationProvisioningResponseBody) Validate() error {
	return dara.Validate(s)
}
