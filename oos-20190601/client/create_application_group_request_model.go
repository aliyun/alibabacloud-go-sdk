// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationName(v string) *CreateApplicationGroupRequest
	GetApplicationName() *string
	SetClientToken(v string) *CreateApplicationGroupRequest
	GetClientToken() *string
	SetCmsGroupId(v string) *CreateApplicationGroupRequest
	GetCmsGroupId() *string
	SetDeployRegionId(v string) *CreateApplicationGroupRequest
	GetDeployRegionId() *string
	SetDescription(v string) *CreateApplicationGroupRequest
	GetDescription() *string
	SetImportTagKey(v string) *CreateApplicationGroupRequest
	GetImportTagKey() *string
	SetImportTagValue(v string) *CreateApplicationGroupRequest
	GetImportTagValue() *string
	SetName(v string) *CreateApplicationGroupRequest
	GetName() *string
	SetRegionId(v string) *CreateApplicationGroupRequest
	GetRegionId() *string
}

type CreateApplicationGroupRequest struct {
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplication
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// -
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the application group in CloudMonitor.
	//
	// example:
	//
	// 218026174
	CmsGroupId *string `json:"CmsGroupId,omitempty" xml:"CmsGroupId,omitempty"`
	// The ID of the region in which the related sources reside.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DeployRegionId *string `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
	// The description of the application group.
	//
	// example:
	//
	// ApplicationGroup
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The key of the tag. You must set both the ImportTagKey and the ImportTagValue parameters, or leave both of them empty. If you do not set the ImportTagKey and ImportTagValue parameters, the application name is used for this parameter by default.
	//
	// example:
	//
	// k1
	ImportTagKey *string `json:"ImportTagKey,omitempty" xml:"ImportTagKey,omitempty"`
	// The value of the tag. You must set both the ImportTagKey and the ImportTagValue parameters, or leave both of them empty. If you do not set the ImportTagKey and ImportTagValue parameters, the application group name is used for this parameter by default.
	//
	// example:
	//
	// v1
	ImportTagValue *string `json:"ImportTagValue,omitempty" xml:"ImportTagValue,omitempty"`
	// The name of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplicationGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateApplicationGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationGroupRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *CreateApplicationGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateApplicationGroupRequest) GetCmsGroupId() *string {
	return s.CmsGroupId
}

func (s *CreateApplicationGroupRequest) GetDeployRegionId() *string {
	return s.DeployRegionId
}

func (s *CreateApplicationGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApplicationGroupRequest) GetImportTagKey() *string {
	return s.ImportTagKey
}

func (s *CreateApplicationGroupRequest) GetImportTagValue() *string {
	return s.ImportTagValue
}

func (s *CreateApplicationGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateApplicationGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApplicationGroupRequest) SetApplicationName(v string) *CreateApplicationGroupRequest {
	s.ApplicationName = &v
	return s
}

func (s *CreateApplicationGroupRequest) SetClientToken(v string) *CreateApplicationGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateApplicationGroupRequest) SetCmsGroupId(v string) *CreateApplicationGroupRequest {
	s.CmsGroupId = &v
	return s
}

func (s *CreateApplicationGroupRequest) SetDeployRegionId(v string) *CreateApplicationGroupRequest {
	s.DeployRegionId = &v
	return s
}

func (s *CreateApplicationGroupRequest) SetDescription(v string) *CreateApplicationGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateApplicationGroupRequest) SetImportTagKey(v string) *CreateApplicationGroupRequest {
	s.ImportTagKey = &v
	return s
}

func (s *CreateApplicationGroupRequest) SetImportTagValue(v string) *CreateApplicationGroupRequest {
	s.ImportTagValue = &v
	return s
}

func (s *CreateApplicationGroupRequest) SetName(v string) *CreateApplicationGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateApplicationGroupRequest) SetRegionId(v string) *CreateApplicationGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApplicationGroupRequest) Validate() error {
	return dara.Validate(s)
}
