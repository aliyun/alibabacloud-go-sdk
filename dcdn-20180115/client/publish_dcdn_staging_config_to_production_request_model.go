// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishDcdnStagingConfigToProductionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *PublishDcdnStagingConfigToProductionRequest
	GetDomainName() *string
	SetFunctionName(v string) *PublishDcdnStagingConfigToProductionRequest
	GetFunctionName() *string
}

type PublishDcdnStagingConfigToProductionRequest struct {
	// The accelerated domain name. You can specify only one domain name in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The name of the feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// aliauth
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
}

func (s PublishDcdnStagingConfigToProductionRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishDcdnStagingConfigToProductionRequest) GoString() string {
	return s.String()
}

func (s *PublishDcdnStagingConfigToProductionRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *PublishDcdnStagingConfigToProductionRequest) GetFunctionName() *string {
	return s.FunctionName
}

func (s *PublishDcdnStagingConfigToProductionRequest) SetDomainName(v string) *PublishDcdnStagingConfigToProductionRequest {
	s.DomainName = &v
	return s
}

func (s *PublishDcdnStagingConfigToProductionRequest) SetFunctionName(v string) *PublishDcdnStagingConfigToProductionRequest {
	s.FunctionName = &v
	return s
}

func (s *PublishDcdnStagingConfigToProductionRequest) Validate() error {
	return dara.Validate(s)
}
