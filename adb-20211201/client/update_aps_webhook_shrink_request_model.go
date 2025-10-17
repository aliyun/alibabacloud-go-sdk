// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApsWebhookShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *UpdateApsWebhookShrinkRequest
	GetDBClusterId() *string
	SetRegionId(v string) *UpdateApsWebhookShrinkRequest
	GetRegionId() *string
	SetWebhookShrink(v string) *UpdateApsWebhookShrinkRequest
	GetWebhookShrink() *string
}

type UpdateApsWebhookShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// exampleDBClusterId
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// exampleRegionId
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	WebhookShrink *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s UpdateApsWebhookShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApsWebhookShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateApsWebhookShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *UpdateApsWebhookShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateApsWebhookShrinkRequest) GetWebhookShrink() *string {
	return s.WebhookShrink
}

func (s *UpdateApsWebhookShrinkRequest) SetDBClusterId(v string) *UpdateApsWebhookShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *UpdateApsWebhookShrinkRequest) SetRegionId(v string) *UpdateApsWebhookShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateApsWebhookShrinkRequest) SetWebhookShrink(v string) *UpdateApsWebhookShrinkRequest {
	s.WebhookShrink = &v
	return s
}

func (s *UpdateApsWebhookShrinkRequest) Validate() error {
	return dara.Validate(s)
}
