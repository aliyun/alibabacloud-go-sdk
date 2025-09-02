// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSuccessInstanceTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *GetSuccessInstanceTrendRequest
	GetProjectId() *int64
}

type GetSuccessInstanceTrendRequest struct {
	// The DataWorks workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9527
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetSuccessInstanceTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSuccessInstanceTrendRequest) GoString() string {
	return s.String()
}

func (s *GetSuccessInstanceTrendRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetSuccessInstanceTrendRequest) SetProjectId(v int64) *GetSuccessInstanceTrendRequest {
	s.ProjectId = &v
	return s
}

func (s *GetSuccessInstanceTrendRequest) Validate() error {
	return dara.Validate(s)
}
