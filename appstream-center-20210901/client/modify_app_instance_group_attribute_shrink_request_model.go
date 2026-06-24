// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppInstanceGroupAttributeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *ModifyAppInstanceGroupAttributeShrinkRequest
	GetAppInstanceGroupId() *string
	SetAppInstanceGroupName(v string) *ModifyAppInstanceGroupAttributeShrinkRequest
	GetAppInstanceGroupName() *string
	SetNetworkShrink(v string) *ModifyAppInstanceGroupAttributeShrinkRequest
	GetNetworkShrink() *string
	SetNodePoolShrink(v string) *ModifyAppInstanceGroupAttributeShrinkRequest
	GetNodePoolShrink() *string
	SetPerSessionPerApp(v bool) *ModifyAppInstanceGroupAttributeShrinkRequest
	GetPerSessionPerApp() *bool
	SetPreOpenAppId(v string) *ModifyAppInstanceGroupAttributeShrinkRequest
	GetPreOpenAppId() *string
	SetPreOpenMode(v string) *ModifyAppInstanceGroupAttributeShrinkRequest
	GetPreOpenMode() *string
	SetProductType(v string) *ModifyAppInstanceGroupAttributeShrinkRequest
	GetProductType() *string
	SetSecurityPolicyShrink(v string) *ModifyAppInstanceGroupAttributeShrinkRequest
	GetSecurityPolicyShrink() *string
	SetSessionTimeout(v int32) *ModifyAppInstanceGroupAttributeShrinkRequest
	GetSessionTimeout() *int32
	SetStoragePolicyShrink(v string) *ModifyAppInstanceGroupAttributeShrinkRequest
	GetStoragePolicyShrink() *string
}

type ModifyAppInstanceGroupAttributeShrinkRequest struct {
	// The delivery group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The delivery group name.
	//
	// example:
	//
	// 办公应用
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The network configuration.
	//
	// > To use this parameter, submit a ticket.
	NetworkShrink *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The resource group object.
	NodePoolShrink *string `json:"NodePool,omitempty" xml:"NodePool,omitempty"`
	// Specifies whether to allow only one application per session.
	//
	// - If enabled, opening multiple applications within a delivery group allocates a separate session for each application, consuming more sessions.
	//
	// example:
	//
	// false
	PerSessionPerApp *bool `json:"PerSessionPerApp,omitempty" xml:"PerSessionPerApp,omitempty"`
	// The AppId of the pre-open application. If the PreOpenMode parameter is set to `SINGLE_APP`, PreOpenAppId cannot be an empty string.
	//
	// example:
	//
	// ca-b2ronxxd****
	PreOpenAppId *string `json:"PreOpenAppId,omitempty" xml:"PreOpenAppId,omitempty"`
	// The pre-open mode.
	//
	// example:
	//
	// OFF
	PreOpenMode *string `json:"PreOpenMode,omitempty" xml:"PreOpenMode,omitempty"`
	// The product type.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The security policy.
	SecurityPolicyShrink *string `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty"`
	// The session retention duration after disconnection, in minutes. After an end user session is disconnected, the session is retained for the specified duration before being logged off. Set this parameter to `-1` to retain the session indefinitely. Valid values: -1 and 3 to 300. Default value: `15`.
	//
	// example:
	//
	// 15
	SessionTimeout *int32 `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	// The storage policy.
	StoragePolicyShrink *string `json:"StoragePolicy,omitempty" xml:"StoragePolicy,omitempty"`
}

func (s ModifyAppInstanceGroupAttributeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) GetAppInstanceGroupName() *string {
	return s.AppInstanceGroupName
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) GetNetworkShrink() *string {
	return s.NetworkShrink
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) GetNodePoolShrink() *string {
	return s.NodePoolShrink
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) GetPerSessionPerApp() *bool {
	return s.PerSessionPerApp
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) GetPreOpenAppId() *string {
	return s.PreOpenAppId
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) GetPreOpenMode() *string {
	return s.PreOpenMode
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) GetSecurityPolicyShrink() *string {
	return s.SecurityPolicyShrink
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) GetSessionTimeout() *int32 {
	return s.SessionTimeout
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) GetStoragePolicyShrink() *string {
	return s.StoragePolicyShrink
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetAppInstanceGroupId(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetAppInstanceGroupName(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetNetworkShrink(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.NetworkShrink = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetNodePoolShrink(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.NodePoolShrink = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetPerSessionPerApp(v bool) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.PerSessionPerApp = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetPreOpenAppId(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.PreOpenAppId = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetPreOpenMode(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.PreOpenMode = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetProductType(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetSecurityPolicyShrink(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.SecurityPolicyShrink = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetSessionTimeout(v int32) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.SessionTimeout = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetStoragePolicyShrink(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.StoragePolicyShrink = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
