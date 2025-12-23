// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DescribeProjectResponseBody
	GetDescription() *string
	SetName(v string) *DescribeProjectResponseBody
	GetName() *string
	SetProjectId(v string) *DescribeProjectResponseBody
	GetProjectId() *string
	SetRequestId(v string) *DescribeProjectResponseBody
	GetRequestId() *string
}

type DescribeProjectResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// this is a project description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_project_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// p-3q9jo749ne5****
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// B897B94B-6754-5D09-AB8C-2E8186CCADC0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DescribeProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProjectResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeProjectResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeProjectResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProjectResponseBody) SetDescription(v string) *DescribeProjectResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeProjectResponseBody) SetName(v string) *DescribeProjectResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeProjectResponseBody) SetProjectId(v string) *DescribeProjectResponseBody {
	s.ProjectId = &v
	return s
}

func (s *DescribeProjectResponseBody) SetRequestId(v string) *DescribeProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
