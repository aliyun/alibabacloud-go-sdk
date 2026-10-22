// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *MoveResourceRequest
	GetId() *string
	SetPath(v string) *MoveResourceRequest
	GetPath() *string
	SetProjectId(v int64) *MoveResourceRequest
	GetProjectId() *int64
}

type MoveResourceRequest struct {
	// The unique identifier of the Data Studio file resource.
	//
	// 	Notice: Before SDK version 8.0.0, this field is of the Long type. In SDK version 8.0.0 and later, this field is of the String type. **This change does not affect the normal use of the SDK. The parameter is still returned based on the type defined in the SDK**. The type change may cause project compilation failures only when you upgrade the SDK across version 8.0.0. In this case, you must manually correct the data type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The destination path for the move operation. Do not include the file resource name in the path.
	//
	// For example, to change the path of the math.py file to root/demo/math.py, set this parameter to root/demo.
	//
	// This parameter is required.
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace Configuration page to obtain the workspace ID.
	//
	// This parameter specifies the DataWorks workspace for this API operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s MoveResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveResourceRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceRequest) GetId() *string {
	return s.Id
}

func (s *MoveResourceRequest) GetPath() *string {
	return s.Path
}

func (s *MoveResourceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *MoveResourceRequest) SetId(v string) *MoveResourceRequest {
	s.Id = &v
	return s
}

func (s *MoveResourceRequest) SetPath(v string) *MoveResourceRequest {
	s.Path = &v
	return s
}

func (s *MoveResourceRequest) SetProjectId(v int64) *MoveResourceRequest {
	s.ProjectId = &v
	return s
}

func (s *MoveResourceRequest) Validate() error {
	return dara.Validate(s)
}
