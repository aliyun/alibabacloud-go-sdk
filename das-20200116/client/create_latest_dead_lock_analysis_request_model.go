// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLatestDeadLockAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateLatestDeadLockAnalysisRequest
	GetInstanceId() *string
	SetNodeId(v string) *CreateLatestDeadLockAnalysisRequest
	GetNodeId() *string
}

type CreateLatestDeadLockAnalysisRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1u5mas9exx7****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// pi-bp16v3824rt73****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s CreateLatestDeadLockAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLatestDeadLockAnalysisRequest) GoString() string {
	return s.String()
}

func (s *CreateLatestDeadLockAnalysisRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateLatestDeadLockAnalysisRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateLatestDeadLockAnalysisRequest) SetInstanceId(v string) *CreateLatestDeadLockAnalysisRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateLatestDeadLockAnalysisRequest) SetNodeId(v string) *CreateLatestDeadLockAnalysisRequest {
	s.NodeId = &v
	return s
}

func (s *CreateLatestDeadLockAnalysisRequest) Validate() error {
	return dara.Validate(s)
}
