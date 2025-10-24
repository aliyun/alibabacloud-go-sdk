// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWaterMarkTemplateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *QueryWaterMarkTemplateListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *QueryWaterMarkTemplateListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryWaterMarkTemplateListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryWaterMarkTemplateListRequest
	GetResourceOwnerId() *int64
	SetWaterMarkTemplateIds(v string) *QueryWaterMarkTemplateListRequest
	GetWaterMarkTemplateIds() *string
}

type QueryWaterMarkTemplateListRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IDs of the watermark templates that you want to query. To obtain the IDs of the watermark templates, you can log on to the **ApsaraVideo Media Processing (MPS) console*	- and choose **Global Settings*	- > **Watermark Templates*	- in the left-side navigation pane. You can query up to 10 watermark templates at a time. Separate multiple IDs of watermark templates with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 3780bd69b2b74540bc7b1096f564****
	WaterMarkTemplateIds *string `json:"WaterMarkTemplateIds,omitempty" xml:"WaterMarkTemplateIds,omitempty"`
}

func (s QueryWaterMarkTemplateListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryWaterMarkTemplateListRequest) GoString() string {
	return s.String()
}

func (s *QueryWaterMarkTemplateListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *QueryWaterMarkTemplateListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryWaterMarkTemplateListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryWaterMarkTemplateListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryWaterMarkTemplateListRequest) GetWaterMarkTemplateIds() *string {
	return s.WaterMarkTemplateIds
}

func (s *QueryWaterMarkTemplateListRequest) SetOwnerAccount(v string) *QueryWaterMarkTemplateListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *QueryWaterMarkTemplateListRequest) SetOwnerId(v int64) *QueryWaterMarkTemplateListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryWaterMarkTemplateListRequest) SetResourceOwnerAccount(v string) *QueryWaterMarkTemplateListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryWaterMarkTemplateListRequest) SetResourceOwnerId(v int64) *QueryWaterMarkTemplateListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryWaterMarkTemplateListRequest) SetWaterMarkTemplateIds(v string) *QueryWaterMarkTemplateListRequest {
	s.WaterMarkTemplateIds = &v
	return s
}

func (s *QueryWaterMarkTemplateListRequest) Validate() error {
	return dara.Validate(s)
}
