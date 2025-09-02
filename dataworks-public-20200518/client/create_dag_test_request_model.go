// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDagTestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizdate(v string) *CreateDagTestRequest
	GetBizdate() *string
	SetName(v string) *CreateDagTestRequest
	GetName() *string
	SetNodeId(v int64) *CreateDagTestRequest
	GetNodeId() *int64
	SetNodeParams(v string) *CreateDagTestRequest
	GetNodeParams() *string
	SetProjectEnv(v string) *CreateDagTestRequest
	GetProjectEnv() *string
}

type CreateDagTestRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2020-05-26 00:00:00
	Bizdate *string `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xm_create_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// bizdate=$bizdate tbods=$tbods
	NodeParams *string `json:"NodeParams,omitempty" xml:"NodeParams,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROD
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
}

func (s CreateDagTestRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDagTestRequest) GoString() string {
	return s.String()
}

func (s *CreateDagTestRequest) GetBizdate() *string {
	return s.Bizdate
}

func (s *CreateDagTestRequest) GetName() *string {
	return s.Name
}

func (s *CreateDagTestRequest) GetNodeId() *int64 {
	return s.NodeId
}

func (s *CreateDagTestRequest) GetNodeParams() *string {
	return s.NodeParams
}

func (s *CreateDagTestRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *CreateDagTestRequest) SetBizdate(v string) *CreateDagTestRequest {
	s.Bizdate = &v
	return s
}

func (s *CreateDagTestRequest) SetName(v string) *CreateDagTestRequest {
	s.Name = &v
	return s
}

func (s *CreateDagTestRequest) SetNodeId(v int64) *CreateDagTestRequest {
	s.NodeId = &v
	return s
}

func (s *CreateDagTestRequest) SetNodeParams(v string) *CreateDagTestRequest {
	s.NodeParams = &v
	return s
}

func (s *CreateDagTestRequest) SetProjectEnv(v string) *CreateDagTestRequest {
	s.ProjectEnv = &v
	return s
}

func (s *CreateDagTestRequest) Validate() error {
	return dara.Validate(s)
}
