// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetNodeRequest
	GetId() *string
	SetProjectId(v int64) *GetNodeRequest
	GetProjectId() *int64
}

type GetNodeRequest struct {
	// The unique identifier of the Data Studio node.
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

func (s GetNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodeRequest) GoString() string {
	return s.String()
}

func (s *GetNodeRequest) GetId() *string {
	return s.Id
}

func (s *GetNodeRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetNodeRequest) SetId(v string) *GetNodeRequest {
	s.Id = &v
	return s
}

func (s *GetNodeRequest) SetProjectId(v int64) *GetNodeRequest {
	s.ProjectId = &v
	return s
}

func (s *GetNodeRequest) Validate() error {
	return dara.Validate(s)
}
