// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorWebhook interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *MonitorWebhook
	GetEnterpriseId() *int64
	SetGmtCreate(v string) *MonitorWebhook
	GetGmtCreate() *string
	SetGmtModified(v string) *MonitorWebhook
	GetGmtModified() *string
	SetId(v int64) *MonitorWebhook
	GetId() *int64
	SetMethod(v string) *MonitorWebhook
	GetMethod() *string
	SetName(v string) *MonitorWebhook
	GetName() *string
	SetType(v string) *MonitorWebhook
	GetType() *string
	SetUrl(v string) *MonitorWebhook
	GetUrl() *string
}

type MonitorWebhook struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	EnterpriseId *int64 `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	// example:
	//
	// 2022-05-01T00:00:00.000Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2022-05-01T00:00:00.000Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
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

func (s MonitorWebhook) String() string {
	return dara.Prettify(s)
}

func (s MonitorWebhook) GoString() string {
	return s.String()
}

func (s *MonitorWebhook) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *MonitorWebhook) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *MonitorWebhook) GetGmtModified() *string {
	return s.GmtModified
}

func (s *MonitorWebhook) GetId() *int64 {
	return s.Id
}

func (s *MonitorWebhook) GetMethod() *string {
	return s.Method
}

func (s *MonitorWebhook) GetName() *string {
	return s.Name
}

func (s *MonitorWebhook) GetType() *string {
	return s.Type
}

func (s *MonitorWebhook) GetUrl() *string {
	return s.Url
}

func (s *MonitorWebhook) SetEnterpriseId(v int64) *MonitorWebhook {
	s.EnterpriseId = &v
	return s
}

func (s *MonitorWebhook) SetGmtCreate(v string) *MonitorWebhook {
	s.GmtCreate = &v
	return s
}

func (s *MonitorWebhook) SetGmtModified(v string) *MonitorWebhook {
	s.GmtModified = &v
	return s
}

func (s *MonitorWebhook) SetId(v int64) *MonitorWebhook {
	s.Id = &v
	return s
}

func (s *MonitorWebhook) SetMethod(v string) *MonitorWebhook {
	s.Method = &v
	return s
}

func (s *MonitorWebhook) SetName(v string) *MonitorWebhook {
	s.Name = &v
	return s
}

func (s *MonitorWebhook) SetType(v string) *MonitorWebhook {
	s.Type = &v
	return s
}

func (s *MonitorWebhook) SetUrl(v string) *MonitorWebhook {
	s.Url = &v
	return s
}

func (s *MonitorWebhook) Validate() error {
	return dara.Validate(s)
}
