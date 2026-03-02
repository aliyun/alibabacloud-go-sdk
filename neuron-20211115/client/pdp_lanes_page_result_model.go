// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpLanesPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetPdpLanes(v []*PdpLaneInfo) *PdpLanesPageResult
	GetPdpLanes() []*PdpLaneInfo
	SetRequestId(v string) *PdpLanesPageResult
	GetRequestId() *string
	SetTotal(v int32) *PdpLanesPageResult
	GetTotal() *int32
}

type PdpLanesPageResult struct {
	PdpLanes []*PdpLaneInfo `json:"pdpLanes,omitempty" xml:"pdpLanes,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s PdpLanesPageResult) String() string {
	return dara.Prettify(s)
}

func (s PdpLanesPageResult) GoString() string {
	return s.String()
}

func (s *PdpLanesPageResult) GetPdpLanes() []*PdpLaneInfo {
	return s.PdpLanes
}

func (s *PdpLanesPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *PdpLanesPageResult) GetTotal() *int32 {
	return s.Total
}

func (s *PdpLanesPageResult) SetPdpLanes(v []*PdpLaneInfo) *PdpLanesPageResult {
	s.PdpLanes = v
	return s
}

func (s *PdpLanesPageResult) SetRequestId(v string) *PdpLanesPageResult {
	s.RequestId = &v
	return s
}

func (s *PdpLanesPageResult) SetTotal(v int32) *PdpLanesPageResult {
	s.Total = &v
	return s
}

func (s *PdpLanesPageResult) Validate() error {
	if s.PdpLanes != nil {
		for _, item := range s.PdpLanes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
