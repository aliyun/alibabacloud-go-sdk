// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoutineCodeRevisionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DescribeRoutineCodeRevisionRequest
	GetName() *string
	SetSelectCodeRevision(v string) *DescribeRoutineCodeRevisionRequest
	GetSelectCodeRevision() *string
}

type DescribeRoutineCodeRevisionRequest struct {
	// The name of the routine. The name must be unique among the routines that belong to the same Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version of the JavaScript code that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1611151912787121550
	SelectCodeRevision *string `json:"SelectCodeRevision,omitempty" xml:"SelectCodeRevision,omitempty"`
}

func (s DescribeRoutineCodeRevisionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoutineCodeRevisionRequest) GoString() string {
	return s.String()
}

func (s *DescribeRoutineCodeRevisionRequest) GetName() *string {
	return s.Name
}

func (s *DescribeRoutineCodeRevisionRequest) GetSelectCodeRevision() *string {
	return s.SelectCodeRevision
}

func (s *DescribeRoutineCodeRevisionRequest) SetName(v string) *DescribeRoutineCodeRevisionRequest {
	s.Name = &v
	return s
}

func (s *DescribeRoutineCodeRevisionRequest) SetSelectCodeRevision(v string) *DescribeRoutineCodeRevisionRequest {
	s.SelectCodeRevision = &v
	return s
}

func (s *DescribeRoutineCodeRevisionRequest) Validate() error {
	return dara.Validate(s)
}
