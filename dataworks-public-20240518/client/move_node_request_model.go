// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *MoveNodeRequest
	GetId() *string
	SetPath(v string) *MoveNodeRequest
	GetPath() *string
	SetProjectId(v int64) *MoveNodeRequest
	GetProjectId() *int64
}

type MoveNodeRequest struct {
	// The ID of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The unique identifier of the Data Studio node.
	//
	// >  This field is of the Long type in SDK versions prior to 8.0.0, and of the String type in SDK versions 8.0.0 and later. This change does not affect normal SDK usage; the parameter will still be returned according to the type defined in the SDK. However, compilation failures may occur due to the type change only when upgrading the SDK across version 8.0.0. In this case, you must manually update the data type.
	//
	// This parameter is required.
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
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

func (s MoveNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveNodeRequest) GoString() string {
	return s.String()
}

func (s *MoveNodeRequest) GetId() *string {
	return s.Id
}

func (s *MoveNodeRequest) GetPath() *string {
	return s.Path
}

func (s *MoveNodeRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *MoveNodeRequest) SetId(v string) *MoveNodeRequest {
	s.Id = &v
	return s
}

func (s *MoveNodeRequest) SetPath(v string) *MoveNodeRequest {
	s.Path = &v
	return s
}

func (s *MoveNodeRequest) SetProjectId(v int64) *MoveNodeRequest {
	s.ProjectId = &v
	return s
}

func (s *MoveNodeRequest) Validate() error {
	return dara.Validate(s)
}
