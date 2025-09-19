// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImageFixCycleConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageFixCycle(v int32) *ModifyImageFixCycleConfigRequest
	GetImageFixCycle() *int32
	SetImageFixSwitch(v string) *ModifyImageFixCycleConfigRequest
	GetImageFixSwitch() *string
	SetImageFixTarget(v string) *ModifyImageFixCycleConfigRequest
	GetImageFixTarget() *string
	SetImageTimeRange(v int32) *ModifyImageFixCycleConfigRequest
	GetImageTimeRange() *int32
}

type ModifyImageFixCycleConfigRequest struct {
	// The cycle of the scheduled fix. Unit: day.
	//
	// example:
	//
	// 7
	ImageFixCycle *int32 `json:"ImageFixCycle,omitempty" xml:"ImageFixCycle,omitempty"`
	// Specifies whether to enable the schedule image fix.
	//
	// 	- **on**: enable
	//
	// 	- **off**: disable
	//
	// example:
	//
	// on
	ImageFixSwitch *string `json:"ImageFixSwitch,omitempty" xml:"ImageFixSwitch,omitempty"`
	// The range of the scheduled fix. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **type**: The type of the image. The value is fixed to repo.
	//
	// 	- **target**: The content of the image. The value is in the format of Namespace/Image repository.
	//
	// example:
	//
	// {\\"type\\":\\"repo\\",\\"target\\":[\\"cdp-uat/zentao\\",\\"qa-dac/yyuan9\\",\\"cafdms-qa/xxl-job-admin\\"]}
	ImageFixTarget *string `json:"ImageFixTarget,omitempty" xml:"ImageFixTarget,omitempty"`
	// The time range during which the image was modified. Unit: day.
	//
	// example:
	//
	// 30
	ImageTimeRange *int32 `json:"ImageTimeRange,omitempty" xml:"ImageTimeRange,omitempty"`
}

func (s ModifyImageFixCycleConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageFixCycleConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyImageFixCycleConfigRequest) GetImageFixCycle() *int32 {
	return s.ImageFixCycle
}

func (s *ModifyImageFixCycleConfigRequest) GetImageFixSwitch() *string {
	return s.ImageFixSwitch
}

func (s *ModifyImageFixCycleConfigRequest) GetImageFixTarget() *string {
	return s.ImageFixTarget
}

func (s *ModifyImageFixCycleConfigRequest) GetImageTimeRange() *int32 {
	return s.ImageTimeRange
}

func (s *ModifyImageFixCycleConfigRequest) SetImageFixCycle(v int32) *ModifyImageFixCycleConfigRequest {
	s.ImageFixCycle = &v
	return s
}

func (s *ModifyImageFixCycleConfigRequest) SetImageFixSwitch(v string) *ModifyImageFixCycleConfigRequest {
	s.ImageFixSwitch = &v
	return s
}

func (s *ModifyImageFixCycleConfigRequest) SetImageFixTarget(v string) *ModifyImageFixCycleConfigRequest {
	s.ImageFixTarget = &v
	return s
}

func (s *ModifyImageFixCycleConfigRequest) SetImageTimeRange(v int32) *ModifyImageFixCycleConfigRequest {
	s.ImageTimeRange = &v
	return s
}

func (s *ModifyImageFixCycleConfigRequest) Validate() error {
	return dara.Validate(s)
}
