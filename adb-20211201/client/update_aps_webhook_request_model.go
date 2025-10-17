// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApsWebhookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *UpdateApsWebhookRequest
	GetDBClusterId() *string
	SetRegionId(v string) *UpdateApsWebhookRequest
	GetRegionId() *string
	SetWebhook(v []*UpdateApsWebhookRequestWebhook) *UpdateApsWebhookRequest
	GetWebhook() []*UpdateApsWebhookRequestWebhook
}

type UpdateApsWebhookRequest struct {
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
	RegionId *string                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Webhook  []*UpdateApsWebhookRequestWebhook `json:"Webhook,omitempty" xml:"Webhook,omitempty" type:"Repeated"`
}

func (s UpdateApsWebhookRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApsWebhookRequest) GoString() string {
	return s.String()
}

func (s *UpdateApsWebhookRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *UpdateApsWebhookRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateApsWebhookRequest) GetWebhook() []*UpdateApsWebhookRequestWebhook {
	return s.Webhook
}

func (s *UpdateApsWebhookRequest) SetDBClusterId(v string) *UpdateApsWebhookRequest {
	s.DBClusterId = &v
	return s
}

func (s *UpdateApsWebhookRequest) SetRegionId(v string) *UpdateApsWebhookRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateApsWebhookRequest) SetWebhook(v []*UpdateApsWebhookRequestWebhook) *UpdateApsWebhookRequest {
	s.Webhook = v
	return s
}

func (s *UpdateApsWebhookRequest) Validate() error {
	if s.Webhook != nil {
		for _, item := range s.Webhook {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateApsWebhookRequestWebhook struct {
	// example:
	//
	// ABC**
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// exampleWebhookName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// https://example.com/webhook
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// Webhook ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// **355****
	WebhookId *int64 `json:"WebhookId,omitempty" xml:"WebhookId,omitempty"`
	// example:
	//
	// dingtalk
	WebhookType *string `json:"WebhookType,omitempty" xml:"WebhookType,omitempty"`
}

func (s UpdateApsWebhookRequestWebhook) String() string {
	return dara.Prettify(s)
}

func (s UpdateApsWebhookRequestWebhook) GoString() string {
	return s.String()
}

func (s *UpdateApsWebhookRequestWebhook) GetKey() *string {
	return s.Key
}

func (s *UpdateApsWebhookRequestWebhook) GetName() *string {
	return s.Name
}

func (s *UpdateApsWebhookRequestWebhook) GetUrl() *string {
	return s.Url
}

func (s *UpdateApsWebhookRequestWebhook) GetWebhookId() *int64 {
	return s.WebhookId
}

func (s *UpdateApsWebhookRequestWebhook) GetWebhookType() *string {
	return s.WebhookType
}

func (s *UpdateApsWebhookRequestWebhook) SetKey(v string) *UpdateApsWebhookRequestWebhook {
	s.Key = &v
	return s
}

func (s *UpdateApsWebhookRequestWebhook) SetName(v string) *UpdateApsWebhookRequestWebhook {
	s.Name = &v
	return s
}

func (s *UpdateApsWebhookRequestWebhook) SetUrl(v string) *UpdateApsWebhookRequestWebhook {
	s.Url = &v
	return s
}

func (s *UpdateApsWebhookRequestWebhook) SetWebhookId(v int64) *UpdateApsWebhookRequestWebhook {
	s.WebhookId = &v
	return s
}

func (s *UpdateApsWebhookRequestWebhook) SetWebhookType(v string) *UpdateApsWebhookRequestWebhook {
	s.WebhookType = &v
	return s
}

func (s *UpdateApsWebhookRequestWebhook) Validate() error {
	return dara.Validate(s)
}
