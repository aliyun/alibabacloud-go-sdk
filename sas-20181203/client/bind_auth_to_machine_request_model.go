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
	SetClientToken(v string) *BindAuthToMachineRequest
	GetClientToken() *string
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
	SetProductCode(v string) *BindAuthToMachineRequest
	GetProductCode() *string
	SetResourceDirectoryAccountId(v int64) *BindAuthToMachineRequest
	GetResourceDirectoryAccountId() *int64
	SetUnBind(v []*string) *BindAuthToMachineRequest
	GetUnBind() []*string
}

type BindAuthToMachineRequest struct {
	// The authorization version of the asset. Valid values:
	//
	// - **6**: Anti-virus Edition.
	//
	// - **5**: Premium Edition.
	//
	// - **3**: Enterprise Edition.
	//
	// - **7**: Ultimate Edition.
	//
	// - **10**: Value-added service Edition.
	//
	// example:
	//
	// 6
	AuthVersion *int32 `json:"AuthVersion,omitempty" xml:"AuthVersion,omitempty"`
	// Specifies whether to enable automatic binding. Valid values:
	//
	// - **0**: Disable automatic binding.
	//
	// - **1**: Enable automatic binding.
	//
	// example:
	//
	// 1
	AutoBind *int32 `json:"AutoBind,omitempty" xml:"AutoBind,omitempty"`
	// The collection of UUIDs to bind.
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
	// The client token that is used to ensure the idempotence of the request. Use a different token for each request. The token supports only ASCII characters and cannot exceed 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The conditions for searching assets. This parameter is in JSON format. Pay attention to letter case when you specify this parameter.
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
	// - **OR**: The search conditions are evaluated using a logical OR.
	//
	// - **AND**: The search conditions are evaluated using a logical AND.
	//
	// example:
	//
	// OR
	LogicalExp *string `json:"LogicalExp,omitempty" xml:"LogicalExp,omitempty"`
	// The order version associated with the pre-binding operation. Valid values:
	//
	// - **level7**: Anti-virus Edition.
	//
	// - **level3**: Premium Edition.
	//
	// - **level2**: Enterprise Edition.
	//
	// - **level8**: Ultimate Edition.
	//
	// - **level10**: Value-added service only.
	//
	// example:
	//
	// level2
	NtmVersion *string `json:"NtmVersion,omitempty" xml:"NtmVersion,omitempty"`
	// The order ID associated with the pre-binding operation.
	//
	// example:
	//
	// 233016**0482
	PreBindOrderId *int64  `json:"PreBindOrderId,omitempty" xml:"PreBindOrderId,omitempty"`
	ProductCode    *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the member accounts (Alibaba Cloud account) in the resource directory.
	//
	// >Call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 16670360956*****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The collection of UUIDs to unbind.
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

func (s *BindAuthToMachineRequest) GetClientToken() *string {
	return s.ClientToken
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

func (s *BindAuthToMachineRequest) GetProductCode() *string {
	return s.ProductCode
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

func (s *BindAuthToMachineRequest) SetClientToken(v string) *BindAuthToMachineRequest {
	s.ClientToken = &v
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

func (s *BindAuthToMachineRequest) SetProductCode(v string) *BindAuthToMachineRequest {
	s.ProductCode = &v
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
