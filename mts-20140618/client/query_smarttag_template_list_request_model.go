// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmarttagTemplateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *QuerySmarttagTemplateListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *QuerySmarttagTemplateListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QuerySmarttagTemplateListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySmarttagTemplateListRequest
	GetResourceOwnerId() *int64
	SetTemplateId(v string) *QuerySmarttagTemplateListRequest
	GetTemplateId() *string
}

type QuerySmarttagTemplateListRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the template. You can obtain the template ID from the response of the [AddSmarttagTemplate](https://help.aliyun.com/document_detail/187759.html) operation. If you set this parameter to a specific value, the information about the corresponding template is returned. If you do not specify this parameter, the operation returns the information about all the templates that are created by the current RAM user.
	//
	// example:
	//
	// 05de22f255284c7a8d2aab535dde****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s QuerySmarttagTemplateListRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySmarttagTemplateListRequest) GoString() string {
	return s.String()
}

func (s *QuerySmarttagTemplateListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *QuerySmarttagTemplateListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySmarttagTemplateListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySmarttagTemplateListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySmarttagTemplateListRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *QuerySmarttagTemplateListRequest) SetOwnerAccount(v string) *QuerySmarttagTemplateListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *QuerySmarttagTemplateListRequest) SetOwnerId(v int64) *QuerySmarttagTemplateListRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySmarttagTemplateListRequest) SetResourceOwnerAccount(v string) *QuerySmarttagTemplateListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySmarttagTemplateListRequest) SetResourceOwnerId(v int64) *QuerySmarttagTemplateListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySmarttagTemplateListRequest) SetTemplateId(v string) *QuerySmarttagTemplateListRequest {
	s.TemplateId = &v
	return s
}

func (s *QuerySmarttagTemplateListRequest) Validate() error {
	return dara.Validate(s)
}
