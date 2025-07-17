// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPids(v []*string) *DeleteAppListRequest
	GetPids() []*string
	SetRegionId(v string) *DeleteAppListRequest
	GetRegionId() *string
}

type DeleteAppListRequest struct {
	// The list of PIDs for the applications monitored by Application Monitoring.
	Pids []*string `json:"Pids,omitempty" xml:"Pids,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAppListRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppListRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppListRequest) GetPids() []*string {
	return s.Pids
}

func (s *DeleteAppListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAppListRequest) SetPids(v []*string) *DeleteAppListRequest {
	s.Pids = v
	return s
}

func (s *DeleteAppListRequest) SetRegionId(v string) *DeleteAppListRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAppListRequest) Validate() error {
	return dara.Validate(s)
}
