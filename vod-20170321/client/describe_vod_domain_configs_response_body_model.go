// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainConfigs(v *DescribeVodDomainConfigsResponseBodyDomainConfigs) *DescribeVodDomainConfigsResponseBody
	GetDomainConfigs() *DescribeVodDomainConfigsResponseBodyDomainConfigs
	SetRequestId(v string) *DescribeVodDomainConfigsResponseBody
	GetRequestId() *string
}

type DescribeVodDomainConfigsResponseBody struct {
	// The configurations of the domain name.
	DomainConfigs *DescribeVodDomainConfigsResponseBodyDomainConfigs `json:"DomainConfigs,omitempty" xml:"DomainConfigs,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F8AA0364-0FDB-4AD5-****-D69FAB8924ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodDomainConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainConfigsResponseBody) GetDomainConfigs() *DescribeVodDomainConfigsResponseBodyDomainConfigs {
	return s.DomainConfigs
}

func (s *DescribeVodDomainConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainConfigsResponseBody) SetDomainConfigs(v *DescribeVodDomainConfigsResponseBodyDomainConfigs) *DescribeVodDomainConfigsResponseBody {
	s.DomainConfigs = v
	return s
}

func (s *DescribeVodDomainConfigsResponseBody) SetRequestId(v string) *DescribeVodDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainConfigsResponseBodyDomainConfigs struct {
	DomainConfig []*DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig `json:"DomainConfig,omitempty" xml:"DomainConfig,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainConfigsResponseBodyDomainConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainConfigsResponseBodyDomainConfigs) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigs) GetDomainConfig() []*DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig {
	return s.DomainConfig
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigs) SetDomainConfig(v []*DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig) *DescribeVodDomainConfigsResponseBodyDomainConfigs {
	s.DomainConfig = v
	return s
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigs) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig struct {
	// The configuration ID.
	//
	// example:
	//
	// 5003576
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The feature parameters.
	FunctionArgs *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty" type:"Struct"`
	// The feature name.
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

func (s DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig) GetConfigId() *string {
	return s.ConfigId
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig) GetFunctionArgs() *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs {
	return s.FunctionArgs
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig) GetStatus() *string {
	return s.Status
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig) SetConfigId(v string) *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.ConfigId = &v
	return s
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig) SetFunctionArgs(v *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.FunctionArgs = v
	return s
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig) SetFunctionName(v string) *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.FunctionName = &v
	return s
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig) SetStatus(v string) *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.Status = &v
	return s
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs struct {
	FunctionArg []*DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg `json:"FunctionArg,omitempty" xml:"FunctionArg,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) GetFunctionArg() []*DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	return s.FunctionArg
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) SetFunctionArg(v []*DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs {
	s.FunctionArg = v
	return s
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg struct {
	// The parameter name.
	//
	// example:
	//
	// file_type
	ArgName *string `json:"ArgName,omitempty" xml:"ArgName,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// txt
	ArgValue *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty"`
}

func (s DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) GetArgName() *string {
	return s.ArgName
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) GetArgValue() *string {
	return s.ArgValue
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) SetArgName(v string) *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	s.ArgName = &v
	return s
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) SetArgValue(v string) *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	s.ArgValue = &v
	return s
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) Validate() error {
	return dara.Validate(s)
}
