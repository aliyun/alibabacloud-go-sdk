// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateVulsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperateType(v string) *OperateVulsRequest
	GetOperateType() *string
	SetType(v string) *OperateVulsRequest
	GetType() *string
	SetUuids(v []*string) *OperateVulsRequest
	GetUuids() []*string
	SetVulNames(v []*string) *OperateVulsRequest
	GetVulNames() []*string
}

type OperateVulsRequest struct {
	// The method to handle the vulnerability. Set the value to **vul_fix**, which indicates fixing the vulnerability.
	//
	// This parameter is required.
	//
	// example:
	//
	// vul_fix
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// The type of vulnerability to fix. Set the value to **cve**, which indicates a Linux software vulnerability.
	//
	// This parameter is required.
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The UUIDs of the servers on which you want to fix the vulnerabilities.
	//
	// This parameter is required.
	Uuids []*string `json:"Uuids,omitempty" xml:"Uuids,omitempty" type:"Repeated"`
	// The names of the vulnerabilities to fix.
	//
	// This parameter is required.
	VulNames []*string `json:"VulNames,omitempty" xml:"VulNames,omitempty" type:"Repeated"`
}

func (s OperateVulsRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateVulsRequest) GoString() string {
	return s.String()
}

func (s *OperateVulsRequest) GetOperateType() *string {
	return s.OperateType
}

func (s *OperateVulsRequest) GetType() *string {
	return s.Type
}

func (s *OperateVulsRequest) GetUuids() []*string {
	return s.Uuids
}

func (s *OperateVulsRequest) GetVulNames() []*string {
	return s.VulNames
}

func (s *OperateVulsRequest) SetOperateType(v string) *OperateVulsRequest {
	s.OperateType = &v
	return s
}

func (s *OperateVulsRequest) SetType(v string) *OperateVulsRequest {
	s.Type = &v
	return s
}

func (s *OperateVulsRequest) SetUuids(v []*string) *OperateVulsRequest {
	s.Uuids = v
	return s
}

func (s *OperateVulsRequest) SetVulNames(v []*string) *OperateVulsRequest {
	s.VulNames = v
	return s
}

func (s *OperateVulsRequest) Validate() error {
	return dara.Validate(s)
}
