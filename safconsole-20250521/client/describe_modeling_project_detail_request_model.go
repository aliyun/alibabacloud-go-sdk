// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelingProjectDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *DescribeModelingProjectDetailRequest
	GetProjectId() *int64
}

type DescribeModelingProjectDetailRequest struct {
	// Project ID.
	//
	// example:
	//
	// 01
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeModelingProjectDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelingProjectDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeModelingProjectDetailRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DescribeModelingProjectDetailRequest) SetProjectId(v int64) *DescribeModelingProjectDetailRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeModelingProjectDetailRequest) Validate() error {
	return dara.Validate(s)
}
