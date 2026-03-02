// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpHistoryConfigPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*PdpHistoryConfig) *PdpHistoryConfigPageResult
	GetData() []*PdpHistoryConfig
	SetRequestId(v string) *PdpHistoryConfigPageResult
	GetRequestId() *string
	SetTotal(v int64) *PdpHistoryConfigPageResult
	GetTotal() *int64
}

type PdpHistoryConfigPageResult struct {
	Data      []*PdpHistoryConfig `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	RequestId *string             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s PdpHistoryConfigPageResult) String() string {
	return dara.Prettify(s)
}

func (s PdpHistoryConfigPageResult) GoString() string {
	return s.String()
}

func (s *PdpHistoryConfigPageResult) GetData() []*PdpHistoryConfig {
	return s.Data
}

func (s *PdpHistoryConfigPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *PdpHistoryConfigPageResult) GetTotal() *int64 {
	return s.Total
}

func (s *PdpHistoryConfigPageResult) SetData(v []*PdpHistoryConfig) *PdpHistoryConfigPageResult {
	s.Data = v
	return s
}

func (s *PdpHistoryConfigPageResult) SetRequestId(v string) *PdpHistoryConfigPageResult {
	s.RequestId = &v
	return s
}

func (s *PdpHistoryConfigPageResult) SetTotal(v int64) *PdpHistoryConfigPageResult {
	s.Total = &v
	return s
}

func (s *PdpHistoryConfigPageResult) Validate() error {
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
