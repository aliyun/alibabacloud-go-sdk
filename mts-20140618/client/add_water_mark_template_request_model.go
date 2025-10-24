// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWaterMarkTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *AddWaterMarkTemplateRequest
	GetConfig() *string
	SetName(v string) *AddWaterMarkTemplateRequest
	GetName() *string
	SetOwnerAccount(v string) *AddWaterMarkTemplateRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddWaterMarkTemplateRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddWaterMarkTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddWaterMarkTemplateRequest
	GetResourceOwnerId() *int64
}

type AddWaterMarkTemplateRequest struct {
	// The configuration of the watermark template. The value is a JSON object. For more information, see the "WaterMarks" section of the [Parameter details](https://help.aliyun.com/document_detail/29253.html) topic.
	//
	// > If you do not require a positive correlation between the size of text in the watermark and the resolution, you can enable adaptation for the watermark. To do so, add `[\\"adaptive\\"]=true` to the TextWaterMark parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Width":"10","Height":"30","Dx":"10","Dy":"5","ReferPos":"TopRight","Type":"Image","Timeline":{"Start":"0","Duration":"10"}}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The name of the watermark template. The value can contain letters and digits and can be up to 128 bytes in size.
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
}

func (s AddWaterMarkTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s AddWaterMarkTemplateRequest) GoString() string {
	return s.String()
}

func (s *AddWaterMarkTemplateRequest) GetConfig() *string {
	return s.Config
}

func (s *AddWaterMarkTemplateRequest) GetName() *string {
	return s.Name
}

func (s *AddWaterMarkTemplateRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddWaterMarkTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddWaterMarkTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddWaterMarkTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddWaterMarkTemplateRequest) SetConfig(v string) *AddWaterMarkTemplateRequest {
	s.Config = &v
	return s
}

func (s *AddWaterMarkTemplateRequest) SetName(v string) *AddWaterMarkTemplateRequest {
	s.Name = &v
	return s
}

func (s *AddWaterMarkTemplateRequest) SetOwnerAccount(v string) *AddWaterMarkTemplateRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddWaterMarkTemplateRequest) SetOwnerId(v int64) *AddWaterMarkTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *AddWaterMarkTemplateRequest) SetResourceOwnerAccount(v string) *AddWaterMarkTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddWaterMarkTemplateRequest) SetResourceOwnerId(v int64) *AddWaterMarkTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddWaterMarkTemplateRequest) Validate() error {
	return dara.Validate(s)
}
