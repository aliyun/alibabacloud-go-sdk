// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyCheckResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckIds(v []*int64) *VerifyCheckResultRequest
	GetCheckIds() []*int64
	SetForce(v bool) *VerifyCheckResultRequest
	GetForce() *bool
	SetInstanceIds(v []*string) *VerifyCheckResultRequest
	GetInstanceIds() []*string
	SetTaskSource(v string) *VerifyCheckResultRequest
	GetTaskSource() *string
}

type VerifyCheckResultRequest struct {
	// The list of check item IDs.
	CheckIds []*int64 `json:"CheckIds,omitempty" xml:"CheckIds,omitempty" type:"Repeated"`
	// Specifies whether to forcibly execute the specified check item. Default value: false.
	//
	// - true: Forcibly executes the specified check item. Forcible execution bypasses frequency and quantity limits and directly initiates the check, which may result in repeated executions within a short period.
	//
	// - false (default): Does not forcibly execute the specified check item. This ensures that the same check item is executed only once within a short period.
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The list of instance IDs for the check item assets.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The source of the task. Valid values:
	//
	// - **YAO_CHI**: ApsaraDB console.
	//
	// example:
	//
	// YAO_CHI
	TaskSource *string `json:"TaskSource,omitempty" xml:"TaskSource,omitempty"`
}

func (s VerifyCheckResultRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyCheckResultRequest) GoString() string {
	return s.String()
}

func (s *VerifyCheckResultRequest) GetCheckIds() []*int64 {
	return s.CheckIds
}

func (s *VerifyCheckResultRequest) GetForce() *bool {
	return s.Force
}

func (s *VerifyCheckResultRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *VerifyCheckResultRequest) GetTaskSource() *string {
	return s.TaskSource
}

func (s *VerifyCheckResultRequest) SetCheckIds(v []*int64) *VerifyCheckResultRequest {
	s.CheckIds = v
	return s
}

func (s *VerifyCheckResultRequest) SetForce(v bool) *VerifyCheckResultRequest {
	s.Force = &v
	return s
}

func (s *VerifyCheckResultRequest) SetInstanceIds(v []*string) *VerifyCheckResultRequest {
	s.InstanceIds = v
	return s
}

func (s *VerifyCheckResultRequest) SetTaskSource(v string) *VerifyCheckResultRequest {
	s.TaskSource = &v
	return s
}

func (s *VerifyCheckResultRequest) Validate() error {
	return dara.Validate(s)
}
