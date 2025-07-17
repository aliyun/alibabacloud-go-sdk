// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetNodeRequest
	GetId() *int64
	SetProjectId(v int64) *GetNodeRequest
	GetProjectId() *int64
}

type GetNodeRequest struct {
	// The ID of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodeRequest) GoString() string {
	return s.String()
}

func (s *GetNodeRequest) GetId() *int64 {
	return s.Id
}

func (s *GetNodeRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetNodeRequest) SetId(v int64) *GetNodeRequest {
	s.Id = &v
	return s
}

func (s *GetNodeRequest) SetProjectId(v int64) *GetNodeRequest {
	s.ProjectId = &v
	return s
}

func (s *GetNodeRequest) Validate() error {
	return dara.Validate(s)
}
