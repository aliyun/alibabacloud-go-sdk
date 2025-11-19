// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCorePackageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *DescribeCorePackageListRequest
	GetApiKey() *string
	SetQueryDeductionInstances(v bool) *DescribeCorePackageListRequest
	GetQueryDeductionInstances() *bool
	SetQueryScenario(v string) *DescribeCorePackageListRequest
	GetQueryScenario() *string
}

type DescribeCorePackageListRequest struct {
	// This parameter is required.
	ApiKey                  *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	QueryDeductionInstances *bool   `json:"QueryDeductionInstances,omitempty" xml:"QueryDeductionInstances,omitempty"`
	QueryScenario           *string `json:"QueryScenario,omitempty" xml:"QueryScenario,omitempty"`
}

func (s DescribeCorePackageListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCorePackageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCorePackageListRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribeCorePackageListRequest) GetQueryDeductionInstances() *bool {
	return s.QueryDeductionInstances
}

func (s *DescribeCorePackageListRequest) GetQueryScenario() *string {
	return s.QueryScenario
}

func (s *DescribeCorePackageListRequest) SetApiKey(v string) *DescribeCorePackageListRequest {
	s.ApiKey = &v
	return s
}

func (s *DescribeCorePackageListRequest) SetQueryDeductionInstances(v bool) *DescribeCorePackageListRequest {
	s.QueryDeductionInstances = &v
	return s
}

func (s *DescribeCorePackageListRequest) SetQueryScenario(v string) *DescribeCorePackageListRequest {
	s.QueryScenario = &v
	return s
}

func (s *DescribeCorePackageListRequest) Validate() error {
	return dara.Validate(s)
}
