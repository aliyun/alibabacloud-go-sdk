// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDcdnDomainStagingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SetDcdnDomainStagingConfigRequest
	GetDomainName() *string
	SetFunctions(v string) *SetDcdnDomainStagingConfigRequest
	GetFunctions() *string
}

type SetDcdnDomainStagingConfigRequest struct {
	// The accelerated domain name. Separate multiple accelerated domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The list of features. Format: `[{"functionArgs":[{"argName":"parameter key","argValue":"parameter value"},{"argName":"xx","argValue":"xx"}],"functionName": feature name"}]`
	//
	// > Separate multiple parameters with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// [{\\"functionArgs\\":[{\\"argName\\":\\"enable\\",\\"argValue\\":\\"on\\",\\"argName\\":\\"pri\\",\\"argValue\\":\\"1\\",\\"argName\\":\\"rule\\",\\"argValue\\":\\"xxx\\"}],\\"functionName\\":\\"edge_function\\"}]
	Functions *string `json:"Functions,omitempty" xml:"Functions,omitempty"`
}

func (s SetDcdnDomainStagingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDcdnDomainStagingConfigRequest) GoString() string {
	return s.String()
}

func (s *SetDcdnDomainStagingConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetDcdnDomainStagingConfigRequest) GetFunctions() *string {
	return s.Functions
}

func (s *SetDcdnDomainStagingConfigRequest) SetDomainName(v string) *SetDcdnDomainStagingConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetDcdnDomainStagingConfigRequest) SetFunctions(v string) *SetDcdnDomainStagingConfigRequest {
	s.Functions = &v
	return s
}

func (s *SetDcdnDomainStagingConfigRequest) Validate() error {
	return dara.Validate(s)
}
