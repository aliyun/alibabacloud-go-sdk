// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveDomainStagingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SetLiveDomainStagingConfigRequest
	GetDomainName() *string
	SetFunctions(v string) *SetLiveDomainStagingConfigRequest
	GetFunctions() *string
	SetOwnerId(v int64) *SetLiveDomainStagingConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetLiveDomainStagingConfigRequest
	GetRegionId() *string
}

type SetLiveDomainStagingConfigRequest struct {
	// The accelerated domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// developer.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The list of features. You must specify the ConfigId field when you want to modify the configurations. For more information, see **Features specified by the Functions parameter**.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"functionArgs":[{"argName":"enable","argValue":"on"},{"argName":"pri","argValue":"1"},{"argName":"rule","argValue":"xxx"}],"functionName":"edge_function"}]
	Functions *string `json:"Functions,omitempty" xml:"Functions,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetLiveDomainStagingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLiveDomainStagingConfigRequest) GoString() string {
	return s.String()
}

func (s *SetLiveDomainStagingConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetLiveDomainStagingConfigRequest) GetFunctions() *string {
	return s.Functions
}

func (s *SetLiveDomainStagingConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetLiveDomainStagingConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetLiveDomainStagingConfigRequest) SetDomainName(v string) *SetLiveDomainStagingConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetLiveDomainStagingConfigRequest) SetFunctions(v string) *SetLiveDomainStagingConfigRequest {
	s.Functions = &v
	return s
}

func (s *SetLiveDomainStagingConfigRequest) SetOwnerId(v int64) *SetLiveDomainStagingConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLiveDomainStagingConfigRequest) SetRegionId(v string) *SetLiveDomainStagingConfigRequest {
	s.RegionId = &v
	return s
}

func (s *SetLiveDomainStagingConfigRequest) Validate() error {
	return dara.Validate(s)
}
