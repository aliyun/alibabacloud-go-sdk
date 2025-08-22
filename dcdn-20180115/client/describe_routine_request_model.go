// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoutineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DescribeRoutineRequest
	GetName() *string
}

type DescribeRoutineRequest struct {
	// The name of the routine. The name must be unique among the routines that belong to the same Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeRoutineRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoutineRequest) GoString() string {
	return s.String()
}

func (s *DescribeRoutineRequest) GetName() *string {
	return s.Name
}

func (s *DescribeRoutineRequest) SetName(v string) *DescribeRoutineRequest {
	s.Name = &v
	return s
}

func (s *DescribeRoutineRequest) Validate() error {
	return dara.Validate(s)
}
