// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScenarioRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteScenarioRequest
	GetRegionId() *string
	SetScenarioId(v int64) *DeleteScenarioRequest
	GetScenarioId() *int64
}

type DeleteScenarioRequest struct {
	// The ID of the region.
	//
	// example:
	//
	// cn-zhangjaikou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the business monitoring job. You can obtain the ID by calling the ListScenario operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 132
	ScenarioId *int64 `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
}

func (s DeleteScenarioRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteScenarioRequest) GoString() string {
	return s.String()
}

func (s *DeleteScenarioRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteScenarioRequest) GetScenarioId() *int64 {
	return s.ScenarioId
}

func (s *DeleteScenarioRequest) SetRegionId(v string) *DeleteScenarioRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteScenarioRequest) SetScenarioId(v int64) *DeleteScenarioRequest {
	s.ScenarioId = &v
	return s
}

func (s *DeleteScenarioRequest) Validate() error {
	return dara.Validate(s)
}
