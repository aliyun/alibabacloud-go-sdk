// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppVulScanCycleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCycle(v string) *ModifyAppVulScanCycleRequest
	GetCycle() *string
}

type ModifyAppVulScanCycleRequest struct {
	// The scan cycle for application vulnerabilities.
	//
	// 	- 1week
	//
	// 	- 2weeks
	//
	// 	- 3days
	//
	// example:
	//
	// 1week
	Cycle *string `json:"Cycle,omitempty" xml:"Cycle,omitempty"`
}

func (s ModifyAppVulScanCycleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppVulScanCycleRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppVulScanCycleRequest) GetCycle() *string {
	return s.Cycle
}

func (s *ModifyAppVulScanCycleRequest) SetCycle(v string) *ModifyAppVulScanCycleRequest {
	s.Cycle = &v
	return s
}

func (s *ModifyAppVulScanCycleRequest) Validate() error {
	return dara.Validate(s)
}
