// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishVodStagingConfigToProductionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *PublishVodStagingConfigToProductionRequest
	GetDomainName() *string
	SetOwnerId(v int64) *PublishVodStagingConfigToProductionRequest
	GetOwnerId() *int64
}

type PublishVodStagingConfigToProductionRequest struct {
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s PublishVodStagingConfigToProductionRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishVodStagingConfigToProductionRequest) GoString() string {
	return s.String()
}

func (s *PublishVodStagingConfigToProductionRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *PublishVodStagingConfigToProductionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PublishVodStagingConfigToProductionRequest) SetDomainName(v string) *PublishVodStagingConfigToProductionRequest {
	s.DomainName = &v
	return s
}

func (s *PublishVodStagingConfigToProductionRequest) SetOwnerId(v int64) *PublishVodStagingConfigToProductionRequest {
	s.OwnerId = &v
	return s
}

func (s *PublishVodStagingConfigToProductionRequest) Validate() error {
	return dara.Validate(s)
}
