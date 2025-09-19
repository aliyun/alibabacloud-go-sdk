// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVulAutoRepairConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReason(v string) *CreateVulAutoRepairConfigRequest
	GetReason() *string
	SetType(v string) *CreateVulAutoRepairConfigRequest
	GetType() *string
	SetVulAutoRepairConfigList(v []*CreateVulAutoRepairConfigRequestVulAutoRepairConfigList) *CreateVulAutoRepairConfigRequest
	GetVulAutoRepairConfigList() []*CreateVulAutoRepairConfigRequestVulAutoRepairConfigList
}

type CreateVulAutoRepairConfigRequest struct {
	// The reason why the vulnerability can be automatically fixed.
	//
	// example:
	//
	// TestAutoRepair
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The type of the vulnerability. Valid values: -**cve**: Linux software vulnerability -**sys**: Windows system vulnerability
	//
	// This parameter is required.
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The vulnerabilities that can be automatically fixed.
	//
	// This parameter is required.
	VulAutoRepairConfigList []*CreateVulAutoRepairConfigRequestVulAutoRepairConfigList `json:"VulAutoRepairConfigList,omitempty" xml:"VulAutoRepairConfigList,omitempty" type:"Repeated"`
}

func (s CreateVulAutoRepairConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVulAutoRepairConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateVulAutoRepairConfigRequest) GetReason() *string {
	return s.Reason
}

func (s *CreateVulAutoRepairConfigRequest) GetType() *string {
	return s.Type
}

func (s *CreateVulAutoRepairConfigRequest) GetVulAutoRepairConfigList() []*CreateVulAutoRepairConfigRequestVulAutoRepairConfigList {
	return s.VulAutoRepairConfigList
}

func (s *CreateVulAutoRepairConfigRequest) SetReason(v string) *CreateVulAutoRepairConfigRequest {
	s.Reason = &v
	return s
}

func (s *CreateVulAutoRepairConfigRequest) SetType(v string) *CreateVulAutoRepairConfigRequest {
	s.Type = &v
	return s
}

func (s *CreateVulAutoRepairConfigRequest) SetVulAutoRepairConfigList(v []*CreateVulAutoRepairConfigRequestVulAutoRepairConfigList) *CreateVulAutoRepairConfigRequest {
	s.VulAutoRepairConfigList = v
	return s
}

func (s *CreateVulAutoRepairConfigRequest) Validate() error {
	return dara.Validate(s)
}

type CreateVulAutoRepairConfigRequestVulAutoRepairConfigList struct {
	// The alias of the vulnerability.
	//
	// This parameter is required.
	//
	// example:
	//
	// CVE-2018-25032:zlib 1.2.11 memory corruption
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The name of the vulnerability.
	//
	// This parameter is required.
	//
	// example:
	//
	// anolisos:8.4:ANSA-2022:0001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateVulAutoRepairConfigRequestVulAutoRepairConfigList) String() string {
	return dara.Prettify(s)
}

func (s CreateVulAutoRepairConfigRequestVulAutoRepairConfigList) GoString() string {
	return s.String()
}

func (s *CreateVulAutoRepairConfigRequestVulAutoRepairConfigList) GetAliasName() *string {
	return s.AliasName
}

func (s *CreateVulAutoRepairConfigRequestVulAutoRepairConfigList) GetName() *string {
	return s.Name
}

func (s *CreateVulAutoRepairConfigRequestVulAutoRepairConfigList) SetAliasName(v string) *CreateVulAutoRepairConfigRequestVulAutoRepairConfigList {
	s.AliasName = &v
	return s
}

func (s *CreateVulAutoRepairConfigRequestVulAutoRepairConfigList) SetName(v string) *CreateVulAutoRepairConfigRequestVulAutoRepairConfigList {
	s.Name = &v
	return s
}

func (s *CreateVulAutoRepairConfigRequestVulAutoRepairConfigList) Validate() error {
	return dara.Validate(s)
}
