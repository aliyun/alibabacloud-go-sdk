// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSourceMapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFidList(v []*string) *DeleteSourceMapRequest
	GetFidList() []*string
	SetPid(v string) *DeleteSourceMapRequest
	GetPid() *string
	SetRegionId(v string) *DeleteSourceMapRequest
	GetRegionId() *string
}

type DeleteSourceMapRequest struct {
	// The IDs of the SourceMap files.
	//
	// This parameter is required.
	FidList []*string `json:"FidList,omitempty" xml:"FidList,omitempty" type:"Repeated"`
	// The process identifier (PID) of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// atc889zkcf@d8deedfa9bf****
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteSourceMapRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSourceMapRequest) GoString() string {
	return s.String()
}

func (s *DeleteSourceMapRequest) GetFidList() []*string {
	return s.FidList
}

func (s *DeleteSourceMapRequest) GetPid() *string {
	return s.Pid
}

func (s *DeleteSourceMapRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSourceMapRequest) SetFidList(v []*string) *DeleteSourceMapRequest {
	s.FidList = v
	return s
}

func (s *DeleteSourceMapRequest) SetPid(v string) *DeleteSourceMapRequest {
	s.Pid = &v
	return s
}

func (s *DeleteSourceMapRequest) SetRegionId(v string) *DeleteSourceMapRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSourceMapRequest) Validate() error {
	return dara.Validate(s)
}
