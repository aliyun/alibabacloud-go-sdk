// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *DescribeElasticPlanRequest
	GetDryRun() *bool
}

type DescribeElasticPlanRequest struct {
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s DescribeElasticPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticPlanRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeElasticPlanRequest) SetDryRun(v bool) *DescribeElasticPlanRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeElasticPlanRequest) Validate() error {
	return dara.Validate(s)
}
