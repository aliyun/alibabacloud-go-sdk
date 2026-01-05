// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *MoveFunctionRequest
	GetId() *string
	SetPath(v string) *MoveFunctionRequest
	GetPath() *string
	SetProjectId(v int64) *MoveFunctionRequest
	GetProjectId() *int64
}

type MoveFunctionRequest struct {
	// The unique identifier of the UDF.
	//
	// >  This field is of the Long type in SDK versions prior to 8.0.0, and of the String type in SDK versions 8.0.0 and later. This change does not affect normal SDK usage; the parameter will still be returned according to the type defined in the SDK.. However, compilation failures may occur due to the type change only when upgrading the SDK across version 8.0.0. In this case, you must manually update the data type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 543217824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The path to which you want to move the UDF. You do not need to specify a UDF name in the path.
	//
	// For example, if you want to move the test UDF to root/demo/test, you must set this parameter to root/demo.
	//
	// This parameter is required.
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the ID.
	//
	// This parameter indicates the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s MoveFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveFunctionRequest) GoString() string {
	return s.String()
}

func (s *MoveFunctionRequest) GetId() *string {
	return s.Id
}

func (s *MoveFunctionRequest) GetPath() *string {
	return s.Path
}

func (s *MoveFunctionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *MoveFunctionRequest) SetId(v string) *MoveFunctionRequest {
	s.Id = &v
	return s
}

func (s *MoveFunctionRequest) SetPath(v string) *MoveFunctionRequest {
	s.Path = &v
	return s
}

func (s *MoveFunctionRequest) SetProjectId(v int64) *MoveFunctionRequest {
	s.ProjectId = &v
	return s
}

func (s *MoveFunctionRequest) Validate() error {
	return dara.Validate(s)
}
