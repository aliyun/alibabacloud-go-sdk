// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgScenedDeleteSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v []*int32) *DsgScenedDeleteSceneRequest
	GetIds() []*int32
}

type DsgScenedDeleteSceneRequest struct {
	// The IDs of level-2 data masking scenarios.
	//
	// This parameter is required.
	Ids []*int32 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
}

func (s DsgScenedDeleteSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgScenedDeleteSceneRequest) GoString() string {
	return s.String()
}

func (s *DsgScenedDeleteSceneRequest) GetIds() []*int32 {
	return s.Ids
}

func (s *DsgScenedDeleteSceneRequest) SetIds(v []*int32) *DsgScenedDeleteSceneRequest {
	s.Ids = v
	return s
}

func (s *DsgScenedDeleteSceneRequest) Validate() error {
	return dara.Validate(s)
}
