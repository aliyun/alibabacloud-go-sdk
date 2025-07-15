// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackLiveStagingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *RollbackLiveStagingConfigRequest
	GetDomainName() *string
	SetFunctionName(v string) *RollbackLiveStagingConfigRequest
	GetFunctionName() *string
	SetOwnerId(v int64) *RollbackLiveStagingConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RollbackLiveStagingConfigRequest
	GetRegionId() *string
}

type RollbackLiveStagingConfigRequest struct {
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

func (s RollbackLiveStagingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackLiveStagingConfigRequest) GoString() string {
	return s.String()
}

func (s *RollbackLiveStagingConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *RollbackLiveStagingConfigRequest) GetFunctionName() *string {
	return s.FunctionName
}

func (s *RollbackLiveStagingConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RollbackLiveStagingConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RollbackLiveStagingConfigRequest) SetDomainName(v string) *RollbackLiveStagingConfigRequest {
	s.DomainName = &v
	return s
}

func (s *RollbackLiveStagingConfigRequest) SetFunctionName(v string) *RollbackLiveStagingConfigRequest {
	s.FunctionName = &v
	return s
}

func (s *RollbackLiveStagingConfigRequest) SetOwnerId(v int64) *RollbackLiveStagingConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *RollbackLiveStagingConfigRequest) SetRegionId(v string) *RollbackLiveStagingConfigRequest {
	s.RegionId = &v
	return s
}

func (s *RollbackLiveStagingConfigRequest) Validate() error {
	return dara.Validate(s)
}
