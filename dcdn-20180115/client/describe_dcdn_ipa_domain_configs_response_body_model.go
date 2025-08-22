// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnIpaDomainConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainConfigs(v *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigs) *DescribeDcdnIpaDomainConfigsResponseBody
	GetDomainConfigs() *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigs
	SetRequestId(v string) *DescribeDcdnIpaDomainConfigsResponseBody
	GetRequestId() *string
}

type DescribeDcdnIpaDomainConfigsResponseBody struct {
	// The configurations of the domain name.
	DomainConfigs *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigs `json:"DomainConfigs,omitempty" xml:"DomainConfigs,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// F8AA0364-0FDB-4AD5-AC74-D69FAB8924ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnIpaDomainConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainConfigsResponseBody) GetDomainConfigs() *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigs {
	return s.DomainConfigs
}

func (s *DescribeDcdnIpaDomainConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnIpaDomainConfigsResponseBody) SetDomainConfigs(v *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigs) *DescribeDcdnIpaDomainConfigsResponseBody {
	s.DomainConfigs = v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsResponseBody) SetRequestId(v string) *DescribeDcdnIpaDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigs struct {
	DomainConfig []*DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig `json:"DomainConfig,omitempty" xml:"DomainConfig,omitempty" type:"Repeated"`
}

func (s DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigs) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigs) GetDomainConfig() []*DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig {
	return s.DomainConfig
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigs) SetDomainConfig(v []*DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig) *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigs {
	s.DomainConfig = v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigs) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig struct {
	// The ID of the configuration.
	//
	// example:
	//
	// 5003576
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The list of features.
	FunctionArgs *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty" type:"Struct"`
	// The name of the feature.
	//
	// example:
	//
	// protogw
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The configuration status of the feature. Valid values:
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

func (s DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig) GetConfigId() *string {
	return s.ConfigId
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig) GetFunctionArgs() *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs {
	return s.FunctionArgs
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig) GetStatus() *string {
	return s.Status
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig) SetConfigId(v string) *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.ConfigId = &v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig) SetFunctionArgs(v *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.FunctionArgs = v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig) SetFunctionName(v string) *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.FunctionName = &v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig) SetStatus(v string) *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.Status = &v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs struct {
	FunctionArg []*DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg `json:"FunctionArg,omitempty" xml:"FunctionArg,omitempty" type:"Repeated"`
}

func (s DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) GetFunctionArg() []*DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	return s.FunctionArg
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) SetFunctionArg(v []*DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs {
	s.FunctionArg = v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg struct {
	// The name of the configuration.
	//
	// example:
	//
	// file_type
	ArgName *string `json:"ArgName,omitempty" xml:"ArgName,omitempty"`
	// The value of the configuration.
	//
	// example:
	//
	// txt
	ArgValue *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty"`
}

func (s DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) GetArgName() *string {
	return s.ArgName
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) GetArgValue() *string {
	return s.ArgValue
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) SetArgName(v string) *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	s.ArgName = &v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) SetArgValue(v string) *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	s.ArgValue = &v
	return s
}

func (s *DescribeDcdnIpaDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) Validate() error {
	return dara.Validate(s)
}
