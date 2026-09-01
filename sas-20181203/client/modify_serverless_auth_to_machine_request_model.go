// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyServerlessAuthToMachineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCriteria(v string) *ModifyServerlessAuthToMachineRequest
	GetAppCriteria() *string
	SetAuthItem(v string) *ModifyServerlessAuthToMachineRequest
	GetAuthItem() *string
	SetAutoBind(v int32) *ModifyServerlessAuthToMachineRequest
	GetAutoBind() *int32
	SetBindAll(v bool) *ModifyServerlessAuthToMachineRequest
	GetBindAll() *bool
	SetBindAppList(v []*string) *ModifyServerlessAuthToMachineRequest
	GetBindAppList() []*string
	SetBindAssetType(v string) *ModifyServerlessAuthToMachineRequest
	GetBindAssetType() *string
	SetBindUuidList(v []*string) *ModifyServerlessAuthToMachineRequest
	GetBindUuidList() []*string
	SetClientToken(v string) *ModifyServerlessAuthToMachineRequest
	GetClientToken() *string
	SetCriteria(v string) *ModifyServerlessAuthToMachineRequest
	GetCriteria() *string
	SetLogicalExp(v string) *ModifyServerlessAuthToMachineRequest
	GetLogicalExp() *string
	SetNtmVersion(v string) *ModifyServerlessAuthToMachineRequest
	GetNtmVersion() *string
	SetPreBind(v int32) *ModifyServerlessAuthToMachineRequest
	GetPreBind() *int32
	SetPreBindOrderId(v int64) *ModifyServerlessAuthToMachineRequest
	GetPreBindOrderId() *int64
	SetResourceDirectoryUid(v int64) *ModifyServerlessAuthToMachineRequest
	GetResourceDirectoryUid() *int64
	SetUnBindAppList(v []*string) *ModifyServerlessAuthToMachineRequest
	GetUnBindAppList() []*string
	SetUnBindUuidList(v []*string) *ModifyServerlessAuthToMachineRequest
	GetUnBindUuidList() []*string
}

