// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetControlPolicyEnablementStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnablementStatus(v string) *GetControlPolicyEnablementStatusResponseBody
	GetEnablementStatus() *string
	SetRequestId(v string) *GetControlPolicyEnablementStatusResponseBody
	GetRequestId() *string
}

type GetControlPolicyEnablementStatusResponseBody struct {
	// The status of the Control Policy feature. Valid values:
	//
	// 	- Enabled: The Control Policy feature is enabled.
	//
	// 	- PendingEnable: The Control Policy feature is being enabled.
	//
	// 	- Disabled: The Control Policy feature is disabled.
	//
	// 	- PendingDisable: The Control Policy feature is being disabled.
	//
	// example:
	//
	// Disabled
	EnablementStatus *string `json:"EnablementStatus,omitempty" xml:"EnablementStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1DC39A4E-3B52-4EFE-9F93-4897D7FFA0C4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetControlPolicyEnablementStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetControlPolicyEnablementStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetControlPolicyEnablementStatusResponseBody) GetEnablementStatus() *string {
	return s.EnablementStatus
}

func (s *GetControlPolicyEnablementStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetControlPolicyEnablementStatusResponseBody) SetEnablementStatus(v string) *GetControlPolicyEnablementStatusResponseBody {
	s.EnablementStatus = &v
	return s
}

func (s *GetControlPolicyEnablementStatusResponseBody) SetRequestId(v string) *GetControlPolicyEnablementStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetControlPolicyEnablementStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
