// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDataAgentMemoryConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDMSUnit(v string) *CheckDataAgentMemoryConfigRequest
	GetDMSUnit() *string
}

type CheckDataAgentMemoryConfigRequest struct {
	// The current Data Management unit.
	//
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
}

func (s CheckDataAgentMemoryConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckDataAgentMemoryConfigRequest) GoString() string {
	return s.String()
}

func (s *CheckDataAgentMemoryConfigRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *CheckDataAgentMemoryConfigRequest) SetDMSUnit(v string) *CheckDataAgentMemoryConfigRequest {
	s.DMSUnit = &v
	return s
}

func (s *CheckDataAgentMemoryConfigRequest) Validate() error {
	return dara.Validate(s)
}
