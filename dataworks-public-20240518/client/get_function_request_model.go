// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetFunctionRequest
	GetId() *int64
	SetProjectId(v int64) *GetFunctionRequest
	GetProjectId() *int64
}

type GetFunctionRequest struct {
	// The ID of the UDF.
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

func (s GetFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionRequest) GoString() string {
	return s.String()
}

func (s *GetFunctionRequest) GetId() *int64 {
	return s.Id
}

func (s *GetFunctionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetFunctionRequest) SetId(v int64) *GetFunctionRequest {
	s.Id = &v
	return s
}

func (s *GetFunctionRequest) SetProjectId(v int64) *GetFunctionRequest {
	s.ProjectId = &v
	return s
}

func (s *GetFunctionRequest) Validate() error {
	return dara.Validate(s)
}
