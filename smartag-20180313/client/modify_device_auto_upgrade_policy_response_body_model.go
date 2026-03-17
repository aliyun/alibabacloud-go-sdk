// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDeviceAutoUpgradePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDeviceAutoUpgradePolicyResponseBody
	GetRequestId() *string
}

type ModifyDeviceAutoUpgradePolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F03C6897-2284-4BC8-94B4-1467BD992A2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDeviceAutoUpgradePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDeviceAutoUpgradePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDeviceAutoUpgradePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDeviceAutoUpgradePolicyResponseBody) SetRequestId(v string) *ModifyDeviceAutoUpgradePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDeviceAutoUpgradePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
