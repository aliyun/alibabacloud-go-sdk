// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudPermissionTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAccountId(v string) *DeleteCloudPermissionTemplateRequest
	GetCloudAccountId() *string
	SetCloudPermissionTemplateId(v string) *DeleteCloudPermissionTemplateRequest
	GetCloudPermissionTemplateId() *string
	SetInstanceId(v string) *DeleteCloudPermissionTemplateRequest
	GetInstanceId() *string
}

type DeleteCloudPermissionTemplateRequest struct {
	// The cloud account ID.
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

func (s DeleteCloudPermissionTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudPermissionTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteCloudPermissionTemplateRequest) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *DeleteCloudPermissionTemplateRequest) GetCloudPermissionTemplateId() *string {
	return s.CloudPermissionTemplateId
}

func (s *DeleteCloudPermissionTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteCloudPermissionTemplateRequest) SetCloudAccountId(v string) *DeleteCloudPermissionTemplateRequest {
	s.CloudAccountId = &v
	return s
}

func (s *DeleteCloudPermissionTemplateRequest) SetCloudPermissionTemplateId(v string) *DeleteCloudPermissionTemplateRequest {
	s.CloudPermissionTemplateId = &v
	return s
}

func (s *DeleteCloudPermissionTemplateRequest) SetInstanceId(v string) *DeleteCloudPermissionTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteCloudPermissionTemplateRequest) Validate() error {
	return dara.Validate(s)
}
