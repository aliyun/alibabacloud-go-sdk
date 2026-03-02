// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorWebhookUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *MonitorWebhookUpdateCmd
	GetId() *int64
	SetMethod(v string) *MonitorWebhookUpdateCmd
	GetMethod() *string
	SetName(v string) *MonitorWebhookUpdateCmd
	GetName() *string
	SetType(v string) *MonitorWebhookUpdateCmd
	GetType() *string
	SetUrl(v string) *MonitorWebhookUpdateCmd
	GetUrl() *string
}

type MonitorWebhookUpdateCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// POST
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 钉钉机器人
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DINGDING
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=**********
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s MonitorWebhookUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s MonitorWebhookUpdateCmd) GoString() string {
	return s.String()
}

func (s *MonitorWebhookUpdateCmd) GetId() *int64 {
	return s.Id
}

func (s *MonitorWebhookUpdateCmd) GetMethod() *string {
	return s.Method
}

func (s *MonitorWebhookUpdateCmd) GetName() *string {
	return s.Name
}

func (s *MonitorWebhookUpdateCmd) GetType() *string {
	return s.Type
}

func (s *MonitorWebhookUpdateCmd) GetUrl() *string {
	return s.Url
}

func (s *MonitorWebhookUpdateCmd) SetId(v int64) *MonitorWebhookUpdateCmd {
	s.Id = &v
	return s
}

func (s *MonitorWebhookUpdateCmd) SetMethod(v string) *MonitorWebhookUpdateCmd {
	s.Method = &v
	return s
}

func (s *MonitorWebhookUpdateCmd) SetName(v string) *MonitorWebhookUpdateCmd {
	s.Name = &v
	return s
}

func (s *MonitorWebhookUpdateCmd) SetType(v string) *MonitorWebhookUpdateCmd {
	s.Type = &v
	return s
}

func (s *MonitorWebhookUpdateCmd) SetUrl(v string) *MonitorWebhookUpdateCmd {
	s.Url = &v
	return s
}

func (s *MonitorWebhookUpdateCmd) Validate() error {
	return dara.Validate(s)
}
