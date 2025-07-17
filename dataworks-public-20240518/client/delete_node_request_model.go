// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteNodeRequest
	GetId() *int64
	SetProjectId(v int64) *DeleteNodeRequest
	GetProjectId() *int64
}

type DeleteNodeRequest struct {
	// The ID of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the ID.
	//
	// You can use this parameter to specify the DataWorks workspace on which you want to perform the API operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteNodeRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteNodeRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteNodeRequest) SetId(v int64) *DeleteNodeRequest {
	s.Id = &v
	return s
}

func (s *DeleteNodeRequest) SetProjectId(v int64) *DeleteNodeRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteNodeRequest) Validate() error {
	return dara.Validate(s)
}
