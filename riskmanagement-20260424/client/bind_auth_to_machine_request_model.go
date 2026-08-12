// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAuthToMachineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *BindAuthToMachineRequest
	GetRegionId() *string
	SetSdkRequest(v *BindAuthToMachineRequestSdkRequest) *BindAuthToMachineRequest
	GetSdkRequest() *BindAuthToMachineRequestSdkRequest
}

type BindAuthToMachineRequest struct {
	// The region ID of the Smart Access Gateway instance.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Security Center SDK request.
	SdkRequest *BindAuthToMachineRequestSdkRequest `json:"SdkRequest,omitempty" xml:"SdkRequest,omitempty" type:"Struct"`
}

func (s BindAuthToMachineRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAuthToMachineRequest) GoString() string {
	return s.String()
}

func (s *BindAuthToMachineRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BindAuthToMachineRequest) GetSdkRequest() *BindAuthToMachineRequestSdkRequest {
	return s.SdkRequest
}

func (s *BindAuthToMachineRequest) SetRegionId(v string) *BindAuthToMachineRequest {
	s.RegionId = &v
	return s
}

func (s *BindAuthToMachineRequest) SetSdkRequest(v *BindAuthToMachineRequestSdkRequest) *BindAuthToMachineRequest {
	s.SdkRequest = v
	return s
}

func (s *BindAuthToMachineRequest) Validate() error {
	if s.SdkRequest != nil {
		if err := s.SdkRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindAuthToMachineRequestSdkRequest struct {
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
	// - **10**: Value-added Service Edition
	//
	// example:
	//
	// 3
	AuthVersion *int32 `json:"AuthVersion,omitempty" xml:"AuthVersion,omitempty"`
	// Specifies whether to enable automatic binding. Valid values:
	//
	// - **0**: disabled
	//
	// - **1**: enabled
	//
	// example:
	//
	// 1
	AutoBind *int32 `json:"AutoBind,omitempty" xml:"AutoBind,omitempty"`
	// The collection of UUIDs to bind.
	//
	// > Bind and UnBind cannot both be empty.
	//
	// Maximum number of child entries: 1000.
	Bind []*string `json:"Bind,omitempty" xml:"Bind,omitempty" type:"Repeated"`
	// Specifies whether to bind all assets. Default value: **false**. Valid values:
	//
	// - **true**: yes
	//
	// - **false**: no
	//
	// example:
	//
	// true
	BindAll *bool `json:"BindAll,omitempty" xml:"BindAll,omitempty"`
	// The search conditions for assets. This parameter is in JSON format. Pay attention to the letter case when you specify this parameter.
	//
	// > You can search for assets by instance ID, instance name, VPC ID, region, public IP address, and other conditions. You can call the DescribeCriteria operation to query the supported search conditions.
	//
	// example:
	//
	// [{\\"name\\":\\"clientStatus\\",\\"value\\":\\"online\\"},{\\"name\\":\\"authVersion\\",\\"value\\":\\"1\\"}]
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// Specifies whether this is a pre-binding operation. Valid values:
	//
	// - **0**: no
	//
	// - **1**: yes
	//
	// > After pre-binding is enabled, the corresponding authorization quota is automatically bound to the specified servers after the purchase is completed.
	//
	// example:
	//
	// 1
	IsPreBind *int32 `json:"IsPreBind,omitempty" xml:"IsPreBind,omitempty"`
	// The logical relationship between multiple search conditions. Valid values:
	//
	// - **OR**: The search conditions are in an **OR*	- relationship.
	//
	// - **AND**: The search conditions are in an **AND*	- relationship.
	//
	// example:
	//
	// AND
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
	// - **level10**: value-added service only
	//
	// example:
	//
	// level2
	NtmVersion *int64 `json:"NtmVersion,omitempty" xml:"NtmVersion,omitempty"`
	// The order ID associated with the pre-binding.
	//
	// > Note: This field is of the Long type. Precision loss may occur during the sequence/deserialization procedure. The value must not exceed 9007199254740991.
	//
	// example:
	//
	// 263076506250432
	PreBindOrderId *int64 `json:"PreBindOrderId,omitempty" xml:"PreBindOrderId,omitempty"`
	// The collection of UUIDs to unbind.
	//
	// > **Bind*	- and **UnBind*	- cannot both be empty.
	UnBind []*string `json:"UnBind,omitempty" xml:"UnBind,omitempty" type:"Repeated"`
}

func (s BindAuthToMachineRequestSdkRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAuthToMachineRequestSdkRequest) GoString() string {
	return s.String()
}

func (s *BindAuthToMachineRequestSdkRequest) GetAuthVersion() *int32 {
	return s.AuthVersion
}

func (s *BindAuthToMachineRequestSdkRequest) GetAutoBind() *int32 {
	return s.AutoBind
}

func (s *BindAuthToMachineRequestSdkRequest) GetBind() []*string {
	return s.Bind
}

func (s *BindAuthToMachineRequestSdkRequest) GetBindAll() *bool {
	return s.BindAll
}

func (s *BindAuthToMachineRequestSdkRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *BindAuthToMachineRequestSdkRequest) GetIsPreBind() *int32 {
	return s.IsPreBind
}

func (s *BindAuthToMachineRequestSdkRequest) GetLogicalExp() *string {
	return s.LogicalExp
}

func (s *BindAuthToMachineRequestSdkRequest) GetNtmVersion() *int64 {
	return s.NtmVersion
}

func (s *BindAuthToMachineRequestSdkRequest) GetPreBindOrderId() *int64 {
	return s.PreBindOrderId
}

func (s *BindAuthToMachineRequestSdkRequest) GetUnBind() []*string {
	return s.UnBind
}

func (s *BindAuthToMachineRequestSdkRequest) SetAuthVersion(v int32) *BindAuthToMachineRequestSdkRequest {
	s.AuthVersion = &v
	return s
}

func (s *BindAuthToMachineRequestSdkRequest) SetAutoBind(v int32) *BindAuthToMachineRequestSdkRequest {
	s.AutoBind = &v
	return s
}

func (s *BindAuthToMachineRequestSdkRequest) SetBind(v []*string) *BindAuthToMachineRequestSdkRequest {
	s.Bind = v
	return s
}

func (s *BindAuthToMachineRequestSdkRequest) SetBindAll(v bool) *BindAuthToMachineRequestSdkRequest {
	s.BindAll = &v
	return s
}

func (s *BindAuthToMachineRequestSdkRequest) SetCriteria(v string) *BindAuthToMachineRequestSdkRequest {
	s.Criteria = &v
	return s
}

func (s *BindAuthToMachineRequestSdkRequest) SetIsPreBind(v int32) *BindAuthToMachineRequestSdkRequest {
	s.IsPreBind = &v
	return s
}

func (s *BindAuthToMachineRequestSdkRequest) SetLogicalExp(v string) *BindAuthToMachineRequestSdkRequest {
	s.LogicalExp = &v
	return s
}

func (s *BindAuthToMachineRequestSdkRequest) SetNtmVersion(v int64) *BindAuthToMachineRequestSdkRequest {
	s.NtmVersion = &v
	return s
}

func (s *BindAuthToMachineRequestSdkRequest) SetPreBindOrderId(v int64) *BindAuthToMachineRequestSdkRequest {
	s.PreBindOrderId = &v
	return s
}

func (s *BindAuthToMachineRequestSdkRequest) SetUnBind(v []*string) *BindAuthToMachineRequestSdkRequest {
	s.UnBind = v
	return s
}

func (s *BindAuthToMachineRequestSdkRequest) Validate() error {
	return dara.Validate(s)
}
