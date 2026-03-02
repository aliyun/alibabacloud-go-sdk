// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpServicePageResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*PdpService) *PdpServicePageResult
	GetData() []*PdpService
	SetRequestId(v string) *PdpServicePageResult
	GetRequestId() *string
	SetTotal(v int64) *PdpServicePageResult
	GetTotal() *int64
}

type PdpServicePageResult struct {
	Data      []*PdpService `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	RequestId *string       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s PdpServicePageResult) String() string {
	return dara.Prettify(s)
}

func (s PdpServicePageResult) GoString() string {
	return s.String()
}

func (s *PdpServicePageResult) GetData() []*PdpService {
	return s.Data
}

func (s *PdpServicePageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *PdpServicePageResult) GetTotal() *int64 {
	return s.Total
}

func (s *PdpServicePageResult) SetData(v []*PdpService) *PdpServicePageResult {
	s.Data = v
	return s
}

func (s *PdpServicePageResult) SetRequestId(v string) *PdpServicePageResult {
	s.RequestId = &v
	return s
}

func (s *PdpServicePageResult) SetTotal(v int64) *PdpServicePageResult {
	s.Total = &v
	return s
}

func (s *PdpServicePageResult) Validate() error {
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
