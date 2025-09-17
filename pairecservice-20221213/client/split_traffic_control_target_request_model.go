// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSplitTrafficControlTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironment(v string) *SplitTrafficControlTargetRequest
	GetEnvironment() *string
	SetInstanceId(v string) *SplitTrafficControlTargetRequest
	GetInstanceId() *string
	SetSetPoints(v []*int64) *SplitTrafficControlTargetRequest
	GetSetPoints() []*int64
	SetSetValues(v []*int64) *SplitTrafficControlTargetRequest
	GetSetValues() []*int64
	SetTimePoints(v []*int64) *SplitTrafficControlTargetRequest
	GetTimePoints() []*int64
}

type SplitTrafficControlTargetRequest struct {
	// example:
	//
	// Prod
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// example:
	//
	// learn-pairec-xxx
	InstanceId *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SetPoints  []*int64 `json:"SetPoints,omitempty" xml:"SetPoints,omitempty" type:"Repeated"`
	SetValues  []*int64 `json:"SetValues,omitempty" xml:"SetValues,omitempty" type:"Repeated"`
	TimePoints []*int64 `json:"TimePoints,omitempty" xml:"TimePoints,omitempty" type:"Repeated"`
}

func (s SplitTrafficControlTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s SplitTrafficControlTargetRequest) GoString() string {
	return s.String()
}

func (s *SplitTrafficControlTargetRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *SplitTrafficControlTargetRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SplitTrafficControlTargetRequest) GetSetPoints() []*int64 {
	return s.SetPoints
}

func (s *SplitTrafficControlTargetRequest) GetSetValues() []*int64 {
	return s.SetValues
}

func (s *SplitTrafficControlTargetRequest) GetTimePoints() []*int64 {
	return s.TimePoints
}

func (s *SplitTrafficControlTargetRequest) SetEnvironment(v string) *SplitTrafficControlTargetRequest {
	s.Environment = &v
	return s
}

func (s *SplitTrafficControlTargetRequest) SetInstanceId(v string) *SplitTrafficControlTargetRequest {
	s.InstanceId = &v
	return s
}

func (s *SplitTrafficControlTargetRequest) SetSetPoints(v []*int64) *SplitTrafficControlTargetRequest {
	s.SetPoints = v
	return s
}

func (s *SplitTrafficControlTargetRequest) SetSetValues(v []*int64) *SplitTrafficControlTargetRequest {
	s.SetValues = v
	return s
}

func (s *SplitTrafficControlTargetRequest) SetTimePoints(v []*int64) *SplitTrafficControlTargetRequest {
	s.TimePoints = v
	return s
}

func (s *SplitTrafficControlTargetRequest) Validate() error {
	return dara.Validate(s)
}
