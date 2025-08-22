// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainStagingConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainConfigs(v []*DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs) *DescribeDcdnDomainStagingConfigResponseBody
	GetDomainConfigs() []*DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs
	SetRequestId(v string) *DescribeDcdnDomainStagingConfigResponseBody
	GetRequestId() *string
}

type DescribeDcdnDomainStagingConfigResponseBody struct {
	// The configurations of accelerated domain names returned.
	DomainConfigs []*DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs `json:"DomainConfigs,omitempty" xml:"DomainConfigs,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// C80705BF-0F76-41FA-BAD1-5B59296A4E59
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainStagingConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainStagingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainStagingConfigResponseBody) GetDomainConfigs() []*DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs {
	return s.DomainConfigs
}

func (s *DescribeDcdnDomainStagingConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainStagingConfigResponseBody) SetDomainConfigs(v []*DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs) *DescribeDcdnDomainStagingConfigResponseBody {
	s.DomainConfigs = v
	return s
}

func (s *DescribeDcdnDomainStagingConfigResponseBody) SetRequestId(v string) *DescribeDcdnDomainStagingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainStagingConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs struct {
	// The ID of the configuration.
	//
	// example:
	//
	// 6295
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The following table describes the features.
	FunctionArgs []*DescribeDcdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty" type:"Repeated"`
	// The name of the feature.
	//
	// example:
	//
	// aliauth
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The status. Valid values:
	//
	// 	- success: The configuration is successful.
	//
	// 	- testing: The configuration is under testing.
	//
	// 	- failed: The task failed.
	//
	// 	- configuring: The feature is being configured.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs) GetConfigId() *string {
	return s.ConfigId
}

func (s *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs) GetFunctionArgs() []*DescribeDcdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs {
	return s.FunctionArgs
}

func (s *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs) GetStatus() *string {
	return s.Status
}

func (s *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs) SetConfigId(v string) *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs {
	s.ConfigId = &v
	return s
}

func (s *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs) SetFunctionArgs(v []*DescribeDcdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs {
	s.FunctionArgs = v
	return s
}

func (s *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs) SetFunctionName(v string) *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs {
	s.FunctionName = &v
	return s
}

func (s *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs) SetStatus(v string) *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs {
	s.Status = &v
	return s
}

func (s *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigs) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs struct {
	// The name of the configuration.
	//
	// example:
	//
	// auth_type
	ArgName *string `json:"ArgName,omitempty" xml:"ArgName,omitempty"`
	// The value of the configuration.
	//
	// example:
	//
	// req_auth
	ArgValue *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty"`
}

func (s DescribeDcdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) GetArgName() *string {
	return s.ArgName
}

func (s *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) GetArgValue() *string {
	return s.ArgValue
}

func (s *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) SetArgName(v string) *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs {
	s.ArgName = &v
	return s
}

func (s *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) SetArgValue(v string) *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs {
	s.ArgValue = &v
	return s
}

func (s *DescribeDcdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) Validate() error {
	return dara.Validate(s)
}
