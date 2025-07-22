// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAutoResourceOptimizeRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleContext(v string) *DisableAutoResourceOptimizeRulesRequest
	GetConsoleContext() *string
	SetInstanceIds(v string) *DisableAutoResourceOptimizeRulesRequest
	GetInstanceIds() *string
}

type DisableAutoResourceOptimizeRulesRequest struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	// The database instance ID.
	//
	// >  Set this parameter to a JSON array that consists of multiple instance IDs. Separate instance IDs with commas (,). Example: `[\\"Instance ID1\\", \\"Instance ID2\\"]`.
	//
	// This parameter is required.
	//
	// example:
	//
	// [\\"rm-2ze8g2am97624****\\",\\"rm-2ze9xrhze0709****\\"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s DisableAutoResourceOptimizeRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableAutoResourceOptimizeRulesRequest) GoString() string {
	return s.String()
}

func (s *DisableAutoResourceOptimizeRulesRequest) GetConsoleContext() *string {
	return s.ConsoleContext
}

func (s *DisableAutoResourceOptimizeRulesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DisableAutoResourceOptimizeRulesRequest) SetConsoleContext(v string) *DisableAutoResourceOptimizeRulesRequest {
	s.ConsoleContext = &v
	return s
}

func (s *DisableAutoResourceOptimizeRulesRequest) SetInstanceIds(v string) *DisableAutoResourceOptimizeRulesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DisableAutoResourceOptimizeRulesRequest) Validate() error {
	return dara.Validate(s)
}
