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
	// Application query condition.
	//
	// example:
	//
	// **7ad7e3a
	AppCriteria *string `json:"AppCriteria,omitempty" xml:"AppCriteria,omitempty"`
	// Instance type. Values:
	//
	// - **SERVERLESS**: Serverless asset
	//
	// example:
	//
	// SERVERLESS
	AuthItem *string `json:"AuthItem,omitempty" xml:"AuthItem,omitempty"`
	// Enable auto-binding. Values:
	//
	// - **0**: Off
	//
	// - **1**: On
	//
	// example:
	//
	// 1
	AutoBind *int32 `json:"AutoBind,omitempty" xml:"AutoBind,omitempty"`
	// Whether to bind all. Default is **false**. Values:
	//
	// - **true**: Yes
	//
	// - **false**: No
	//
	// example:
	//
	// false
	BindAll *bool `json:"BindAll,omitempty" xml:"BindAll,omitempty"`
	// List of application IDs to be bound.
	//
	// > Obtained through the [ListMachineApps](~~ListMachineApps~~) interface.
	BindAppList []*string `json:"BindAppList,omitempty" xml:"BindAppList,omitempty" type:"Repeated"`
	// Type of asset to operate on. Values:
	//
	// - **INSTANCE**: Instance
	//
	// - **APP**: Application
	//
	// example:
	//
	// APP
	BindAssetType *string `json:"BindAssetType,omitempty" xml:"BindAssetType,omitempty"`
	// List of asset UUIDs to be bound.
	BindUuidList []*string `json:"BindUuidList,omitempty" xml:"BindUuidList,omitempty" type:"Repeated"`
	// Set the conditions for searching assets. This parameter is in JSON format, and case sensitivity should be noted when entering parameters.
	//
	// > Supports searching assets using instance ID, instance name, VPC ID, region, public IP address, etc. You can call the [DescribeCriteria](~~DescribeCriteria~~) interface to query supported search conditions.
	//
	// example:
	//
	// [{"name":"vulStatus","value":"YES","logicalExp":"AND"}]
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// Set the logical relationship between multiple search conditions. Values:
	//
	// - **OR**: Indicates an **or*	- relationship between multiple conditions.
	//
	// - **AND**: Indicates an **and*	- relationship between multiple conditions.
	//
	// example:
	//
	// OR
	LogicalExp *string `json:"LogicalExp,omitempty" xml:"LogicalExp,omitempty"`
	// NTM version code, used for pre-binding.
	//
	// example:
	//
	// level2
	NtmVersion *string `json:"NtmVersion,omitempty" xml:"NtmVersion,omitempty"`
	// Whether it is a pre-bind operation. Values:
	//
	// - **0**: No
	//
	// - **1**: Yes
	//
	//
	// > After enabling pre-binding, the specified server will automatically bind the corresponding version\\"s authorization count after the purchase is completed.
	//
	// example:
	//
	// 1
	PreBind *int32 `json:"PreBind,omitempty" xml:"PreBind,omitempty"`
	// Pre-bind order ID.
	//
	// example:
	//
	// 233016**0482
	PreBindOrderId *int64 `json:"PreBindOrderId,omitempty" xml:"PreBindOrderId,omitempty"`
	// UID of the associated resource directory.
	//
	// example:
	//
	// 123456
	ResourceDirectoryUid *int64 `json:"ResourceDirectoryUid,omitempty" xml:"ResourceDirectoryUid,omitempty"`
	// List of application IDs to be unbound.
	//
	// > Obtained through the [ListMachineApps](~~ListMachineApps~~) interface.
	UnBindAppList []*string `json:"UnBindAppList,omitempty" xml:"UnBindAppList,omitempty" type:"Repeated"`
	// List of asset UUIDs to be unbound.
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
