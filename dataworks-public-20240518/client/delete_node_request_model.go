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
	// The unique identifier of the data development node.
	//
	// 	Notice: This field was of the Long type in SDK versions earlier than 8.0.0 and is of the String type in SDK 8.0.0 and later. **This change does not affect normal SDK usage, and the parameter is still returned in the type defined in the SDK**. When upgrading across SDK version 8.0.0, the type change may cause project compilation failures, and you must manually correct the data type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the workspace settings page to obtain the workspace ID.
	//
	// This parameter specifies the DataWorks workspace for this API call.
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