type ModifyServerlessAuthToMachineRequest struct {
	// The application query conditions.
	//
	// example:
	//
	// **7ad7e3a
	AppCriteria *string `json:"AppCriteria,omitempty" xml:"AppCriteria,omitempty"`
	// The instance type. Valid values:
	//
	// - **SERVERLESS**: Serverless asset.
	//
	// example:
	//
	// SERVERLESS
	AuthItem *string `json:"AuthItem,omitempty" xml:"AuthItem,omitempty"`
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
	// Specifies whether to bind all assets. Default value: **false**. Valid values:
	//
	// - **true**: Bind all assets.
	//
	// - **false**: Do not bind all assets.
	//
	// example:
	//
	// false
	BindAll *bool `json:"BindAll,omitempty" xml:"BindAll,omitempty"`
	// The list of application IDs to bind.
	//
	// > Obtain the IDs by calling the [ListMachineApps](~~ListMachineApps~~) operation.
	BindAppList []*string `json:"BindAppList,omitempty" xml:"BindAppList,omitempty" type:"Repeated"`
	// The Asset Type for the operation. Valid values:
	//
	// - **INSTANCE**: Instance.
	//
	// - **APP**: Application.
	//
	// example:
	//
	// APP
	BindAssetType *string `json:"BindAssetType,omitempty" xml:"BindAssetType,omitempty"`
	// The list of asset UUIDs to bind.
	BindUuidList []*string `json:"BindUuidList,omitempty" xml:"BindUuidList,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. Use a different token for each request. The token supports only ASCII characters and cannot exceed 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The search conditions for assets. This parameter is in JSON format. Pay attention to letter case when you specify this parameter.
	//
	// > You can search for assets by instance ID, instance name, VPC ID, region, public IP address, and other conditions. Call the [DescribeCriteria](~~DescribeCriteria~~) operation to query the supported search conditions.
	//
	// example:
	//
	// [{"name":"vulStatus","value":"YES","logicalExp":"AND"}]
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The logical relationship among multiple search conditions. Valid values:
	//
	// - **OR**: The search conditions are evaluated with a logical OR.
	//
	// - **AND**: The search conditions are evaluated with a logical AND.
	//
	// example:
	//
	// OR
	LogicalExp *string `json:"LogicalExp,omitempty" xml:"LogicalExp,omitempty"`
	// The NTM version code for pre-binding.
	//
	// example:
	//
	// level2
	NtmVersion *string `json:"NtmVersion,omitempty" xml:"NtmVersion,omitempty"`
	// Specifies whether to enable pre-binding. Valid values:
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
	PreBind *int32 `json:"PreBind,omitempty" xml:"PreBind,omitempty"`
	// The pre-binding order ID.
	//
	// example:
	//
	// 233016**0482
	PreBindOrderId *int64 `json:"PreBindOrderId,omitempty" xml:"PreBindOrderId,omitempty"`
	// The UID of the resource directory.
	//
	// example:
	//
	// 123456
	ResourceDirectoryUid *int64 `json:"ResourceDirectoryUid,omitempty" xml:"ResourceDirectoryUid,omitempty"`
	// The list of application IDs to unbind.
	//
	// > Obtain the IDs by calling the [ListMachineApps](~~ListMachineApps~~) operation.
	UnBindAppList []*string `json:"UnBindAppList,omitempty" xml:"UnBindAppList,omitempty" type:"Repeated"`
	// The list of asset UUIDs to unbind.
	UnBindUuidList []*string `json:"UnBindUuidList,omitempty" xml:"UnBindUuidList,omitempty" type:"Repeated"`
}

func (s ModifyServerlessAuthToMachineRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyServerlessAuthToMachineRequest) GoString() string {
	return s.String()
}

func (s *ModifyServerlessAuthToMachineRequest) GetAppCriteria() *string {
	return s.AppCriteria
}

func (s *ModifyServerlessAuthToMachineRequest) GetAuthItem() *string {
	return s.AuthItem
}

func (s *ModifyServerlessAuthToMachineRequest) GetAutoBind() *int32 {
	return s.AutoBind
}

func (s *ModifyServerlessAuthToMachineRequest) GetBindAll() *bool {
	return s.BindAll
}

func (s *ModifyServerlessAuthToMachineRequest) GetBindAppList() []*string {
	return s.BindAppList
}

func (s *ModifyServerlessAuthToMachineRequest) GetBindAssetType() *string {
	return s.BindAssetType
}

func (s *ModifyServerlessAuthToMachineRequest) GetBindUuidList() []*string {
	return s.BindUuidList
}

func (s *ModifyServerlessAuthToMachineRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyServerlessAuthToMachineRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *ModifyServerlessAuthToMachineRequest) GetLogicalExp() *string {
	return s.LogicalExp
}

func (s *ModifyServerlessAuthToMachineRequest) GetNtmVersion() *string {
	return s.NtmVersion
}

func (s *ModifyServerlessAuthToMachineRequest) GetPreBind() *int32 {
	return s.PreBind
}

func (s *ModifyServerlessAuthToMachineRequest) GetPreBindOrderId() *int64 {
	return s.PreBindOrderId
}

func (s *ModifyServerlessAuthToMachineRequest) GetResourceDirectoryUid() *int64 {
	return s.ResourceDirectoryUid
}

func (s *ModifyServerlessAuthToMachineRequest) GetUnBindAppList() []*string {
	return s.UnBindAppList
}

func (s *ModifyServerlessAuthToMachineRequest) GetUnBindUuidList() []*string {
	return s.UnBindUuidList
}

func (s *ModifyServerlessAuthToMachineRequest) SetAppCriteria(v string) *ModifyServerlessAuthToMachineRequest {
	s.AppCriteria = &v
	return s
}

func (s *ModifyServerlessAuthToMachineRequest) SetAuthItem(v string) *ModifyServerlessAuthToMachineRequest {
	s.AuthItem = &v
	return s
}

func (s *ModifyServerlessAuthToMachineRequest) SetAutoBind(v int32) *ModifyServerlessAuthToMachineRequest {
	s.AutoBind = &v
	return s
}

func (s *ModifyServerlessAuthToMachineRequest) SetBindAll(v bool) *ModifyServerlessAuthToMachineRequest {
	s.BindAll = &v
	return s
}

func (s *ModifyServerlessAuthToMachineRequest) SetBindAppList(v []*string) *ModifyServerlessAuthToMachineRequest {
	s.BindAppList = v
	return s
}

func (s *ModifyServerlessAuthToMachineRequest) SetBindAssetType(v string) *ModifyServerlessAuthToMachineRequest {
	s.BindAssetType = &v
	return s
}

func (s *ModifyServerlessAuthToMachineRequest) SetBindUuidList(v []*string) *ModifyServerlessAuthToMachineRequest {
	s.BindUuidList = v
	return s
}

func (s *ModifyServerlessAuthToMachineRequest) SetClientToken(v string) *ModifyServerlessAuthToMachineRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyServerlessAuthToMachineRequest) SetCriteria(v string) *ModifyServerlessAuthToMachineRequest {
	s.Criteria = &v
	return s
}

func (s *ModifyServerlessAuthToMachineRequest) SetLogicalExp(v string) *ModifyServerlessAuthToMachineRequest {
	s.LogicalExp = &v
	return s
}

func (s *ModifyServerlessAuthToMachineRequest) SetNtmVersion(v string) *ModifyServerlessAuthToMachineRequest {
	s.NtmVersion = &v
	return s
}

func (s *ModifyServerlessAuthToMachineRequest) SetPreBind(v int32) *ModifyServerlessAuthToMachineRequest {
	s.PreBind = &v
	return s
}

func (s *ModifyServerlessAuthToMachineRequest) SetPreBindOrderId(v int64) *ModifyServerlessAuthToMachineRequest {
	s.PreBindOrderId = &v
	return s
}

func (s *ModifyServerlessAuthToMachineRequest) SetResourceDirectoryUid(v int64) *ModifyServerlessAuthToMachineRequest {
	s.ResourceDirectoryUid = &v
	return s
}

func (s *ModifyServerlessAuthToMachineRequest) SetUnBindAppList(v []*string) *ModifyServerlessAuthToMachineRequest {
	s.UnBindAppList = v
	return s
}

func (s *ModifyServerlessAuthToMachineRequest) SetUnBindUuidList(v []*string) *ModifyServerlessAuthToMachineRequest {
	s.UnBindUuidList = v
	return s
}

func (s *ModifyServerlessAuthToMachineRequest) Validate() error {
	return dara.Validate(s)
}
