// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAuthToMachineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthVersion(v int32) *BindAuthToMachineRequest
	GetAuthVersion() *int32
	SetAutoBind(v int32) *BindAuthToMachineRequest
	GetAutoBind() *int32
	SetBind(v []*string) *BindAuthToMachineRequest
	GetBind() []*string
	SetBindAll(v bool) *BindAuthToMachineRequest
	GetBindAll() *bool
	SetCriteria(v string) *BindAuthToMachineRequest
	GetCriteria() *string
	SetIsPreBind(v int32) *BindAuthToMachineRequest
	GetIsPreBind() *int32
	SetLogicalExp(v string) *BindAuthToMachineRequest
	GetLogicalExp() *string
	SetNtmVersion(v string) *BindAuthToMachineRequest
	GetNtmVersion() *string
	SetPreBindOrderId(v int64) *BindAuthToMachineRequest
	GetPreBindOrderId() *int64
	SetResourceDirectoryAccountId(v int64) *BindAuthToMachineRequest
	GetResourceDirectoryAccountId() *int64
	SetUnBind(v []*string) *BindAuthToMachineRequest
	GetUnBind() []*string
}

type BindAuthToMachineRequest struct {
	// The authorization version of the asset. Valid values:
	//
	// - **6**: Anti-virus Edition
	//
	// - **5**: Advanced Edition
	//
	// - **3**: Enterprise Edition
	//
	// - **7**: Ultimate Edition
	//
	// - **10**: Value-added Service Edition.
	//
	// example:
	//
	// 6
	AuthVersion *int32 `json:"AuthVersion,omitempty" xml:"AuthVersion,omitempty"`
	// Specifies whether to enable automatic binding. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
	//
	// example:
	//
	// 1
	AutoBind *int32 `json:"AutoBind,omitempty" xml:"AutoBind,omitempty"`
	// The UUIDs of the servers to bind.
	//
	// > **Bind*	- and **UnBind*	- cannot both be empty.
	Bind []*string `json:"Bind,omitempty" xml:"Bind,omitempty" type:"Repeated"`
	// Specifies whether to bind all assets. Default value: **false**. Valid values:
	//
	// - **true**: Bind all assets.
	//
	// - **false**: Do not bind all assets.
	//
	// example:
	//
	// true
	BindAll *bool `json:"BindAll,omitempty" xml:"BindAll,omitempty"`
	// The search conditions for assets. This parameter is in JSON format. Note that the parameter values are case-sensitive.
	//
	// > You can search for assets by instance ID, instance name, VPC ID, region, or public IP address. Call the [DescribeCriteria](~~DescribeCriteria~~) operation to query the supported search conditions.
	//
	// example:
	//
	// [{"name":"riskStatus","value":"YES"},{"name":"internetIp","value":"1.2.XX.XX"}]
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// Specifies whether this is a pre-binding operation. Valid values:
	//
	// - **0**: No.
	//
	// - **1**: Yes.
	//
	//
	// > After pre-binding is enabled, the corresponding authorization quota is automatically bound to the specified servers after the purchase is completed.
	//
	// example:
	//
	// 1
	IsPreBind *int32 `json:"IsPreBind,omitempty" xml:"IsPreBind,omitempty"`
	// The logical relationship among multiple search conditions. Default value: **OR**. Valid values:
	//
	// - **OR**: The search conditions are evaluated with a logical OR.
	//
	// - **AND**: The search conditions are evaluated with a logical AND.
	//
	// example:
	//
	// OR
	LogicalExp *string `json:"LogicalExp,omitempty" xml:"LogicalExp,omitempty"`
	// The order version associated with the pre-binding. Valid values:
	//
	// - **level7**: Anti-virus Edition
	//
	// - **level3**: Advanced Edition
	//
	// - **level2**: Enterprise Edition
	//
	// - **level8**: Ultimate Edition
	//
	// - **level10**: Value-added service only.
	//
	// example:
	//
	// level2
	NtmVersion *string `json:"NtmVersion,omitempty" xml:"NtmVersion,omitempty"`
	// The order ID associated with the pre-binding.
	//
	// example:
	//
	// 233016**0482
	PreBindOrderId *int64 `json:"PreBindOrderId,omitempty" xml:"PreBindOrderId,omitempty"`
	// The ID of the member accounts in the resource folder (Alibaba Cloud account).
	//
	// > Invoke the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 16670360956*****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The UUIDs of the servers to unbind.
	//
	// > **Bind*	- and **UnBind*	- cannot both be empty.
	UnBind []*string `json:"UnBind,omitempty" xml:"UnBind,omitempty" type:"Repeated"`
}

func (s BindAuthToMachineRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAuthToMachineRequest) GoString() string {
	return s.String()
}

func (s *BindAuthToMachineRequest) GetAuthVersion() *int32 {
	return s.AuthVersion
}

func (s *BindAuthToMachineRequest) GetAutoBind() *int32 {
	return s.AutoBind
}

func (s *BindAuthToMachineRequest) GetBind() []*string {
	return s.Bind
}

func (s *BindAuthToMachineRequest) GetBindAll() *bool {
	return s.BindAll
}

func (s *BindAuthToMachineRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *BindAuthToMachineRequest) GetIsPreBind() *int32 {
	return s.IsPreBind
}

func (s *BindAuthToMachineRequest) GetLogicalExp() *string {
	return s.LogicalExp
}

func (s *BindAuthToMachineRequest) GetNtmVersion() *string {
	return s.NtmVersion
}

func (s *BindAuthToMachineRequest) GetPreBindOrderId() *int64 {
	return s.PreBindOrderId
}

func (s *BindAuthToMachineRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *BindAuthToMachineRequest) GetUnBind() []*string {
	return s.UnBind
}

func (s *BindAuthToMachineRequest) SetAuthVersion(v int32) *BindAuthToMachineRequest {
	s.AuthVersion = &v
	return s
}

func (s *BindAuthToMachineRequest) SetAutoBind(v int32) *BindAuthToMachineRequest {
	s.AutoBind = &v
	return s
}

func (s *BindAuthToMachineRequest) SetBind(v []*string) *BindAuthToMachineRequest {
	s.Bind = v
	return s
}

func (s *BindAuthToMachineRequest) SetBindAll(v bool) *BindAuthToMachineRequest {
	s.BindAll = &v
	return s
}

func (s *BindAuthToMachineRequest) SetCriteria(v string) *BindAuthToMachineRequest {
	s.Criteria = &v
	return s
}

func (s *BindAuthToMachineRequest) SetIsPreBind(v int32) *BindAuthToMachineRequest {
	s.IsPreBind = &v
	return s
}

func (s *BindAuthToMachineRequest) SetLogicalExp(v string) *BindAuthToMachineRequest {
	s.LogicalExp = &v
	return s
}

func (s *BindAuthToMachineRequest) SetNtmVersion(v string) *BindAuthToMachineRequest {
	s.NtmVersion = &v
	return s
}

func (s *BindAuthToMachineRequest) SetPreBindOrderId(v int64) *BindAuthToMachineRequest {
	s.PreBindOrderId = &v
	return s
}

func (s *BindAuthToMachineRequest) SetResourceDirectoryAccountId(v int64) *BindAuthToMachineRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *BindAuthToMachineRequest) SetUnBind(v []*string) *BindAuthToMachineRequest {
	s.UnBind = v
	return s
}

func (s *BindAuthToMachineRequest) Validate() error {
	return dara.Validate(s)
}
