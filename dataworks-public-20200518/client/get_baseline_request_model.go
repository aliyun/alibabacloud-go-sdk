// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBaselineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineId(v int64) *GetBaselineRequest
	GetBaselineId() *int64
	SetProjectId(v int64) *GetBaselineRequest
	GetProjectId() *int64
}

type GetBaselineRequest struct {
	// The baseline ID. You can call the [ListBaselines](https://help.aliyun.com/document_detail/2261507.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The ID of the workspace to which the baseline belongs. You can call the [ListBaselines](https://help.aliyun.com/document_detail/2261507.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetBaselineRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineRequest) GoString() string {
	return s.String()
}

func (s *GetBaselineRequest) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *GetBaselineRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetBaselineRequest) SetBaselineId(v int64) *GetBaselineRequest {
	s.BaselineId = &v
	return s
}

func (s *GetBaselineRequest) SetProjectId(v int64) *GetBaselineRequest {
	s.ProjectId = &v
	return s
}

func (s *GetBaselineRequest) Validate() error {
	return dara.Validate(s)
}
