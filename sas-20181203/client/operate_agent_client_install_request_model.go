// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateAgentClientInstallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v string) *OperateAgentClientInstallRequest
	GetInstanceIds() *string
	SetLang(v string) *OperateAgentClientInstallRequest
	GetLang() *string
	SetOs(v string) *OperateAgentClientInstallRequest
	GetOs() *string
	SetRegion(v string) *OperateAgentClientInstallRequest
	GetRegion() *string
	SetUuids(v string) *OperateAgentClientInstallRequest
	GetUuids() *string
}

type OperateAgentClientInstallRequest struct {
	// The IDs of the servers on which you want to install the Security Center agent. Separate multiple IDs with commas (,).
	//
	// > : You must specify at least one of **InstanceIds*	- and **Uuids**. If you specify **InstanceIds**, you must also specify **Region*	- and **Os**.
	//
	// example:
	//
	// i-uf6j8vq9l4r5ntht****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The operating system of the servers. Valid values:
	//
	// 	- **linux**
	//
	// 	- **windows**
	//
	// example:
	//
	// linux
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The region where the servers reside. Valid values include the following regions:
	//
	// 	- cn-hangzhou: China (Hangzhou)
	//
	// 	- cn-beijing: China (Beijing)
	//
	// 	- cn-shanghai: China (Shanghai)
	//
	// 	- cn-zhangjiakou: China (Zhangjiakou)
	//
	// 	- cn-shenzhen: China (Shenzhen)
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The UUIDs of the servers on which you want to install the Security Center agent. Separate multiple UUIDs with commas (,).
	//
	// > You must specify at least one of the **InstanceIds*	- and **Uuids*	- parameters before you can call this operation.
	//
	// example:
	//
	// 1587bedb-fdb4-48c4-9330-************
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s OperateAgentClientInstallRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateAgentClientInstallRequest) GoString() string {
	return s.String()
}

func (s *OperateAgentClientInstallRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *OperateAgentClientInstallRequest) GetLang() *string {
	return s.Lang
}

func (s *OperateAgentClientInstallRequest) GetOs() *string {
	return s.Os
}

func (s *OperateAgentClientInstallRequest) GetRegion() *string {
	return s.Region
}

func (s *OperateAgentClientInstallRequest) GetUuids() *string {
	return s.Uuids
}

func (s *OperateAgentClientInstallRequest) SetInstanceIds(v string) *OperateAgentClientInstallRequest {
	s.InstanceIds = &v
	return s
}

func (s *OperateAgentClientInstallRequest) SetLang(v string) *OperateAgentClientInstallRequest {
	s.Lang = &v
	return s
}

func (s *OperateAgentClientInstallRequest) SetOs(v string) *OperateAgentClientInstallRequest {
	s.Os = &v
	return s
}

func (s *OperateAgentClientInstallRequest) SetRegion(v string) *OperateAgentClientInstallRequest {
	s.Region = &v
	return s
}

func (s *OperateAgentClientInstallRequest) SetUuids(v string) *OperateAgentClientInstallRequest {
	s.Uuids = &v
	return s
}

func (s *OperateAgentClientInstallRequest) Validate() error {
	return dara.Validate(s)
}
