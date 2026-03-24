// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDomainConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainConfigs(v *DescribeCdnDomainConfigsResponseBodyDomainConfigs) *DescribeCdnDomainConfigsResponseBody
	GetDomainConfigs() *DescribeCdnDomainConfigsResponseBodyDomainConfigs
	SetRequestId(v string) *DescribeCdnDomainConfigsResponseBody
	GetRequestId() *string
}

type DescribeCdnDomainConfigsResponseBody struct {
	DomainConfigs *DescribeCdnDomainConfigsResponseBodyDomainConfigs `json:"DomainConfigs,omitempty" xml:"DomainConfigs,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// C80705BF-0F76-41FA-BAD1-5B59296A4E59
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnDomainConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainConfigsResponseBody) GetDomainConfigs() *DescribeCdnDomainConfigsResponseBodyDomainConfigs {
	return s.DomainConfigs
}

func (s *DescribeCdnDomainConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnDomainConfigsResponseBody) SetDomainConfigs(v *DescribeCdnDomainConfigsResponseBodyDomainConfigs) *DescribeCdnDomainConfigsResponseBody {
	s.DomainConfigs = v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBody) SetRequestId(v string) *DescribeCdnDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBody) Validate() error {
	if s.DomainConfigs != nil {
		if err := s.DomainConfigs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnDomainConfigsResponseBodyDomainConfigs struct {
	DomainConfig []*DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig `json:"DomainConfig,omitempty" xml:"DomainConfig,omitempty" type:"Repeated"`
}

func (s DescribeCdnDomainConfigsResponseBodyDomainConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainConfigsResponseBodyDomainConfigs) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigs) GetDomainConfig() []*DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	return s.DomainConfig
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigs) SetDomainConfig(v []*DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) *DescribeCdnDomainConfigsResponseBodyDomainConfigs {
	s.DomainConfig = v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigs) Validate() error {
	if s.DomainConfig != nil {
		for _, item := range s.DomainConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig struct {
	ConfigId     *string                                                                    `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	FunctionArgs *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty" type:"Struct"`
	FunctionName *string                                                                    `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	ParentId     *string                                                                    `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Status       *string                                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) GetConfigId() *string {
	return s.ConfigId
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) GetFunctionArgs() *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs {
	return s.FunctionArgs
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) GetStatus() *string {
	return s.Status
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetConfigId(v string) *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.ConfigId = &v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetFunctionArgs(v *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.FunctionArgs = v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetFunctionName(v string) *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.FunctionName = &v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetParentId(v string) *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.ParentId = &v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetStatus(v string) *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.Status = &v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) Validate() error {
	if s.FunctionArgs != nil {
		if err := s.FunctionArgs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs struct {
	FunctionArg []*DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg `json:"FunctionArg,omitempty" xml:"FunctionArg,omitempty" type:"Repeated"`
}

func (s DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) GetFunctionArg() []*DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	return s.FunctionArg
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) SetFunctionArg(v []*DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs {
	s.FunctionArg = v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) Validate() error {
	if s.FunctionArg != nil {
		for _, item := range s.FunctionArg {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg struct {
	ArgName  *string `json:"ArgName,omitempty" xml:"ArgName,omitempty"`
	ArgValue *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty"`
}

func (s DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) GetArgName() *string {
	return s.ArgName
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) GetArgValue() *string {
	return s.ArgValue
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) SetArgName(v string) *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	s.ArgName = &v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) SetArgValue(v string) *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	s.ArgValue = &v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) Validate() error {
	return dara.Validate(s)
}
