// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetFunctionRequest
	GetId() *string
	SetProjectId(v int64) *GetFunctionRequest
	GetProjectId() *int64
}

type GetFunctionRequest struct {
	// The unique identifier of the UDF function.
	//
	// >  This field is of type Long in SDK versions prior to 8.0.0, and of type String in SDK version 8.0.0 and later. This change does not affect the normal use of the SDK; parameters are still returned according to the type defined in the SDK. Compilation failures due to the type change may occur only when upgrading the SDK across version 8.0.0, in which case users need to manually correct the data type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s *GetFunctionRequest) GetId() *string {
	return s.Id
}

func (s *GetFunctionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetFunctionRequest) SetId(v string) *GetFunctionRequest {
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
