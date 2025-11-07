// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBlackListStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *QueryBlackListStrategyRequest
	GetRegionId() *string
}

type QueryBlackListStrategyRequest struct {
	// Region ID
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s QueryBlackListStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryBlackListStrategyRequest) GoString() string {
	return s.String()
}

func (s *QueryBlackListStrategyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryBlackListStrategyRequest) SetRegionId(v string) *QueryBlackListStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *QueryBlackListStrategyRequest) Validate() error {
	return dara.Validate(s)
}
