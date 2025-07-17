// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteResourceRequest
	GetId() *int64
	SetProjectId(v int64) *DeleteResourceRequest
	GetProjectId() *int64
}

type DeleteResourceRequest struct {
	// The ID of the file resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the ID. You can use this parameter to specify the DataWorks workspace on which you want to perform the API operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteResourceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteResourceRequest) SetId(v int64) *DeleteResourceRequest {
	s.Id = &v
	return s
}

func (s *DeleteResourceRequest) SetProjectId(v int64) *DeleteResourceRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteResourceRequest) Validate() error {
	return dara.Validate(s)
}
