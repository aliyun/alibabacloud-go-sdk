// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoutineRelatedDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DescribeRoutineRelatedDomainsRequest
	GetName() *string
}

type DescribeRoutineRelatedDomainsRequest struct {
	// The name of the routine. The name is unique in the same account.
	//
	// This parameter is required.
	//
	// example:
	//
	// routine_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeRoutineRelatedDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoutineRelatedDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRoutineRelatedDomainsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeRoutineRelatedDomainsRequest) SetName(v string) *DescribeRoutineRelatedDomainsRequest {
	s.Name = &v
	return s
}

func (s *DescribeRoutineRelatedDomainsRequest) Validate() error {
	return dara.Validate(s)
}
