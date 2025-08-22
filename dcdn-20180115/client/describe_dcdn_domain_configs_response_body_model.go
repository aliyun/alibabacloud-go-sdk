// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainConfigs(v *DescribeDcdnDomainConfigsResponseBodyDomainConfigs) *DescribeDcdnDomainConfigsResponseBody
	GetDomainConfigs() *DescribeDcdnDomainConfigsResponseBodyDomainConfigs
	SetRequestId(v string) *DescribeDcdnDomainConfigsResponseBody
	GetRequestId() *string
}

type DescribeDcdnDomainConfigsResponseBody struct {
	// The configurations of the domain name.
	DomainConfigs *DescribeDcdnDomainConfigsResponseBodyDomainConfigs `json:"DomainConfigs,omitempty" xml:"DomainConfigs,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// F8AA0364-0FDB-4AD5-AC74-D69FAB8924ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainConfigsResponseBody) GetDomainConfigs() *DescribeDcdnDomainConfigsResponseBodyDomainConfigs {
	return s.DomainConfigs
}

func (s *DescribeDcdnDomainConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainConfigsResponseBody) SetDomainConfigs(v *DescribeDcdnDomainConfigsResponseBodyDomainConfigs) *DescribeDcdnDomainConfigsResponseBody {
	s.DomainConfigs = v
	return s
}

func (s *DescribeDcdnDomainConfigsResponseBody) SetRequestId(v string) *DescribeDcdnDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainConfigsResponseBodyDomainConfigs struct {
	DomainConfig []*DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig `json:"DomainConfig,omitempty" xml:"DomainConfig,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainConfigsResponseBodyDomainConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainConfigsResponseBodyDomainConfigs) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigs) GetDomainConfig() []*DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	return s.DomainConfig
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigs) SetDomainConfig(v []*DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig) *DescribeDcdnDomainConfigsResponseBodyDomainConfigs {
	s.DomainConfig = v
	return s
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigs) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig struct {
	// The ID of the configuration.
	//
	// example:
	//
	// 5068995
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configurations of the features.
	FunctionArgs *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty" type:"Struct"`
	// The feature name.
	//
	// example:
	//
	// set_req_host_header
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The ID of the advanced condition configuration.
	//
	// example:
	//
	// 1234567
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The status of the configuration. Valid values:
	//
	// 	- **success**: successful
	//
	// 	- **testing**: testing
	//
	// 	- **failed**: The configuration failed.
	//
	// 	- **configuring**: The configuration is in progress.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig) GetConfigId() *string {
	return s.ConfigId
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig) GetFunctionArgs() *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs {
	return s.FunctionArgs
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig) GetStatus() *string {
	return s.Status
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetConfigId(v string) *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.ConfigId = &v
	return s
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetFunctionArgs(v *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.FunctionArgs = v
	return s
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetFunctionName(v string) *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.FunctionName = &v
	return s
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetParentId(v string) *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.ParentId = &v
	return s
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetStatus(v string) *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.Status = &v
	return s
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs struct {
	FunctionArg []*DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg `json:"FunctionArg,omitempty" xml:"FunctionArg,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) GetFunctionArg() []*DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	return s.FunctionArg
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) SetFunctionArg(v []*DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs {
	s.FunctionArg = v
	return s
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg struct {
	// The name of the configuration.
	//
	// example:
	//
	// ttl
	ArgName *string `json:"ArgName,omitempty" xml:"ArgName,omitempty"`
	// The value of the configuration.
	//
	// example:
	//
	// 13
	ArgValue *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty"`
}

func (s DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) GetArgName() *string {
	return s.ArgName
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) GetArgValue() *string {
	return s.ArgValue
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) SetArgName(v string) *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	s.ArgName = &v
	return s
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) SetArgValue(v string) *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	s.ArgValue = &v
	return s
}

func (s *DescribeDcdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) Validate() error {
	return dara.Validate(s)
}
