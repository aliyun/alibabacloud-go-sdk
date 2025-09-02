// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBaselineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineId(v int64) *DeleteBaselineRequest
	GetBaselineId() *int64
	SetProjectId(v int64) *DeleteBaselineRequest
	GetProjectId() *int64
}

type DeleteBaselineRequest struct {
	// The baseline ID. You can call the [ListBaselines](https://help.aliyun.com/document_detail/2261507.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The ID of the workspace to which the baseline belongs. You can call the ListBaselines operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteBaselineRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBaselineRequest) GoString() string {
	return s.String()
}

func (s *DeleteBaselineRequest) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *DeleteBaselineRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteBaselineRequest) SetBaselineId(v int64) *DeleteBaselineRequest {
	s.BaselineId = &v
	return s
}

func (s *DeleteBaselineRequest) SetProjectId(v int64) *DeleteBaselineRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteBaselineRequest) Validate() error {
	return dara.Validate(s)
}
