// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *CopyTemplateRequest
	GetCustSpaceId() *string
	SetLanguage(v string) *CopyTemplateRequest
	GetLanguage() *string
	SetOwnerId(v int64) *CopyTemplateRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CopyTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CopyTemplateRequest
	GetResourceOwnerId() *int64
	SetSceneTemplateCode(v string) *CopyTemplateRequest
	GetSceneTemplateCode() *string
	SetSceneTemplateName(v string) *CopyTemplateRequest
	GetSceneTemplateName() *string
}

type CopyTemplateRequest struct {
	// example:
	//
	// 示例值示例值示例值
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Language             *string `json:"Language,omitempty" xml:"Language,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12456
	SceneTemplateCode *string `json:"SceneTemplateCode,omitempty" xml:"SceneTemplateCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	SceneTemplateName *string `json:"SceneTemplateName,omitempty" xml:"SceneTemplateName,omitempty"`
}

func (s CopyTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyTemplateRequest) GoString() string {
	return s.String()
}

func (s *CopyTemplateRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *CopyTemplateRequest) GetLanguage() *string {
	return s.Language
}

func (s *CopyTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CopyTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CopyTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CopyTemplateRequest) GetSceneTemplateCode() *string {
	return s.SceneTemplateCode
}

func (s *CopyTemplateRequest) GetSceneTemplateName() *string {
	return s.SceneTemplateName
}

func (s *CopyTemplateRequest) SetCustSpaceId(v string) *CopyTemplateRequest {
	s.CustSpaceId = &v
	return s
}

func (s *CopyTemplateRequest) SetLanguage(v string) *CopyTemplateRequest {
	s.Language = &v
	return s
}

func (s *CopyTemplateRequest) SetOwnerId(v int64) *CopyTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *CopyTemplateRequest) SetResourceOwnerAccount(v string) *CopyTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CopyTemplateRequest) SetResourceOwnerId(v int64) *CopyTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CopyTemplateRequest) SetSceneTemplateCode(v string) *CopyTemplateRequest {
	s.SceneTemplateCode = &v
	return s
}

func (s *CopyTemplateRequest) SetSceneTemplateName(v string) *CopyTemplateRequest {
	s.SceneTemplateName = &v
	return s
}

func (s *CopyTemplateRequest) Validate() error {
	return dara.Validate(s)
}
