// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApsWebhookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateApsWebhookRequest
	GetDBClusterId() *string
	SetJobType(v string) *CreateApsWebhookRequest
	GetJobType() *string
	SetRegionId(v string) *CreateApsWebhookRequest
	GetRegionId() *string
	SetWebhook(v []*CreateApsWebhookRequestWebhook) *CreateApsWebhookRequest
	GetWebhook() []*CreateApsWebhookRequestWebhook
}

type CreateApsWebhookRequest struct {
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
	Webhook []*CreateApsWebhookRequestWebhook `json:"Webhook,omitempty" xml:"Webhook,omitempty" type:"Repeated"`
}

func (s CreateApsWebhookRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApsWebhookRequest) GoString() string {
	return s.String()
}

func (s *CreateApsWebhookRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateApsWebhookRequest) GetJobType() *string {
	return s.JobType
}

func (s *CreateApsWebhookRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApsWebhookRequest) GetWebhook() []*CreateApsWebhookRequestWebhook {
	return s.Webhook
}

func (s *CreateApsWebhookRequest) SetDBClusterId(v string) *CreateApsWebhookRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateApsWebhookRequest) SetJobType(v string) *CreateApsWebhookRequest {
	s.JobType = &v
	return s
}

func (s *CreateApsWebhookRequest) SetRegionId(v string) *CreateApsWebhookRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApsWebhookRequest) SetWebhook(v []*CreateApsWebhookRequestWebhook) *CreateApsWebhookRequest {
	s.Webhook = v
	return s
}

func (s *CreateApsWebhookRequest) Validate() error {
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

type CreateApsWebhookRequestWebhook struct {
	// Signed key.
	//
	// example:
	//
	// ***
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name of the webhook.
	//
	// example:
	//
	// MyWebhookName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request path.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/webhook
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The notification method. Valid values: dingtalk. lark.
	//
	// This parameter is required.
	//
	// example:
	//
	// dingtalk
	WebhookType *string `json:"WebhookType,omitempty" xml:"WebhookType,omitempty"`
}

func (s CreateApsWebhookRequestWebhook) String() string {
	return dara.Prettify(s)
}

func (s CreateApsWebhookRequestWebhook) GoString() string {
	return s.String()
}

func (s *CreateApsWebhookRequestWebhook) GetKey() *string {
	return s.Key
}

func (s *CreateApsWebhookRequestWebhook) GetName() *string {
	return s.Name
}

func (s *CreateApsWebhookRequestWebhook) GetUrl() *string {
	return s.Url
}

func (s *CreateApsWebhookRequestWebhook) GetWebhookType() *string {
	return s.WebhookType
}

func (s *CreateApsWebhookRequestWebhook) SetKey(v string) *CreateApsWebhookRequestWebhook {
	s.Key = &v
	return s
}

func (s *CreateApsWebhookRequestWebhook) SetName(v string) *CreateApsWebhookRequestWebhook {
	s.Name = &v
	return s
}

func (s *CreateApsWebhookRequestWebhook) SetUrl(v string) *CreateApsWebhookRequestWebhook {
	s.Url = &v
	return s
}

func (s *CreateApsWebhookRequestWebhook) SetWebhookType(v string) *CreateApsWebhookRequestWebhook {
	s.WebhookType = &v
	return s
}

func (s *CreateApsWebhookRequestWebhook) Validate() error {
	return dara.Validate(s)
}
