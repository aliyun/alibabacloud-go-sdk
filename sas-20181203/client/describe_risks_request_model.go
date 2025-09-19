// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRisksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeRisksRequest
	GetLang() *string
	SetLimit(v int32) *DescribeRisksRequest
	GetLimit() *int32
	SetResourceDirectoryAccountId(v int64) *DescribeRisksRequest
	GetResourceDirectoryAccountId() *int64
	SetRiskId(v int64) *DescribeRisksRequest
	GetRiskId() *int64
	SetRiskName(v string) *DescribeRisksRequest
	GetRiskName() *string
}

type DescribeRisksRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return. Default value: 20.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain the IDs.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The baseline ID.
	//
	// >  You can call the [DescribeCheckWarningSummary](~~DescribeCheckWarningSummary~~) operation to query the baseline IDs.
	//
	// example:
	//
	// 75
	RiskId *int64 `json:"RiskId,omitempty" xml:"RiskId,omitempty"`
	// The name of the baseline.
	//
	// example:
	//
	// docker
	RiskName *string `json:"RiskName,omitempty" xml:"RiskName,omitempty"`
}

func (s DescribeRisksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRisksRequest) GoString() string {
	return s.String()
}

func (s *DescribeRisksRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRisksRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *DescribeRisksRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeRisksRequest) GetRiskId() *int64 {
	return s.RiskId
}

func (s *DescribeRisksRequest) GetRiskName() *string {
	return s.RiskName
}

func (s *DescribeRisksRequest) SetLang(v string) *DescribeRisksRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRisksRequest) SetLimit(v int32) *DescribeRisksRequest {
	s.Limit = &v
	return s
}

func (s *DescribeRisksRequest) SetResourceDirectoryAccountId(v int64) *DescribeRisksRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeRisksRequest) SetRiskId(v int64) *DescribeRisksRequest {
	s.RiskId = &v
	return s
}

func (s *DescribeRisksRequest) SetRiskName(v string) *DescribeRisksRequest {
	s.RiskName = &v
	return s
}

func (s *DescribeRisksRequest) Validate() error {
	return dara.Validate(s)
}
