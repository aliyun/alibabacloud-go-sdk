// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSourceMapShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFidListShrink(v string) *DeleteSourceMapShrinkRequest
	GetFidListShrink() *string
	SetPid(v string) *DeleteSourceMapShrinkRequest
	GetPid() *string
	SetRegionId(v string) *DeleteSourceMapShrinkRequest
	GetRegionId() *string
}

type DeleteSourceMapShrinkRequest struct {
	// The IDs of the SourceMap files.
	//
	// This parameter is required.
	FidListShrink *string `json:"FidList,omitempty" xml:"FidList,omitempty"`
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

func (s DeleteSourceMapShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSourceMapShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteSourceMapShrinkRequest) GetFidListShrink() *string {
	return s.FidListShrink
}

func (s *DeleteSourceMapShrinkRequest) GetPid() *string {
	return s.Pid
}

func (s *DeleteSourceMapShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSourceMapShrinkRequest) SetFidListShrink(v string) *DeleteSourceMapShrinkRequest {
	s.FidListShrink = &v
	return s
}

func (s *DeleteSourceMapShrinkRequest) SetPid(v string) *DeleteSourceMapShrinkRequest {
	s.Pid = &v
	return s
}

func (s *DeleteSourceMapShrinkRequest) SetRegionId(v string) *DeleteSourceMapShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSourceMapShrinkRequest) Validate() error {
	return dara.Validate(s)
}
