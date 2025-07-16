// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTrafficState(v string) *ReleaseServiceRequest
	GetTrafficState() *string
	SetWeight(v int32) *ReleaseServiceRequest
	GetWeight() *int32
}

type ReleaseServiceRequest struct {
	// The traffic state. Valid values:
	//
	// 	- standalone: independent traffic.
	//
	// 	- grouping: grouped traffic.
	//
	// example:
	//
	// grouping
	TrafficState *string `json:"TrafficState,omitempty" xml:"TrafficState,omitempty"`
	// The weight of the service. Valid values: [-1, 1000].
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ReleaseServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseServiceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseServiceRequest) GetTrafficState() *string {
	return s.TrafficState
}

func (s *ReleaseServiceRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *ReleaseServiceRequest) SetTrafficState(v string) *ReleaseServiceRequest {
	s.TrafficState = &v
	return s
}

func (s *ReleaseServiceRequest) SetWeight(v int32) *ReleaseServiceRequest {
	s.Weight = &v
	return s
}

func (s *ReleaseServiceRequest) Validate() error {
	return dara.Validate(s)
}
