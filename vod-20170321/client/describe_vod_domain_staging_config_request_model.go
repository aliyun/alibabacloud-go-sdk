// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainStagingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainStagingConfigRequest
	GetDomainName() *string
	SetFunctionNames(v string) *DescribeVodDomainStagingConfigRequest
	GetFunctionNames() *string
	SetOwnerId(v int64) *DescribeVodDomainStagingConfigRequest
	GetOwnerId() *int64
}

type DescribeVodDomainStagingConfigRequest struct {
	// This parameter is required.
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVodDomainStagingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainStagingConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainStagingConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainStagingConfigRequest) GetFunctionNames() *string {
	return s.FunctionNames
}

func (s *DescribeVodDomainStagingConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainStagingConfigRequest) SetDomainName(v string) *DescribeVodDomainStagingConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainStagingConfigRequest) SetFunctionNames(v string) *DescribeVodDomainStagingConfigRequest {
	s.FunctionNames = &v
	return s
}

func (s *DescribeVodDomainStagingConfigRequest) SetOwnerId(v int64) *DescribeVodDomainStagingConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainStagingConfigRequest) Validate() error {
	return dara.Validate(s)
}
