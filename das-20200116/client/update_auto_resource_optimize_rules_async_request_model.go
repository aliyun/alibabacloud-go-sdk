// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoResourceOptimizeRulesAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleContext(v string) *UpdateAutoResourceOptimizeRulesAsyncRequest
	GetConsoleContext() *string
	SetInstanceIds(v string) *UpdateAutoResourceOptimizeRulesAsyncRequest
	GetInstanceIds() *string
	SetResultId(v string) *UpdateAutoResourceOptimizeRulesAsyncRequest
	GetResultId() *string
	SetTableFragmentationRatio(v float64) *UpdateAutoResourceOptimizeRulesAsyncRequest
	GetTableFragmentationRatio() *float64
	SetTableSpaceSize(v float64) *UpdateAutoResourceOptimizeRulesAsyncRequest
	GetTableSpaceSize() *float64
}

type UpdateAutoResourceOptimizeRulesAsyncRequest struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	// The database instance IDs.
	//
	// >  Set this parameter to a JSON array that consists of multiple instance IDs. Separate instance IDs with commas (,). Example: `[\\"Instance ID1\\", \\"Instance ID2\\"]`.
	//
	// This parameter is required.
	//
	// example:
	//
	// [\\"rm-2ze8g2am97624****\\",\\"rm-2ze9xrhze0709****\\"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The ID of the asynchronous request.
	//
	// >  Asynchronous calls do not immediately return the complete results. To obtain the complete results, you must use the value of **ResultId*	- returned in the response to re-initiate the call until the value of **isFinish*	- is **true**.***	- In this case, you must call this operation at least twice.
	//
	// example:
	//
	// async__507044db6c4eadfa2dab9b084e80****
	ResultId *string `json:"ResultId,omitempty" xml:"ResultId,omitempty"`
	// The fragmentation rate that triggers automatic fragment recycling of a single physical table. Valid values: **0.10*	- to **0.99**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	TableFragmentationRatio *float64 `json:"TableFragmentationRatio,omitempty" xml:"TableFragmentationRatio,omitempty"`
	// The minimum storage usage that triggers automatic fragment recycling of a single physical table. Valid values: **5*	- to **100**. Unit: GB.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	TableSpaceSize *float64 `json:"TableSpaceSize,omitempty" xml:"TableSpaceSize,omitempty"`
}

func (s UpdateAutoResourceOptimizeRulesAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoResourceOptimizeRulesAsyncRequest) GoString() string {
	return s.String()
}

func (s *UpdateAutoResourceOptimizeRulesAsyncRequest) GetConsoleContext() *string {
	return s.ConsoleContext
}

func (s *UpdateAutoResourceOptimizeRulesAsyncRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *UpdateAutoResourceOptimizeRulesAsyncRequest) GetResultId() *string {
	return s.ResultId
}

func (s *UpdateAutoResourceOptimizeRulesAsyncRequest) GetTableFragmentationRatio() *float64 {
	return s.TableFragmentationRatio
}

func (s *UpdateAutoResourceOptimizeRulesAsyncRequest) GetTableSpaceSize() *float64 {
	return s.TableSpaceSize
}

func (s *UpdateAutoResourceOptimizeRulesAsyncRequest) SetConsoleContext(v string) *UpdateAutoResourceOptimizeRulesAsyncRequest {
	s.ConsoleContext = &v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncRequest) SetInstanceIds(v string) *UpdateAutoResourceOptimizeRulesAsyncRequest {
	s.InstanceIds = &v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncRequest) SetResultId(v string) *UpdateAutoResourceOptimizeRulesAsyncRequest {
	s.ResultId = &v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncRequest) SetTableFragmentationRatio(v float64) *UpdateAutoResourceOptimizeRulesAsyncRequest {
	s.TableFragmentationRatio = &v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncRequest) SetTableSpaceSize(v float64) *UpdateAutoResourceOptimizeRulesAsyncRequest {
	s.TableSpaceSize = &v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncRequest) Validate() error {
	return dara.Validate(s)
}
