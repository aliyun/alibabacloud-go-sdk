// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskCheckSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeRiskCheckSummaryRequest
	GetLang() *string
	SetResourceDirectoryAccountId(v string) *DescribeRiskCheckSummaryRequest
	GetResourceDirectoryAccountId() *string
	SetResourceOwnerId(v int64) *DescribeRiskCheckSummaryRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *DescribeRiskCheckSummaryRequest
	GetSourceIp() *string
}

type DescribeRiskCheckSummaryRequest struct {
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The Alibaba Cloud account ID of the member account in the resource directory for multi-account security management.
	//
	// example:
	//
	// 1232428423234****
	ResourceDirectoryAccountId *string `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	ResourceOwnerId            *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IP address of the access source that you want to query.
	//
	// example:
	//
	// 1.2.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeRiskCheckSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskCheckSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckSummaryRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRiskCheckSummaryRequest) GetResourceDirectoryAccountId() *string {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeRiskCheckSummaryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRiskCheckSummaryRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeRiskCheckSummaryRequest) SetLang(v string) *DescribeRiskCheckSummaryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRiskCheckSummaryRequest) SetResourceDirectoryAccountId(v string) *DescribeRiskCheckSummaryRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeRiskCheckSummaryRequest) SetResourceOwnerId(v int64) *DescribeRiskCheckSummaryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRiskCheckSummaryRequest) SetSourceIp(v string) *DescribeRiskCheckSummaryRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRiskCheckSummaryRequest) Validate() error {
	return dara.Validate(s)
}
