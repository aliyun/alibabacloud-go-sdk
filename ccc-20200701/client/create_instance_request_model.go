// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdminRamIdList(v string) *CreateInstanceRequest
	GetAdminRamIdList() *string
	SetDescription(v string) *CreateInstanceRequest
	GetDescription() *string
	SetDomainName(v string) *CreateInstanceRequest
	GetDomainName() *string
	SetName(v string) *CreateInstanceRequest
	GetName() *string
	SetNumberList(v string) *CreateInstanceRequest
	GetNumberList() *string
}

type CreateInstanceRequest struct {
	// List of RAM IDs to be assigned as administrators of this Cloud Contact Center instance. After successful creation, the RAM IDs in this list will be automatically imported into the newly created instance with administrator privileges.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["24861380681070****","105980354482****"]
	AdminRamIdList *string `json:"AdminRamIdList,omitempty" xml:"AdminRamIdList,omitempty"`
	// Description of the instance. The length must not exceed 256 characters.
	//
	// example:
	//
	// 云联络中心的测试实例。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Second-level domain name of the Cloud Contact Center instance, which serves as the instance ID and is globally unique. It must be 4 to 48 characters long and can only contain uppercase and lowercase English letters, digits, underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Instance name. The length must be between 4 and 32 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试实例
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// List of phone numbers for the Cloud Contact Center instance to be created. The current account must have usage rights to these numbers, and the numbers must not be associated with any other instance.
	//
	// example:
	//
	// ["0830011xxxx", "0830312xxxx"]
	NumberList *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetAdminRamIdList() *string {
	return s.AdminRamIdList
}

func (s *CreateInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateInstanceRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateInstanceRequest) GetName() *string {
	return s.Name
}

func (s *CreateInstanceRequest) GetNumberList() *string {
	return s.NumberList
}

func (s *CreateInstanceRequest) SetAdminRamIdList(v string) *CreateInstanceRequest {
	s.AdminRamIdList = &v
	return s
}

func (s *CreateInstanceRequest) SetDescription(v string) *CreateInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateInstanceRequest) SetDomainName(v string) *CreateInstanceRequest {
	s.DomainName = &v
	return s
}

func (s *CreateInstanceRequest) SetName(v string) *CreateInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateInstanceRequest) SetNumberList(v string) *CreateInstanceRequest {
	s.NumberList = &v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
	return dara.Validate(s)
}
