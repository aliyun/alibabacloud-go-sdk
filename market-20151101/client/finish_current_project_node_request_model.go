// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishCurrentProjectNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *FinishCurrentProjectNodeRequest
	GetInstanceId() *string
	SetNodeId(v int64) *FinishCurrentProjectNodeRequest
	GetNodeId() *int64
	SetRemark(v string) *FinishCurrentProjectNodeRequest
	GetRemark() *string
	SetTemplateForm(v string) *FinishCurrentProjectNodeRequest
	GetTemplateForm() *string
}

type FinishCurrentProjectNodeRequest struct {
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
	// 1924
	NodeId       *int64  `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	TemplateForm *string `json:"TemplateForm,omitempty" xml:"TemplateForm,omitempty"`
}

func (s FinishCurrentProjectNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s FinishCurrentProjectNodeRequest) GoString() string {
	return s.String()
}

func (s *FinishCurrentProjectNodeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *FinishCurrentProjectNodeRequest) GetNodeId() *int64 {
	return s.NodeId
}

func (s *FinishCurrentProjectNodeRequest) GetRemark() *string {
	return s.Remark
}

func (s *FinishCurrentProjectNodeRequest) GetTemplateForm() *string {
	return s.TemplateForm
}

func (s *FinishCurrentProjectNodeRequest) SetInstanceId(v string) *FinishCurrentProjectNodeRequest {
	s.InstanceId = &v
	return s
}

func (s *FinishCurrentProjectNodeRequest) SetNodeId(v int64) *FinishCurrentProjectNodeRequest {
	s.NodeId = &v
	return s
}

func (s *FinishCurrentProjectNodeRequest) SetRemark(v string) *FinishCurrentProjectNodeRequest {
	s.Remark = &v
	return s
}

func (s *FinishCurrentProjectNodeRequest) SetTemplateForm(v string) *FinishCurrentProjectNodeRequest {
	s.TemplateForm = &v
	return s
}

func (s *FinishCurrentProjectNodeRequest) Validate() error {
	return dara.Validate(s)
}
