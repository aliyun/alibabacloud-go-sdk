// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpConfigPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*PdpConfig) *PdpConfigPageResult
	GetData() []*PdpConfig
	SetRequestId(v string) *PdpConfigPageResult
	GetRequestId() *string
	SetTotal(v int64) *PdpConfigPageResult
	GetTotal() *int64
}

type PdpConfigPageResult struct {
	Data      []*PdpConfig `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	RequestId *string      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s PdpConfigPageResult) String() string {
	return dara.Prettify(s)
}

func (s PdpConfigPageResult) GoString() string {
	return s.String()
}

func (s *PdpConfigPageResult) GetData() []*PdpConfig {
	return s.Data
}

func (s *PdpConfigPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *PdpConfigPageResult) GetTotal() *int64 {
	return s.Total
}

func (s *PdpConfigPageResult) SetData(v []*PdpConfig) *PdpConfigPageResult {
	s.Data = v
	return s
}

func (s *PdpConfigPageResult) SetRequestId(v string) *PdpConfigPageResult {
	s.RequestId = &v
	return s
}

func (s *PdpConfigPageResult) SetTotal(v int64) *PdpConfigPageResult {
	s.Total = &v
	return s
}

func (s *PdpConfigPageResult) Validate() error {
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
