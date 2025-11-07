// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNisInspectionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInspectionTaskId(v string) *DescribeNisInspectionTaskRequest
	GetInspectionTaskId() *string
}

type DescribeNisInspectionTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ni-8svmpe0yso2bhzr7fh79
	InspectionTaskId *string `json:"InspectionTaskId,omitempty" xml:"InspectionTaskId,omitempty"`
}

func (s DescribeNisInspectionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisInspectionTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeNisInspectionTaskRequest) GetInspectionTaskId() *string {
	return s.InspectionTaskId
}

func (s *DescribeNisInspectionTaskRequest) SetInspectionTaskId(v string) *DescribeNisInspectionTaskRequest {
	s.InspectionTaskId = &v
	return s
}

func (s *DescribeNisInspectionTaskRequest) Validate() error {
	return dara.Validate(s)
}
