// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetVodDomainStagingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SetVodDomainStagingConfigRequest
	GetDomainName() *string
	SetFunctions(v string) *SetVodDomainStagingConfigRequest
	GetFunctions() *string
	SetOwnerId(v int64) *SetVodDomainStagingConfigRequest
	GetOwnerId() *int64
}

type SetVodDomainStagingConfigRequest struct {
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	Functions *string `json:"Functions,omitempty" xml:"Functions,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s SetVodDomainStagingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetVodDomainStagingConfigRequest) GoString() string {
	return s.String()
}

func (s *SetVodDomainStagingConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetVodDomainStagingConfigRequest) GetFunctions() *string {
	return s.Functions
}

func (s *SetVodDomainStagingConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetVodDomainStagingConfigRequest) SetDomainName(v string) *SetVodDomainStagingConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetVodDomainStagingConfigRequest) SetFunctions(v string) *SetVodDomainStagingConfigRequest {
	s.Functions = &v
	return s
}

func (s *SetVodDomainStagingConfigRequest) SetOwnerId(v int64) *SetVodDomainStagingConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetVodDomainStagingConfigRequest) Validate() error {
	return dara.Validate(s)
}
