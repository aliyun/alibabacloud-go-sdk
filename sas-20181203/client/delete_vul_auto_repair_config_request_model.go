// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVulAutoRepairConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliasName(v string) *DeleteVulAutoRepairConfigRequest
	GetAliasName() *string
	SetConfigIdList(v []*int64) *DeleteVulAutoRepairConfigRequest
	GetConfigIdList() []*int64
	SetType(v string) *DeleteVulAutoRepairConfigRequest
	GetType() *string
}

type DeleteVulAutoRepairConfigRequest struct {
	// The alias of the vulnerability.
	//
	// example:
	//
	// RHSA-2017:0184-Important: mysql security update
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The IDs of the configurations.
	//
	// >  You can call the [ListVulAutoRepairConfig](~~ListVulAutoRepairConfig~~) operation to query the IDs.
	ConfigIdList []*int64 `json:"ConfigIdList,omitempty" xml:"ConfigIdList,omitempty" type:"Repeated"`
	// The type of the vulnerability. Valid values:
	//
	// 	- cve: Linux software vulnerability
	//
	// 	- sys: Windows system vulnerability
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeleteVulAutoRepairConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVulAutoRepairConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteVulAutoRepairConfigRequest) GetAliasName() *string {
	return s.AliasName
}

func (s *DeleteVulAutoRepairConfigRequest) GetConfigIdList() []*int64 {
	return s.ConfigIdList
}

func (s *DeleteVulAutoRepairConfigRequest) GetType() *string {
	return s.Type
}

func (s *DeleteVulAutoRepairConfigRequest) SetAliasName(v string) *DeleteVulAutoRepairConfigRequest {
	s.AliasName = &v
	return s
}

func (s *DeleteVulAutoRepairConfigRequest) SetConfigIdList(v []*int64) *DeleteVulAutoRepairConfigRequest {
	s.ConfigIdList = v
	return s
}

func (s *DeleteVulAutoRepairConfigRequest) SetType(v string) *DeleteVulAutoRepairConfigRequest {
	s.Type = &v
	return s
}

func (s *DeleteVulAutoRepairConfigRequest) Validate() error {
	return dara.Validate(s)
}
