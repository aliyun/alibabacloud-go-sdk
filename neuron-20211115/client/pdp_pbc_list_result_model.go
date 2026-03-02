// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpPbcListResult interface {
	dara.Model
	String() string
	GoString() string
	SetPdpPbcs(v []*PdpPbc) *PdpPbcListResult
	GetPdpPbcs() []*PdpPbc
	SetRequestId(v string) *PdpPbcListResult
	GetRequestId() *string
	SetTotal(v int32) *PdpPbcListResult
	GetTotal() *int32
}

type PdpPbcListResult struct {
	PdpPbcs   []*PdpPbc `json:"pdpPbcs,omitempty" xml:"pdpPbcs,omitempty" type:"Repeated"`
	RequestId *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s PdpPbcListResult) String() string {
	return dara.Prettify(s)
}

func (s PdpPbcListResult) GoString() string {
	return s.String()
}

func (s *PdpPbcListResult) GetPdpPbcs() []*PdpPbc {
	return s.PdpPbcs
}

func (s *PdpPbcListResult) GetRequestId() *string {
	return s.RequestId
}

func (s *PdpPbcListResult) GetTotal() *int32 {
	return s.Total
}

func (s *PdpPbcListResult) SetPdpPbcs(v []*PdpPbc) *PdpPbcListResult {
	s.PdpPbcs = v
	return s
}

func (s *PdpPbcListResult) SetRequestId(v string) *PdpPbcListResult {
	s.RequestId = &v
	return s
}

func (s *PdpPbcListResult) SetTotal(v int32) *PdpPbcListResult {
	s.Total = &v
	return s
}

func (s *PdpPbcListResult) Validate() error {
	if s.PdpPbcs != nil {
		for _, item := range s.PdpPbcs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
