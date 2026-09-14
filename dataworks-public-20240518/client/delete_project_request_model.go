// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteProjectRequest
	GetId() *int64
}

type DeleteProjectRequest struct {
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://dataworks.console.aliyun.com/workspace/list) and go to the workspace management page to obtain the ID.
	//
	// This parameter specifies the DataWorks workspace for this API call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteProjectRequest) SetId(v int64) *DeleteProjectRequest {
	s.Id = &v
	return s
}

func (s *DeleteProjectRequest) Validate() error {
	return dara.Validate(s)
}
