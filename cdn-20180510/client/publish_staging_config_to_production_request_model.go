// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishStagingConfigToProductionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *PublishStagingConfigToProductionRequest
	GetDomainName() *string
}

type PublishStagingConfigToProductionRequest struct {
	// The accelerated domain name. You can specify only one domain name in each call.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s PublishStagingConfigToProductionRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishStagingConfigToProductionRequest) GoString() string {
	return s.String()
}

func (s *PublishStagingConfigToProductionRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *PublishStagingConfigToProductionRequest) SetDomainName(v string) *PublishStagingConfigToProductionRequest {
	s.DomainName = &v
	return s
}

func (s *PublishStagingConfigToProductionRequest) Validate() error {
	return dara.Validate(s)
}
