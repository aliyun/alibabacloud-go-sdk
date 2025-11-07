// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAllCustomizeFlowStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteAllCustomizeFlowStrategyRequest
	GetRegionId() *string
	SetUserId(v int64) *DeleteAllCustomizeFlowStrategyRequest
	GetUserId() *int64
}

type DeleteAllCustomizeFlowStrategyRequest struct {
	// Region ID
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// User ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 126005125163xxxx
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteAllCustomizeFlowStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAllCustomizeFlowStrategyRequest) GoString() string {
	return s.String()
}

func (s *DeleteAllCustomizeFlowStrategyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAllCustomizeFlowStrategyRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *DeleteAllCustomizeFlowStrategyRequest) SetRegionId(v string) *DeleteAllCustomizeFlowStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAllCustomizeFlowStrategyRequest) SetUserId(v int64) *DeleteAllCustomizeFlowStrategyRequest {
	s.UserId = &v
	return s
}

func (s *DeleteAllCustomizeFlowStrategyRequest) Validate() error {
	return dara.Validate(s)
}
