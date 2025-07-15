// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainConfigs(v *DescribeLiveDomainConfigsResponseBodyDomainConfigs) *DescribeLiveDomainConfigsResponseBody
	GetDomainConfigs() *DescribeLiveDomainConfigsResponseBodyDomainConfigs
	SetRequestId(v string) *DescribeLiveDomainConfigsResponseBody
	GetRequestId() *string
}

type DescribeLiveDomainConfigsResponseBody struct {
	// The configurations of the domain name.
	DomainConfigs *DescribeLiveDomainConfigsResponseBodyDomainConfigs `json:"DomainConfigs,omitempty" xml:"DomainConfigs,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F8AA0364-0FDB-4AD5-AC74-D69FAB8924ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveDomainConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainConfigsResponseBody) GetDomainConfigs() *DescribeLiveDomainConfigsResponseBodyDomainConfigs {
	return s.DomainConfigs
}

func (s *DescribeLiveDomainConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainConfigsResponseBody) SetDomainConfigs(v *DescribeLiveDomainConfigsResponseBodyDomainConfigs) *DescribeLiveDomainConfigsResponseBody {
	s.DomainConfigs = v
	return s
}

func (s *DescribeLiveDomainConfigsResponseBody) SetRequestId(v string) *DescribeLiveDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveDomainConfigsResponseBodyDomainConfigs struct {
	DomainConfig []*DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfig `json:"DomainConfig,omitempty" xml:"DomainConfig,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainConfigsResponseBodyDomainConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainConfigsResponseBodyDomainConfigs) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainConfigsResponseBodyDomainConfigs) GetDomainConfig() []*DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfig {
	return s.DomainConfig
}

func (s *DescribeLiveDomainConfigsResponseBodyDomainConfigs) SetDomainConfig(v []*DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfig) *DescribeLiveDomainConfigsResponseBodyDomainConfigs {
	s.DomainConfig = v
	return s
}

func (s *DescribeLiveDomainConfigsResponseBodyDomainConfigs) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfig struct {
	// The configuration ID.
	//
	// example:
	//
	// 5003576
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The attributes of the feature.
	FunctionArgs *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty" type:"Struct"`
	// The name of the feature.
	//
	// example:
	//
	// set_req_host_header
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

func (s DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfig) GetConfigId() *string {
	return s.ConfigId
}

func (s *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfig) GetFunctionArgs() *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs {
	return s.FunctionArgs
}

func (s *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfig) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfig) GetStatus() *string {
	return s.Status
}

func (s *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfig) SetConfigId(v string) *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.ConfigId = &v
	return s
}

func (s *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfig) SetFunctionArgs(v *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.FunctionArgs = v
	return s
}

func (s *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfig) SetFunctionName(v string) *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.FunctionName = &v
	return s
}

func (s *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfig) SetStatus(v string) *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.Status = &v
	return s
}

func (s *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs struct {
	FunctionArg []*DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg `json:"FunctionArg,omitempty" xml:"FunctionArg,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) GetFunctionArg() []*DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	return s.FunctionArg
}

func (s *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) SetFunctionArg(v []*DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs {
	s.FunctionArg = v
	return s
}

func (s *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg struct {
	// The name of the attribute.
	//
	// example:
	//
	// domain_name
	ArgName *string `json:"ArgName,omitempty" xml:"ArgName,omitempty"`
	// The attribute value.
	//
	// example:
	//
	// developer.aliyundoc.com
	ArgValue *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty"`
}

func (s DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) GetArgName() *string {
	return s.ArgName
}

func (s *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) GetArgValue() *string {
	return s.ArgValue
}

func (s *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) SetArgName(v string) *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	s.ArgName = &v
	return s
}

func (s *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) SetArgValue(v string) *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	s.ArgValue = &v
	return s
}

func (s *DescribeLiveDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) Validate() error {
	return dara.Validate(s)
}
