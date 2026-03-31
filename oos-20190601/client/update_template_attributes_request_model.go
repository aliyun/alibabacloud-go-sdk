// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v string) *UpdateTemplateAttributesRequest
	GetAccountIds() *string
	SetIsFavorite(v bool) *UpdateTemplateAttributesRequest
	GetIsFavorite() *bool
	SetRegionId(v string) *UpdateTemplateAttributesRequest
	GetRegionId() *string
	SetSharePermissionAction(v string) *UpdateTemplateAttributesRequest
	GetSharePermissionAction() *string
	SetShareTemplateVersion(v string) *UpdateTemplateAttributesRequest
	GetShareTemplateVersion() *string
	SetTemplateName(v string) *UpdateTemplateAttributesRequest
	GetTemplateName() *string
}

type UpdateTemplateAttributesRequest struct {
	// example:
	//
	// ["1319994761488600"]
	AccountIds *string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty"`
	// example:
	//
	// true
	IsFavorite *bool `json:"IsFavorite,omitempty" xml:"IsFavorite,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Share
	SharePermissionAction *string `json:"SharePermissionAction,omitempty" xml:"SharePermissionAction,omitempty"`
	// example:
	//
	// v1
	ShareTemplateVersion *string `json:"ShareTemplateVersion,omitempty" xml:"ShareTemplateVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s UpdateTemplateAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateAttributesRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateAttributesRequest) GetAccountIds() *string {
	return s.AccountIds
}

func (s *UpdateTemplateAttributesRequest) GetIsFavorite() *bool {
	return s.IsFavorite
}

func (s *UpdateTemplateAttributesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateTemplateAttributesRequest) GetSharePermissionAction() *string {
	return s.SharePermissionAction
}

func (s *UpdateTemplateAttributesRequest) GetShareTemplateVersion() *string {
	return s.ShareTemplateVersion
}

func (s *UpdateTemplateAttributesRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *UpdateTemplateAttributesRequest) SetAccountIds(v string) *UpdateTemplateAttributesRequest {
	s.AccountIds = &v
	return s
}

func (s *UpdateTemplateAttributesRequest) SetIsFavorite(v bool) *UpdateTemplateAttributesRequest {
	s.IsFavorite = &v
	return s
}

func (s *UpdateTemplateAttributesRequest) SetRegionId(v string) *UpdateTemplateAttributesRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateTemplateAttributesRequest) SetSharePermissionAction(v string) *UpdateTemplateAttributesRequest {
	s.SharePermissionAction = &v
	return s
}

func (s *UpdateTemplateAttributesRequest) SetShareTemplateVersion(v string) *UpdateTemplateAttributesRequest {
	s.ShareTemplateVersion = &v
	return s
}

func (s *UpdateTemplateAttributesRequest) SetTemplateName(v string) *UpdateTemplateAttributesRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateTemplateAttributesRequest) Validate() error {
	return dara.Validate(s)
}
