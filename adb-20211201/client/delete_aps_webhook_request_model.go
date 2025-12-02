// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApsWebhookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteApsWebhookRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DeleteApsWebhookRequest
	GetRegionId() *string
	SetWebhookId(v string) *DeleteApsWebhookRequest
	GetWebhookId() *string
}

type DeleteApsWebhookRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-uf63i4ij56b***
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the webhook to be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// ***1*595*
	WebhookId *string `json:"WebhookId,omitempty" xml:"WebhookId,omitempty"`
}

func (s DeleteApsWebhookRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApsWebhookRequest) GoString() string {
	return s.String()
}

func (s *DeleteApsWebhookRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteApsWebhookRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteApsWebhookRequest) GetWebhookId() *string {
	return s.WebhookId
}

func (s *DeleteApsWebhookRequest) SetDBClusterId(v string) *DeleteApsWebhookRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteApsWebhookRequest) SetRegionId(v string) *DeleteApsWebhookRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteApsWebhookRequest) SetWebhookId(v string) *DeleteApsWebhookRequest {
	s.WebhookId = &v
	return s
}

func (s *DeleteApsWebhookRequest) Validate() error {
	return dara.Validate(s)
}
