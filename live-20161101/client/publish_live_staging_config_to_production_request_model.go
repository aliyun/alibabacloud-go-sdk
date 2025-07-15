// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishLiveStagingConfigToProductionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *PublishLiveStagingConfigToProductionRequest
	GetDomainName() *string
	SetFunctionName(v string) *PublishLiveStagingConfigToProductionRequest
	GetFunctionName() *string
	SetOwnerId(v int64) *PublishLiveStagingConfigToProductionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *PublishLiveStagingConfigToProductionRequest
	GetRegionId() *string
}

type PublishLiveStagingConfigToProductionRequest struct {
	// The accelerated domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// developer.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The name of the feature. For more information about how to obtain the feature name, see [DescribeLiveDomainStagingConfig](https://help.aliyun.com/document_detail/297374.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// aliauth
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s PublishLiveStagingConfigToProductionRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishLiveStagingConfigToProductionRequest) GoString() string {
	return s.String()
}

func (s *PublishLiveStagingConfigToProductionRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *PublishLiveStagingConfigToProductionRequest) GetFunctionName() *string {
	return s.FunctionName
}

func (s *PublishLiveStagingConfigToProductionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PublishLiveStagingConfigToProductionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PublishLiveStagingConfigToProductionRequest) SetDomainName(v string) *PublishLiveStagingConfigToProductionRequest {
	s.DomainName = &v
	return s
}

func (s *PublishLiveStagingConfigToProductionRequest) SetFunctionName(v string) *PublishLiveStagingConfigToProductionRequest {
	s.FunctionName = &v
	return s
}

func (s *PublishLiveStagingConfigToProductionRequest) SetOwnerId(v int64) *PublishLiveStagingConfigToProductionRequest {
	s.OwnerId = &v
	return s
}

func (s *PublishLiveStagingConfigToProductionRequest) SetRegionId(v string) *PublishLiveStagingConfigToProductionRequest {
	s.RegionId = &v
	return s
}

func (s *PublishLiveStagingConfigToProductionRequest) Validate() error {
	return dara.Validate(s)
}
