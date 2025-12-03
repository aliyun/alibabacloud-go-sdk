// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeNodeCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ResizeNodeCountRequest
	GetClusterId() *string
	SetNodeCount(v int32) *ResizeNodeCountRequest
	GetNodeCount() *int32
	SetVSwitchId(v string) *ResizeNodeCountRequest
	GetVSwitchId() *string
	SetZoneId(v string) *ResizeNodeCountRequest
	GetZoneId() *string
}

type ResizeNodeCountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-bp16o0pd52e3y****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// example:
	//
	// vsw-bp191otqj1ssyl****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ResizeNodeCountRequest) String() string {
	return dara.Prettify(s)
}

func (s ResizeNodeCountRequest) GoString() string {
	return s.String()
}

func (s *ResizeNodeCountRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ResizeNodeCountRequest) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *ResizeNodeCountRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ResizeNodeCountRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ResizeNodeCountRequest) SetClusterId(v string) *ResizeNodeCountRequest {
	s.ClusterId = &v
	return s
}

func (s *ResizeNodeCountRequest) SetNodeCount(v int32) *ResizeNodeCountRequest {
	s.NodeCount = &v
	return s
}

func (s *ResizeNodeCountRequest) SetVSwitchId(v string) *ResizeNodeCountRequest {
	s.VSwitchId = &v
	return s
}

func (s *ResizeNodeCountRequest) SetZoneId(v string) *ResizeNodeCountRequest {
	s.ZoneId = &v
	return s
}

func (s *ResizeNodeCountRequest) Validate() error {
	return dara.Validate(s)
}
