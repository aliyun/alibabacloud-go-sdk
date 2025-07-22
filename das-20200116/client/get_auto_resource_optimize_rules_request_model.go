// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoResourceOptimizeRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleContext(v string) *GetAutoResourceOptimizeRulesRequest
	GetConsoleContext() *string
	SetInstanceIds(v string) *GetAutoResourceOptimizeRulesRequest
	GetInstanceIds() *string
}

type GetAutoResourceOptimizeRulesRequest struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	// The database instance IDs.
	//
	// 	- Specify the parameter value as a JSON array, such as `[\\"Database account 1\\",\\"Database account 2\\"]`. Separate database instance IDs with commas (,).
	//
	// 	- By default, if you leave this parameter empty, all database instances for which the automatic fragment recycling feature has been enabled within the current Alibaba Cloud account are returned. The following types of database instances are returned:
	//
	//     	- Database instances for which the automatic fragment recycling feature is currently enabled.
	//
	//     	- Database instances for which the automatic fragment recycling feature was once enabled but is currently disabled, including those for which DAS Enterprise Edition has been disabled but excluding those that have been released.
	//
	// example:
	//
	// [\\"rm-2ze8g2am97624****\\",\\"rm-2vc54m2a6pd6p****\\",\\"rm-2ze9xrhze0709****\\",\\"rm-2ze8g2am97627****\\"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s GetAutoResourceOptimizeRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAutoResourceOptimizeRulesRequest) GoString() string {
	return s.String()
}

func (s *GetAutoResourceOptimizeRulesRequest) GetConsoleContext() *string {
	return s.ConsoleContext
}

func (s *GetAutoResourceOptimizeRulesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *GetAutoResourceOptimizeRulesRequest) SetConsoleContext(v string) *GetAutoResourceOptimizeRulesRequest {
	s.ConsoleContext = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesRequest) SetInstanceIds(v string) *GetAutoResourceOptimizeRulesRequest {
	s.InstanceIds = &v
	return s
}

func (s *GetAutoResourceOptimizeRulesRequest) Validate() error {
	return dara.Validate(s)
}
