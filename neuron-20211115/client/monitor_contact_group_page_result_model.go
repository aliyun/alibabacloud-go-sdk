// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorContactGroupPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*MonitorContactGroup) *MonitorContactGroupPageResult
	GetData() []*MonitorContactGroup
	SetIntTotal(v int32) *MonitorContactGroupPageResult
	GetIntTotal() *int32
	SetRequestId(v string) *MonitorContactGroupPageResult
	GetRequestId() *string
	SetTotal(v int64) *MonitorContactGroupPageResult
	GetTotal() *int64
}

type MonitorContactGroupPageResult struct {
	Data      []*MonitorContactGroup `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	IntTotal  *int32                 `json:"intTotal,omitempty" xml:"intTotal,omitempty"`
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s MonitorContactGroupPageResult) String() string {
	return dara.Prettify(s)
}

func (s MonitorContactGroupPageResult) GoString() string {
	return s.String()
}

func (s *MonitorContactGroupPageResult) GetData() []*MonitorContactGroup {
	return s.Data
}

func (s *MonitorContactGroupPageResult) GetIntTotal() *int32 {
	return s.IntTotal
}

func (s *MonitorContactGroupPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *MonitorContactGroupPageResult) GetTotal() *int64 {
	return s.Total
}

func (s *MonitorContactGroupPageResult) SetData(v []*MonitorContactGroup) *MonitorContactGroupPageResult {
	s.Data = v
	return s
}

func (s *MonitorContactGroupPageResult) SetIntTotal(v int32) *MonitorContactGroupPageResult {
	s.IntTotal = &v
	return s
}

func (s *MonitorContactGroupPageResult) SetRequestId(v string) *MonitorContactGroupPageResult {
	s.RequestId = &v
	return s
}

func (s *MonitorContactGroupPageResult) SetTotal(v int64) *MonitorContactGroupPageResult {
	s.Total = &v
	return s
}

func (s *MonitorContactGroupPageResult) Validate() error {
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
