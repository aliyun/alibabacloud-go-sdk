// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUnknownThreatDetectStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdList(v []*int64) *DeleteUnknownThreatDetectStrategyRequest
	GetIdList() []*int64
}

type DeleteUnknownThreatDetectStrategyRequest struct {
	// example:
	//
	// 123
	IdList []*int64 `json:"IdList,omitempty" xml:"IdList,omitempty" type:"Repeated"`
}

func (s DeleteUnknownThreatDetectStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUnknownThreatDetectStrategyRequest) GoString() string {
	return s.String()
}

func (s *DeleteUnknownThreatDetectStrategyRequest) GetIdList() []*int64 {
	return s.IdList
}

func (s *DeleteUnknownThreatDetectStrategyRequest) SetIdList(v []*int64) *DeleteUnknownThreatDetectStrategyRequest {
	s.IdList = v
	return s
}

func (s *DeleteUnknownThreatDetectStrategyRequest) Validate() error {
	return dara.Validate(s)
}
