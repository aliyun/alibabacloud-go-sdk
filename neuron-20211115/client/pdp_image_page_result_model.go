// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpImagePageResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*PdpImage) *PdpImagePageResult
	GetData() []*PdpImage
	SetRequestId(v string) *PdpImagePageResult
	GetRequestId() *string
	SetTotal(v int64) *PdpImagePageResult
	GetTotal() *int64
}

type PdpImagePageResult struct {
	Data      []*PdpImage `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	RequestId *string     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s PdpImagePageResult) String() string {
	return dara.Prettify(s)
}

func (s PdpImagePageResult) GoString() string {
	return s.String()
}

func (s *PdpImagePageResult) GetData() []*PdpImage {
	return s.Data
}

func (s *PdpImagePageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *PdpImagePageResult) GetTotal() *int64 {
	return s.Total
}

func (s *PdpImagePageResult) SetData(v []*PdpImage) *PdpImagePageResult {
	s.Data = v
	return s
}

func (s *PdpImagePageResult) SetRequestId(v string) *PdpImagePageResult {
	s.RequestId = &v
	return s
}

func (s *PdpImagePageResult) SetTotal(v int64) *PdpImagePageResult {
	s.Total = &v
	return s
}

func (s *PdpImagePageResult) Validate() error {
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
