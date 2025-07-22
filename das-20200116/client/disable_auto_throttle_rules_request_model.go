// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAutoThrottleRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleContext(v string) *DisableAutoThrottleRulesRequest
	GetConsoleContext() *string
	SetInstanceIds(v string) *DisableAutoThrottleRulesRequest
	GetInstanceIds() *string
}

type DisableAutoThrottleRulesRequest struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	// The database instance IDs.
	//
	// >  Set this parameter to a JSON array that consists of multiple instance IDs. Separate instance IDs with commas (,). Example: `[\\"Instance ID1\\",\\"Instance ID2\\"]`.
	//
	// This parameter is required.
	//
	// example:
	//
	// [\\"rm-2ze8g2am97624****\\",\\"rm-2ze9xrhze0709****\\"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s DisableAutoThrottleRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableAutoThrottleRulesRequest) GoString() string {
	return s.String()
}

func (s *DisableAutoThrottleRulesRequest) GetConsoleContext() *string {
	return s.ConsoleContext
}

func (s *DisableAutoThrottleRulesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DisableAutoThrottleRulesRequest) SetConsoleContext(v string) *DisableAutoThrottleRulesRequest {
	s.ConsoleContext = &v
	return s
}

func (s *DisableAutoThrottleRulesRequest) SetInstanceIds(v string) *DisableAutoThrottleRulesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DisableAutoThrottleRulesRequest) Validate() error {
	return dara.Validate(s)
}
