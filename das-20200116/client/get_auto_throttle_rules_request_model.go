// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoThrottleRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleContext(v string) *GetAutoThrottleRulesRequest
	GetConsoleContext() *string
	SetInstanceIds(v string) *GetAutoThrottleRulesRequest
	GetInstanceIds() *string
}

type GetAutoThrottleRulesRequest struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	// The database instance IDs.
	//
	// 	- Set this parameter to a JSON array that consists of multiple instance IDs. Separate instance IDs with commas (,). Example: `[\\"Instance ID1\\",\\"Instance ID2\\"]`.
	//
	// 	- By default, if you leave this parameter empty, all database instances for which the automatic SQL throttling feature has been enabled within the current Alibaba Cloud account are returned. The following types of database instances are returned:
	//
	//     	- Database instances for which the automatic SQL throttling feature is currently enabled.
	//
	//     	- Database instances for which the automatic SQL throttling feature was once enabled but is currently disabled. Released database instances are not included.
	//
	// example:
	//
	// [\\"rm-2ze8g2am97624****\\",\\"rm-2vc54m2a6pd6p****\\",\\"rm-2ze9xrhze0709****\\"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s GetAutoThrottleRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAutoThrottleRulesRequest) GoString() string {
	return s.String()
}

func (s *GetAutoThrottleRulesRequest) GetConsoleContext() *string {
	return s.ConsoleContext
}

func (s *GetAutoThrottleRulesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *GetAutoThrottleRulesRequest) SetConsoleContext(v string) *GetAutoThrottleRulesRequest {
	s.ConsoleContext = &v
	return s
}

func (s *GetAutoThrottleRulesRequest) SetInstanceIds(v string) *GetAutoThrottleRulesRequest {
	s.InstanceIds = &v
	return s
}

func (s *GetAutoThrottleRulesRequest) Validate() error {
	return dara.Validate(s)
}
