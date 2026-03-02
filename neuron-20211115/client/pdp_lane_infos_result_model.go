// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpLaneInfosResult interface {
	dara.Model
	String() string
	GoString() string
	SetPdpLanes(v []*PdpLaneInfo) *PdpLaneInfosResult
	GetPdpLanes() []*PdpLaneInfo
	SetRequestId(v string) *PdpLaneInfosResult
	GetRequestId() *string
}

type PdpLaneInfosResult struct {
	PdpLanes  []*PdpLaneInfo `json:"pdpLanes,omitempty" xml:"pdpLanes,omitempty" type:"Repeated"`
	RequestId *string        `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PdpLaneInfosResult) String() string {
	return dara.Prettify(s)
}

func (s PdpLaneInfosResult) GoString() string {
	return s.String()
}

func (s *PdpLaneInfosResult) GetPdpLanes() []*PdpLaneInfo {
	return s.PdpLanes
}

func (s *PdpLaneInfosResult) GetRequestId() *string {
	return s.RequestId
}

func (s *PdpLaneInfosResult) SetPdpLanes(v []*PdpLaneInfo) *PdpLaneInfosResult {
	s.PdpLanes = v
	return s
}

func (s *PdpLaneInfosResult) SetRequestId(v string) *PdpLaneInfosResult {
	s.RequestId = &v
	return s
}

func (s *PdpLaneInfosResult) Validate() error {
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
