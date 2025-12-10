// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTunnelQos interface {
	dara.Model
	String() string
	GoString() string
	SetMaxBandwidth(v int64) *TunnelQos
	GetMaxBandwidth() *int64
	SetMaxQps(v int32) *TunnelQos
	GetMaxQps() *int32
}

type TunnelQos struct {
	// example:
	//
	// 1073741824
	MaxBandwidth *int64 `json:"MaxBandwidth,omitempty" xml:"MaxBandwidth,omitempty"`
	// example:
	//
	// 100
	MaxQps *int32 `json:"MaxQps,omitempty" xml:"MaxQps,omitempty"`
}

func (s TunnelQos) String() string {
	return dara.Prettify(s)
}

func (s TunnelQos) GoString() string {
	return s.String()
}

func (s *TunnelQos) GetMaxBandwidth() *int64 {
	return s.MaxBandwidth
}

func (s *TunnelQos) GetMaxQps() *int32 {
	return s.MaxQps
}

func (s *TunnelQos) SetMaxBandwidth(v int64) *TunnelQos {
	s.MaxBandwidth = &v
	return s
}

func (s *TunnelQos) SetMaxQps(v int32) *TunnelQos {
	s.MaxQps = &v
	return s
}

func (s *TunnelQos) Validate() error {
	return dara.Validate(s)
}
