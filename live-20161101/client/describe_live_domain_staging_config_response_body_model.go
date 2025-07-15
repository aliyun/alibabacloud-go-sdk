// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainStagingConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainConfigs(v []*DescribeLiveDomainStagingConfigResponseBodyDomainConfigs) *DescribeLiveDomainStagingConfigResponseBody
	GetDomainConfigs() []*DescribeLiveDomainStagingConfigResponseBodyDomainConfigs
	SetRequestId(v string) *DescribeLiveDomainStagingConfigResponseBody
	GetRequestId() *string
}

type DescribeLiveDomainStagingConfigResponseBody struct {
	// The feature configurations of the accelerated domain name.
	DomainConfigs []*DescribeLiveDomainStagingConfigResponseBodyDomainConfigs `json:"DomainConfigs,omitempty" xml:"DomainConfigs,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C80705BF-0F76-41FA-BAD1-5B59296A4E59
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveDomainStagingConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainStagingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainStagingConfigResponseBody) GetDomainConfigs() []*DescribeLiveDomainStagingConfigResponseBodyDomainConfigs {
	return s.DomainConfigs
}

func (s *DescribeLiveDomainStagingConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainStagingConfigResponseBody) SetDomainConfigs(v []*DescribeLiveDomainStagingConfigResponseBodyDomainConfigs) *DescribeLiveDomainStagingConfigResponseBody {
	s.DomainConfigs = v
	return s
}

func (s *DescribeLiveDomainStagingConfigResponseBody) SetRequestId(v string) *DescribeLiveDomainStagingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainStagingConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveDomainStagingConfigResponseBodyDomainConfigs struct {
	// The configuration ID.
	//
	// example:
	//
	// 6295
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configurations of the feature.
	FunctionArgs []*DescribeLiveDomainStagingConfigResponseBodyDomainConfigsFunctionArgs `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty" type:"Repeated"`
	// The name of the feature.
	//
	// example:
	//
	// aliauth
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The configuration status. Valid values:
	//
	// 	- **success**
	//
	// 	- **testing**
	//
	// 	- **failed**
	//
	// 	- **configuring**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeLiveDomainStagingConfigResponseBodyDomainConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainStagingConfigResponseBodyDomainConfigs) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainStagingConfigResponseBodyDomainConfigs) GetConfigId() *string {
	return s.ConfigId
}

func (s *DescribeLiveDomainStagingConfigResponseBodyDomainConfigs) GetFunctionArgs() []*DescribeLiveDomainStagingConfigResponseBodyDomainConfigsFunctionArgs {
	return s.FunctionArgs
}

func (s *DescribeLiveDomainStagingConfigResponseBodyDomainConfigs) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DescribeLiveDomainStagingConfigResponseBodyDomainConfigs) GetStatus() *string {
	return s.Status
}

func (s *DescribeLiveDomainStagingConfigResponseBodyDomainConfigs) SetConfigId(v string) *DescribeLiveDomainStagingConfigResponseBodyDomainConfigs {
	s.ConfigId = &v
	return s
}

func (s *DescribeLiveDomainStagingConfigResponseBodyDomainConfigs) SetFunctionArgs(v []*DescribeLiveDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) *DescribeLiveDomainStagingConfigResponseBodyDomainConfigs {
	s.FunctionArgs = v
	return s
}

func (s *DescribeLiveDomainStagingConfigResponseBodyDomainConfigs) SetFunctionName(v string) *DescribeLiveDomainStagingConfigResponseBodyDomainConfigs {
	s.FunctionName = &v
	return s
}

func (s *DescribeLiveDomainStagingConfigResponseBodyDomainConfigs) SetStatus(v string) *DescribeLiveDomainStagingConfigResponseBodyDomainConfigs {
	s.Status = &v
	return s
}

func (s *DescribeLiveDomainStagingConfigResponseBodyDomainConfigs) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveDomainStagingConfigResponseBodyDomainConfigsFunctionArgs struct {
	// The name of the parameter.
	//
	// example:
	//
	// auth_type
	ArgName *string `json:"ArgName,omitempty" xml:"ArgName,omitempty"`
	// The configured value.
	//
	// example:
	//
	// req_auth
	ArgValue *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty"`
}

func (s DescribeLiveDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) GetArgName() *string {
	return s.ArgName
}

func (s *DescribeLiveDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) GetArgValue() *string {
	return s.ArgValue
}

func (s *DescribeLiveDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) SetArgName(v string) *DescribeLiveDomainStagingConfigResponseBodyDomainConfigsFunctionArgs {
	s.ArgName = &v
	return s
}

func (s *DescribeLiveDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) SetArgValue(v string) *DescribeLiveDomainStagingConfigResponseBodyDomainConfigsFunctionArgs {
	s.ArgValue = &v
	return s
}

func (s *DescribeLiveDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) Validate() error {
	return dara.Validate(s)
}
