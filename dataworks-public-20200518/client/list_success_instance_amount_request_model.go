// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSuccessInstanceAmountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *ListSuccessInstanceAmountRequest
	GetProjectId() *int64
}

type ListSuccessInstanceAmountRequest struct {
	// The ID of the DataWorks workspace. You can log on to the DataWorks console and go to the workspace configuration page to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9527
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListSuccessInstanceAmountRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSuccessInstanceAmountRequest) GoString() string {
	return s.String()
}

func (s *ListSuccessInstanceAmountRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListSuccessInstanceAmountRequest) SetProjectId(v int64) *ListSuccessInstanceAmountRequest {
	s.ProjectId = &v
	return s
}

func (s *ListSuccessInstanceAmountRequest) Validate() error {
	return dara.Validate(s)
}
