// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableControlPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnablementStatus(v string) *DisableControlPolicyResponseBody
	GetEnablementStatus() *string
	SetRequestId(v string) *DisableControlPolicyResponseBody
	GetRequestId() *string
}

type DisableControlPolicyResponseBody struct {
	// The status of the Control Policy feature. Valid values:
	//
	// 	- Enabled: The feature is enabled.
	//
	// 	- PendingEnable: The feature is being enabled.
	//
	// 	- Disabled: The feature is disabled.
	//
	// 	- PendingDisable: The feature is being disabled.
	//
	// example:
	//
	// PendingDisable
	EnablementStatus *string `json:"EnablementStatus,omitempty" xml:"EnablementStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7C709979-451D-4C92-835D-7DDCCAA562E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableControlPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DisableControlPolicyResponseBody) GetEnablementStatus() *string {
	return s.EnablementStatus
}

func (s *DisableControlPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableControlPolicyResponseBody) SetEnablementStatus(v string) *DisableControlPolicyResponseBody {
	s.EnablementStatus = &v
	return s
}

func (s *DisableControlPolicyResponseBody) SetRequestId(v string) *DisableControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableControlPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
