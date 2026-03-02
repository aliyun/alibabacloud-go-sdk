// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcListResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*Pbc) *PbcListResult
	GetData() []*Pbc
	SetRequestId(v string) *PbcListResult
	GetRequestId() *string
	SetTotal(v int32) *PbcListResult
	GetTotal() *int32
}

type PbcListResult struct {
	Data      []*Pbc  `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s PbcListResult) String() string {
	return dara.Prettify(s)
}

func (s PbcListResult) GoString() string {
	return s.String()
}

func (s *PbcListResult) GetData() []*Pbc {
	return s.Data
}

func (s *PbcListResult) GetRequestId() *string {
	return s.RequestId
}

func (s *PbcListResult) GetTotal() *int32 {
	return s.Total
}

func (s *PbcListResult) SetData(v []*Pbc) *PbcListResult {
	s.Data = v
	return s
}

func (s *PbcListResult) SetRequestId(v string) *PbcListResult {
	s.RequestId = &v
	return s
}

func (s *PbcListResult) SetTotal(v int32) *PbcListResult {
	s.Total = &v
	return s
}

func (s *PbcListResult) Validate() error {
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
