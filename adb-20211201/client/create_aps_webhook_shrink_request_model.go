// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApsWebhookShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateApsWebhookShrinkRequest
	GetDBClusterId() *string
	SetJobType(v string) *CreateApsWebhookShrinkRequest
	GetJobType() *string
	SetRegionId(v string) *CreateApsWebhookShrinkRequest
	GetRegionId() *string
	SetWebhookShrink(v string) *CreateApsWebhookShrinkRequest
	GetWebhookShrink() *string
}

type CreateApsWebhookShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ResultExport
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	WebhookShrink *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s CreateApsWebhookShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApsWebhookShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateApsWebhookShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateApsWebhookShrinkRequest) GetJobType() *string {
	return s.JobType
}

func (s *CreateApsWebhookShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApsWebhookShrinkRequest) GetWebhookShrink() *string {
	return s.WebhookShrink
}

func (s *CreateApsWebhookShrinkRequest) SetDBClusterId(v string) *CreateApsWebhookShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateApsWebhookShrinkRequest) SetJobType(v string) *CreateApsWebhookShrinkRequest {
	s.JobType = &v
	return s
}

func (s *CreateApsWebhookShrinkRequest) SetRegionId(v string) *CreateApsWebhookShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApsWebhookShrinkRequest) SetWebhookShrink(v string) *CreateApsWebhookShrinkRequest {
	s.WebhookShrink = &v
	return s
}

func (s *CreateApsWebhookShrinkRequest) Validate() error {
	return dara.Validate(s)
}
