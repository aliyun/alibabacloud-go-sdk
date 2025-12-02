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
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The type of the task. Valid value: Task type. SLS or OSS Export Task: ResultExport.
	//
	// This parameter is required.
	//
	// example:
	//
	// ResultExport
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The ID of the region in which to create the dedicated block storage cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The array of webhooks.
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
