// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorWebhookPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*MonitorWebhook) *MonitorWebhookPageResult
	GetData() []*MonitorWebhook
	SetIntTotal(v int32) *MonitorWebhookPageResult
	GetIntTotal() *int32
	SetRequestId(v string) *MonitorWebhookPageResult
	GetRequestId() *string
	SetTotal(v int64) *MonitorWebhookPageResult
	GetTotal() *int64
}

type MonitorWebhookPageResult struct {
	Data      []*MonitorWebhook `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	IntTotal  *int32            `json:"intTotal,omitempty" xml:"intTotal,omitempty"`
	RequestId *string           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s MonitorWebhookPageResult) String() string {
	return dara.Prettify(s)
}

func (s MonitorWebhookPageResult) GoString() string {
	return s.String()
}

func (s *MonitorWebhookPageResult) GetData() []*MonitorWebhook {
	return s.Data
}

func (s *MonitorWebhookPageResult) GetIntTotal() *int32 {
	return s.IntTotal
}

func (s *MonitorWebhookPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *MonitorWebhookPageResult) GetTotal() *int64 {
	return s.Total
}

func (s *MonitorWebhookPageResult) SetData(v []*MonitorWebhook) *MonitorWebhookPageResult {
	s.Data = v
	return s
}

func (s *MonitorWebhookPageResult) SetIntTotal(v int32) *MonitorWebhookPageResult {
	s.IntTotal = &v
	return s
}

func (s *MonitorWebhookPageResult) SetRequestId(v string) *MonitorWebhookPageResult {
	s.RequestId = &v
	return s
}

func (s *MonitorWebhookPageResult) SetTotal(v int64) *MonitorWebhookPageResult {
	s.Total = &v
	return s
}

func (s *MonitorWebhookPageResult) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
