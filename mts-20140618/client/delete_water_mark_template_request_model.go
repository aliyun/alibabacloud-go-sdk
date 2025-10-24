// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWaterMarkTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeleteWaterMarkTemplateRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteWaterMarkTemplateRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteWaterMarkTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteWaterMarkTemplateRequest
	GetResourceOwnerId() *int64
	SetWaterMarkTemplateId(v string) *DeleteWaterMarkTemplateRequest
	GetWaterMarkTemplateId() *string
}

type DeleteWaterMarkTemplateRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the watermark template that you want to delete. To obtain the template ID, you can log on to the **ApsaraVideo Media Processing (MPS) console*	- and choose **Global Settings*	- > **Watermark Templates*	- in the left-side navigation pane.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3780bd69b2b74540bc7b1096f564****
	WaterMarkTemplateId *string `json:"WaterMarkTemplateId,omitempty" xml:"WaterMarkTemplateId,omitempty"`
}

func (s DeleteWaterMarkTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWaterMarkTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteWaterMarkTemplateRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteWaterMarkTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteWaterMarkTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteWaterMarkTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteWaterMarkTemplateRequest) GetWaterMarkTemplateId() *string {
	return s.WaterMarkTemplateId
}

func (s *DeleteWaterMarkTemplateRequest) SetOwnerAccount(v string) *DeleteWaterMarkTemplateRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteWaterMarkTemplateRequest) SetOwnerId(v int64) *DeleteWaterMarkTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteWaterMarkTemplateRequest) SetResourceOwnerAccount(v string) *DeleteWaterMarkTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteWaterMarkTemplateRequest) SetResourceOwnerId(v int64) *DeleteWaterMarkTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteWaterMarkTemplateRequest) SetWaterMarkTemplateId(v string) *DeleteWaterMarkTemplateRequest {
	s.WaterMarkTemplateId = &v
	return s
}

func (s *DeleteWaterMarkTemplateRequest) Validate() error {
	return dara.Validate(s)
}
