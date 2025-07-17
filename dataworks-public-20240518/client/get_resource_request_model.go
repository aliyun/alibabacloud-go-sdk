// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetResourceRequest
	GetId() *int64
	SetProjectId(v int64) *GetResourceRequest
	GetProjectId() *int64
}

type GetResourceRequest struct {
	// The ID of the file resource.
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

func (s GetResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceRequest) GoString() string {
	return s.String()
}

func (s *GetResourceRequest) GetId() *int64 {
	return s.Id
}

func (s *GetResourceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetResourceRequest) SetId(v int64) *GetResourceRequest {
	s.Id = &v
	return s
}

func (s *GetResourceRequest) SetProjectId(v int64) *GetResourceRequest {
	s.ProjectId = &v
	return s
}

func (s *GetResourceRequest) Validate() error {
	return dara.Validate(s)
}
