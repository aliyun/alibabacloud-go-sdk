// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpListTokenResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*RuntimeToken) *PdpListTokenResult
	GetData() []*RuntimeToken
	SetRequestId(v string) *PdpListTokenResult
	GetRequestId() *string
	SetTotal(v int64) *PdpListTokenResult
	GetTotal() *int64
}

type PdpListTokenResult struct {
	Data      []*RuntimeToken `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	RequestId *string         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Total     *int64          `json:"total,omitempty" xml:"total,omitempty"`
}

func (s PdpListTokenResult) String() string {
	return dara.Prettify(s)
}

func (s PdpListTokenResult) GoString() string {
	return s.String()
}

func (s *PdpListTokenResult) GetData() []*RuntimeToken {
	return s.Data
}

func (s *PdpListTokenResult) GetRequestId() *string {
	return s.RequestId
}

func (s *PdpListTokenResult) GetTotal() *int64 {
	return s.Total
}

func (s *PdpListTokenResult) SetData(v []*RuntimeToken) *PdpListTokenResult {
	s.Data = v
	return s
}

func (s *PdpListTokenResult) SetRequestId(v string) *PdpListTokenResult {
	s.RequestId = &v
	return s
}

func (s *PdpListTokenResult) SetTotal(v int64) *PdpListTokenResult {
	s.Total = &v
	return s
}

func (s *PdpListTokenResult) Validate() error {
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
