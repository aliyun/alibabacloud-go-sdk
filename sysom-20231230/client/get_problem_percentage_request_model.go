// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProblemPercentageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCluster(v string) *GetProblemPercentageRequest
	GetCluster() *string
	SetEnd(v float32) *GetProblemPercentageRequest
	GetEnd() *float32
	SetInstance(v string) *GetProblemPercentageRequest
	GetInstance() *string
	SetStart(v float32) *GetProblemPercentageRequest
	GetStart() *float32
}

type GetProblemPercentageRequest struct {
	// Cluster ID
	//
	// example:
	//
	// 2ijff4be-bf24-4070-89ca-c47c879b0g32
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// End Time
	//
	// This parameter is required.
	//
	// example:
	//
	// 1725801327754
	End *float32 `json:"end,omitempty" xml:"end,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// Start Time
	//
	// This parameter is required.
	//
	// example:
	//
	// 1725797727754
	Start *float32 `json:"start,omitempty" xml:"start,omitempty"`
}

func (s GetProblemPercentageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProblemPercentageRequest) GoString() string {
	return s.String()
}

func (s *GetProblemPercentageRequest) GetCluster() *string {
	return s.Cluster
}

func (s *GetProblemPercentageRequest) GetEnd() *float32 {
	return s.End
}

func (s *GetProblemPercentageRequest) GetInstance() *string {
	return s.Instance
}

func (s *GetProblemPercentageRequest) GetStart() *float32 {
	return s.Start
}

func (s *GetProblemPercentageRequest) SetCluster(v string) *GetProblemPercentageRequest {
	s.Cluster = &v
	return s
}

func (s *GetProblemPercentageRequest) SetEnd(v float32) *GetProblemPercentageRequest {
	s.End = &v
	return s
}

func (s *GetProblemPercentageRequest) SetInstance(v string) *GetProblemPercentageRequest {
	s.Instance = &v
	return s
}

func (s *GetProblemPercentageRequest) SetStart(v float32) *GetProblemPercentageRequest {
	s.Start = &v
	return s
}

func (s *GetProblemPercentageRequest) Validate() error {
	return dara.Validate(s)
}
