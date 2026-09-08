// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLFlowWaste interface {
	dara.Model
	String() string
	GoString() string
	SetUsefulSec(v int64) *RLFlowWaste
	GetUsefulSec() *int64
}

type RLFlowWaste struct {
	// The cumulative duration of trained trajectories, in seconds.
	//
	// example:
	//
	// 183
	UsefulSec *int64 `json:"UsefulSec,omitempty" xml:"UsefulSec,omitempty"`
}

func (s RLFlowWaste) String() string {
	return dara.Prettify(s)
}

func (s RLFlowWaste) GoString() string {
	return s.String()
}

func (s *RLFlowWaste) GetUsefulSec() *int64 {
	return s.UsefulSec
}

func (s *RLFlowWaste) SetUsefulSec(v int64) *RLFlowWaste {
	s.UsefulSec = &v
	return s
}

func (s *RLFlowWaste) Validate() error {
	return dara.Validate(s)
}
