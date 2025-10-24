// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWaterMarkTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *UpdateWaterMarkTemplateRequest
	GetConfig() *string
	SetName(v string) *UpdateWaterMarkTemplateRequest
	GetName() *string
	SetOwnerAccount(v string) *UpdateWaterMarkTemplateRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateWaterMarkTemplateRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateWaterMarkTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateWaterMarkTemplateRequest
	GetResourceOwnerId() *int64
	SetWaterMarkTemplateId(v string) *UpdateWaterMarkTemplateRequest
	GetWaterMarkTemplateId() *string
}

type UpdateWaterMarkTemplateRequest struct {
	// The updated configuration of the watermark template. The value is a JSON object. For more information, see [Parameter details](https://help.aliyun.com/document_detail/29253.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Width":"10","Height":"30","Dx":"10","Dy":"5","Type":"Image","Timeline":{"Start":"0","Duration":"10"}}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The new name of the watermark template. The value can contain letters and digits and can be up to 128 bytes in size.
	//
	// This parameter is required.
	//
	// example:
	//
	// example-watermark-****
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the watermark template that you want to update. To obtain the ID of the watermark template, you can log on to the **ApsaraVideo Media Processing (MPS) console*	- and choose **Global Settings*	- > **Watermark Templates*	- in the left-side navigation pane.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3780bd69b2b74540bc7b1096f564****
	WaterMarkTemplateId *string `json:"WaterMarkTemplateId,omitempty" xml:"WaterMarkTemplateId,omitempty"`
}

func (s UpdateWaterMarkTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWaterMarkTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateWaterMarkTemplateRequest) GetConfig() *string {
	return s.Config
}

func (s *UpdateWaterMarkTemplateRequest) GetName() *string {
	return s.Name
}

func (s *UpdateWaterMarkTemplateRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateWaterMarkTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateWaterMarkTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateWaterMarkTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateWaterMarkTemplateRequest) GetWaterMarkTemplateId() *string {
	return s.WaterMarkTemplateId
}

func (s *UpdateWaterMarkTemplateRequest) SetConfig(v string) *UpdateWaterMarkTemplateRequest {
	s.Config = &v
	return s
}

func (s *UpdateWaterMarkTemplateRequest) SetName(v string) *UpdateWaterMarkTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateWaterMarkTemplateRequest) SetOwnerAccount(v string) *UpdateWaterMarkTemplateRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateWaterMarkTemplateRequest) SetOwnerId(v int64) *UpdateWaterMarkTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateWaterMarkTemplateRequest) SetResourceOwnerAccount(v string) *UpdateWaterMarkTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateWaterMarkTemplateRequest) SetResourceOwnerId(v int64) *UpdateWaterMarkTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateWaterMarkTemplateRequest) SetWaterMarkTemplateId(v string) *UpdateWaterMarkTemplateRequest {
	s.WaterMarkTemplateId = &v
	return s
}

func (s *UpdateWaterMarkTemplateRequest) Validate() error {
	return dara.Validate(s)
}
