// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeleteTemplateRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteTemplateRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteTemplateRequest
	GetResourceOwnerId() *int64
	SetTemplateId(v string) *DeleteTemplateRequest
	GetTemplateId() *string
}

type DeleteTemplateRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the custom transcoding template that you want to delete. To obtain the ID of the custom transcoding template, you can log on to the **ApsaraVideo Media Processing (MPS) console*	- and choose **Global Settings*	- > **Encoding Templates*	- in the left-side navigation pane.
	//
	// This parameter is required.
	//
	// example:
	//
	// 16f01ad6175e4230ac42bb5182cd****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplateRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteTemplateRequest) SetOwnerAccount(v string) *DeleteTemplateRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTemplateRequest) SetOwnerId(v int64) *DeleteTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTemplateRequest) SetResourceOwnerAccount(v string) *DeleteTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTemplateRequest) SetResourceOwnerId(v int64) *DeleteTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTemplateRequest) SetTemplateId(v string) *DeleteTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteTemplateRequest) Validate() error {
	return dara.Validate(s)
}
