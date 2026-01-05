// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteNodeRequest
	GetId() *string
	SetProjectId(v int64) *DeleteNodeRequest
	GetProjectId() *int64
}

type DeleteNodeRequest struct {
	// The unique identifier of the Data Studio node.
	//
	// >  This field is of the Long type in SDK versions prior to 8.0.0, and of the String type in SDK versions 8.0.0 and later. This change does not affect normal SDK usage; the parameter will still be returned according to the type defined in the SDK. However, compilation failures may occur due to the type change only when upgrading the SDK across version 8.0.0. In this case, you must manually update the data type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s *DeleteNodeRequest) GetId() *string {
	return s.Id
}

func (s *DeleteNodeRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteNodeRequest) SetId(v string) *DeleteNodeRequest {
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
