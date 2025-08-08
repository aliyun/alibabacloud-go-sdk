// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ResumeProjectRequest
	GetInstanceId() *string
	SetNodeId(v int64) *ResumeProjectRequest
	GetNodeId() *int64
	SetRemark(v string) *ResumeProjectRequest
	GetRemark() *string
}

type ResumeProjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4****89
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1922
	NodeId *int64  `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s ResumeProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeProjectRequest) GoString() string {
	return s.String()
}

func (s *ResumeProjectRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResumeProjectRequest) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ResumeProjectRequest) GetRemark() *string {
	return s.Remark
}

func (s *ResumeProjectRequest) SetInstanceId(v string) *ResumeProjectRequest {
	s.InstanceId = &v
	return s
}

func (s *ResumeProjectRequest) SetNodeId(v int64) *ResumeProjectRequest {
	s.NodeId = &v
	return s
}

func (s *ResumeProjectRequest) SetRemark(v string) *ResumeProjectRequest {
	s.Remark = &v
	return s
}

func (s *ResumeProjectRequest) Validate() error {
	return dara.Validate(s)
}
