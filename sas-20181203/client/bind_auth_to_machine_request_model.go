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
	SetUnBind(v []*string) *BindAuthToMachineRequest
	GetUnBind() []*string
}

type BindAuthToMachineRequest struct {
	// The edition of Security Center that is authorized to scan the asset. Valid values:
	//
	// 	- **6**: Anti-virus
	//
	// 	- **5**: Advanced
	//
	// 	- **3**: Enterprise
	//
	// 	- **7**: Ultimate
	//
	// 	- **10**: Value-added Plan
	//
	// example:
	//
	// 6
	AuthVersion *int32 `json:"AuthVersion,omitempty" xml:"AuthVersion,omitempty"`
	// Specifies whether to automatically bind servers to Security Center. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 1
	AutoBind *int32 `json:"AutoBind,omitempty" xml:"AutoBind,omitempty"`
	// The UUIDs of the servers that you want to bind to Security Center.
	//
	// >  You must specify at least one of the **Bind*	- and **UnBind*	- parameters.
	Bind []*string `json:"Bind,omitempty" xml:"Bind,omitempty" type:"Repeated"`
	// Specifies whether to bind all servers to Security Center. Default value: **false**. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	BindAll *bool `json:"BindAll,omitempty" xml:"BindAll,omitempty"`
	// The search conditions that are used to filter servers. The value of this parameter is in the JSON format and is case-sensitive.
	//
	// >  A search condition can be an instance ID, instance name, virtual private cloud (VPC) ID, region, or public IP address. You can call the [DescribeCriteria](~~DescribeCriteria~~) operation to query the supported search conditions.
	//
	// example:
	//
	// [{"name":"riskStatus","value":"YES"},{"name":"internetIp","value":"1.2.XX.XX"}]
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// Specifies whether to specify servers for protection when you purchase Security Center. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// >  If you specify servers, the servers are automatically added to Security Center for protection after the purchase order is complete.
	//
	// example:
	//
	// 1
	IsPreBind *int32 `json:"IsPreBind,omitempty" xml:"IsPreBind,omitempty"`
	// The logical relationship that you want to use to evaluate multiple search conditions. Default value: **OR**. Valid values:
	//
	// 	- **OR**
	//
	// 	- **AND**
	//
	// example:
	//
	// OR
	LogicalExp *string `json:"LogicalExp,omitempty" xml:"LogicalExp,omitempty"`
	// The edition of Security Center that you purchase in the order. Valid values:
	//
	// 	- **level7**: Anti-virus
	//
	// 	- **level3**: Advanced
	//
	// 	- **level2**: Enterprise
	//
	// 	- **level8**: Ultimate
	//
	// 	- **level10**: Value-added Plan
	//
	// example:
	//
	// level2
	NtmVersion *string `json:"NtmVersion,omitempty" xml:"NtmVersion,omitempty"`
	// The ID of the order in which Security Center is purchased and servers are specified for protection.
	//
	// example:
	//
	// 233016**0482
	PreBindOrderId *int64 `json:"PreBindOrderId,omitempty" xml:"PreBindOrderId,omitempty"`
	// The UUIDs of the servers that you want to unbind from Security Center.
	//
	// >  You must specify at least one of the **Bind*	- and **UnBind*	- parameters.
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

func (s *BindAuthToMachineRequest) SetUnBind(v []*string) *BindAuthToMachineRequest {
	s.UnBind = v
	return s
}

func (s *BindAuthToMachineRequest) Validate() error {
	return dara.Validate(s)
}
