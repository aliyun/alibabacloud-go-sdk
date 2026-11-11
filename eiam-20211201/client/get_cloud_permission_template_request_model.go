// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudPermissionTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAccountId(v string) *GetCloudPermissionTemplateRequest
	GetCloudAccountId() *string
	SetCloudPermissionTemplateId(v string) *GetCloudPermissionTemplateRequest
	GetCloudPermissionTemplateId() *string
	SetInstanceId(v string) *GetCloudPermissionTemplateRequest
	GetInstanceId() *string
}

type GetCloudPermissionTemplateRequest struct {
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

func (s GetCloudPermissionTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCloudPermissionTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetCloudPermissionTemplateRequest) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *GetCloudPermissionTemplateRequest) GetCloudPermissionTemplateId() *string {
	return s.CloudPermissionTemplateId
}

func (s *GetCloudPermissionTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCloudPermissionTemplateRequest) SetCloudAccountId(v string) *GetCloudPermissionTemplateRequest {
	s.CloudAccountId = &v
	return s
}

func (s *GetCloudPermissionTemplateRequest) SetCloudPermissionTemplateId(v string) *GetCloudPermissionTemplateRequest {
	s.CloudPermissionTemplateId = &v
	return s
}

func (s *GetCloudPermissionTemplateRequest) SetInstanceId(v string) *GetCloudPermissionTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCloudPermissionTemplateRequest) Validate() error {
	return dara.Validate(s)
}
