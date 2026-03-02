// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorContactPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*MonitorContact) *MonitorContactPageResult
	GetData() []*MonitorContact
	SetIntTotal(v int32) *MonitorContactPageResult
	GetIntTotal() *int32
	SetRequestId(v string) *MonitorContactPageResult
	GetRequestId() *string
	SetTotal(v int64) *MonitorContactPageResult
	GetTotal() *int64
}

type MonitorContactPageResult struct {
	Data      []*MonitorContact `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	IntTotal  *int32            `json:"intTotal,omitempty" xml:"intTotal,omitempty"`
	RequestId *string           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s MonitorContactPageResult) String() string {
	return dara.Prettify(s)
}

func (s MonitorContactPageResult) GoString() string {
	return s.String()
}

func (s *MonitorContactPageResult) GetData() []*MonitorContact {
	return s.Data
}

func (s *MonitorContactPageResult) GetIntTotal() *int32 {
	return s.IntTotal
}

func (s *MonitorContactPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *MonitorContactPageResult) GetTotal() *int64 {
	return s.Total
}

func (s *MonitorContactPageResult) SetData(v []*MonitorContact) *MonitorContactPageResult {
	s.Data = v
	return s
}

func (s *MonitorContactPageResult) SetIntTotal(v int32) *MonitorContactPageResult {
	s.IntTotal = &v
	return s
}

func (s *MonitorContactPageResult) SetRequestId(v string) *MonitorContactPageResult {
	s.RequestId = &v
	return s
}

func (s *MonitorContactPageResult) SetTotal(v int64) *MonitorContactPageResult {
	s.Total = &v
	return s
}

func (s *MonitorContactPageResult) Validate() error {
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
