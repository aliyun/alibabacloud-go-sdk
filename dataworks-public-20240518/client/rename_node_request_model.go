// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *RenameNodeRequest
	GetId() *string
	SetName(v string) *RenameNodeRequest
	GetName() *string
	SetProjectId(v int64) *RenameNodeRequest
	GetProjectId() *int64
}

type RenameNodeRequest struct {
	// The unique identifier of the Data Studio node.
	//
	// >  This field is of the Long type in SDK versions prior to 8.0.0, and of the String type in SDK versions 8.0.0 and later. This change does not affect normal SDK usage; the parameter will still be returned according to the type defined in the SDK. However, compilation failures may occur due to the type change only when upgrading the SDK across version 8.0.0. In this case, you must manually update the data type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The unique identifier of the Data Studio node.
	//
	// >  This field is of the Long type in SDK versions prior to 8.0.0, and of the String type in SDK versions 8.0.0 and later. This change does not affect the normal use of the SDK. The parameter is returned based on the type defined in the SDK. Compilation failures caused by the type change may occur only when you upgrade the SDK across version 8.0.0. In this case, you must manually update the data type.
	//
	// This parameter is required.
	//
	// example:
	//
	// Rename
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s RenameNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s RenameNodeRequest) GoString() string {
	return s.String()
}

func (s *RenameNodeRequest) GetId() *string {
	return s.Id
}

func (s *RenameNodeRequest) GetName() *string {
	return s.Name
}

func (s *RenameNodeRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *RenameNodeRequest) SetId(v string) *RenameNodeRequest {
	s.Id = &v
	return s
}

func (s *RenameNodeRequest) SetName(v string) *RenameNodeRequest {
	s.Name = &v
	return s
}

func (s *RenameNodeRequest) SetProjectId(v int64) *RenameNodeRequest {
	s.ProjectId = &v
	return s
}

func (s *RenameNodeRequest) Validate() error {
	return dara.Validate(s)
}
