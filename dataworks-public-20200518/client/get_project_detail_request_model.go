// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *GetProjectDetailRequest
	GetProjectId() *int64
}

type GetProjectDetailRequest struct {
	// The DataWorks workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 27
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetProjectDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProjectDetailRequest) GoString() string {
	return s.String()
}

func (s *GetProjectDetailRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetProjectDetailRequest) SetProjectId(v int64) *GetProjectDetailRequest {
	s.ProjectId = &v
	return s
}

func (s *GetProjectDetailRequest) Validate() error {
	return dara.Validate(s)
}
