// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCloudPermissionTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAccountId(v string) *DisableCloudPermissionTemplateRequest
	GetCloudAccountId() *string
	SetCloudPermissionTemplateId(v string) *DisableCloudPermissionTemplateRequest
	GetCloudPermissionTemplateId() *string
	SetInstanceId(v string) *DisableCloudPermissionTemplateRequest
	GetInstanceId() *string
}

type DisableCloudPermissionTemplateRequest struct {
	// The ID of the cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac_mkv7rgt4d7i4u7zqtzev2mxxxx
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// The ID of the permission template.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpt_mkv7rgt4d7i4u7zqtzev2mxxxx
	CloudPermissionTemplateId *string `json:"CloudPermissionTemplateId,omitempty" xml:"CloudPermissionTemplateId,omitempty"`
	// The ID of the IDaaS EIAM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableCloudPermissionTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableCloudPermissionTemplateRequest) GoString() string {
	return s.String()
}

func (s *DisableCloudPermissionTemplateRequest) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *DisableCloudPermissionTemplateRequest) GetCloudPermissionTemplateId() *string {
	return s.CloudPermissionTemplateId
}

func (s *DisableCloudPermissionTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableCloudPermissionTemplateRequest) SetCloudAccountId(v string) *DisableCloudPermissionTemplateRequest {
	s.CloudAccountId = &v
	return s
}

func (s *DisableCloudPermissionTemplateRequest) SetCloudPermissionTemplateId(v string) *DisableCloudPermissionTemplateRequest {
	s.CloudPermissionTemplateId = &v
	return s
}

func (s *DisableCloudPermissionTemplateRequest) SetInstanceId(v string) *DisableCloudPermissionTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableCloudPermissionTemplateRequest) Validate() error {
	return dara.Validate(s)
}
