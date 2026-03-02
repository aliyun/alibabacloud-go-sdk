// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMonitorNotifyObjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *ListMonitorNotifyObjectsRequest
	GetEnterpriseId() *int64
	SetName(v string) *ListMonitorNotifyObjectsRequest
	GetName() *string
	SetType(v int32) *ListMonitorNotifyObjectsRequest
	GetType() *int32
	SetWebhookType(v string) *ListMonitorNotifyObjectsRequest
	GetWebhookType() *string
}

type ListMonitorNotifyObjectsRequest struct {
	// example:
	//
	// 2
	EnterpriseId *int64  `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// DINGDING
	WebhookType *string `json:"webhookType,omitempty" xml:"webhookType,omitempty"`
}

func (s ListMonitorNotifyObjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMonitorNotifyObjectsRequest) GoString() string {
	return s.String()
}

func (s *ListMonitorNotifyObjectsRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ListMonitorNotifyObjectsRequest) GetName() *string {
	return s.Name
}

func (s *ListMonitorNotifyObjectsRequest) GetType() *int32 {
	return s.Type
}

func (s *ListMonitorNotifyObjectsRequest) GetWebhookType() *string {
	return s.WebhookType
}

func (s *ListMonitorNotifyObjectsRequest) SetEnterpriseId(v int64) *ListMonitorNotifyObjectsRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ListMonitorNotifyObjectsRequest) SetName(v string) *ListMonitorNotifyObjectsRequest {
	s.Name = &v
	return s
}

func (s *ListMonitorNotifyObjectsRequest) SetType(v int32) *ListMonitorNotifyObjectsRequest {
	s.Type = &v
	return s
}

func (s *ListMonitorNotifyObjectsRequest) SetWebhookType(v string) *ListMonitorNotifyObjectsRequest {
	s.WebhookType = &v
	return s
}

func (s *ListMonitorNotifyObjectsRequest) Validate() error {
	return dara.Validate(s)
}
