// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainConfigs(v []*DescribeVsDomainConfigsResponseBodyDomainConfigs) *DescribeVsDomainConfigsResponseBody
	GetDomainConfigs() []*DescribeVsDomainConfigsResponseBodyDomainConfigs
	SetRequestId(v string) *DescribeVsDomainConfigsResponseBody
	GetRequestId() *string
}

type DescribeVsDomainConfigsResponseBody struct {
	DomainConfigs []*DescribeVsDomainConfigsResponseBodyDomainConfigs `json:"DomainConfigs,omitempty" xml:"DomainConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// D94D0E1E-E71B-562D-8C18-969BB3653FBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVsDomainConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainConfigsResponseBody) GetDomainConfigs() []*DescribeVsDomainConfigsResponseBodyDomainConfigs {
	return s.DomainConfigs
}

func (s *DescribeVsDomainConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVsDomainConfigsResponseBody) SetDomainConfigs(v []*DescribeVsDomainConfigsResponseBodyDomainConfigs) *DescribeVsDomainConfigsResponseBody {
	s.DomainConfigs = v
	return s
}

func (s *DescribeVsDomainConfigsResponseBody) SetRequestId(v string) *DescribeVsDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsDomainConfigsResponseBody) Validate() error {
	if s.DomainConfigs != nil {
		for _, item := range s.DomainConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVsDomainConfigsResponseBodyDomainConfigs struct {
	// example:
	//
	// 6295
	ConfigId     *string                                                         `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	FunctionArgs []*DescribeVsDomainConfigsResponseBodyDomainConfigsFunctionArgs `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty" type:"Repeated"`
	// example:
	//
	// aliauth
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeVsDomainConfigsResponseBodyDomainConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainConfigsResponseBodyDomainConfigs) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainConfigsResponseBodyDomainConfigs) GetConfigId() *string {
	return s.ConfigId
}

func (s *DescribeVsDomainConfigsResponseBodyDomainConfigs) GetFunctionArgs() []*DescribeVsDomainConfigsResponseBodyDomainConfigsFunctionArgs {
	return s.FunctionArgs
}

func (s *DescribeVsDomainConfigsResponseBodyDomainConfigs) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DescribeVsDomainConfigsResponseBodyDomainConfigs) GetStatus() *string {
	return s.Status
}

func (s *DescribeVsDomainConfigsResponseBodyDomainConfigs) SetConfigId(v string) *DescribeVsDomainConfigsResponseBodyDomainConfigs {
	s.ConfigId = &v
	return s
}

func (s *DescribeVsDomainConfigsResponseBodyDomainConfigs) SetFunctionArgs(v []*DescribeVsDomainConfigsResponseBodyDomainConfigsFunctionArgs) *DescribeVsDomainConfigsResponseBodyDomainConfigs {
	s.FunctionArgs = v
	return s
}

func (s *DescribeVsDomainConfigsResponseBodyDomainConfigs) SetFunctionName(v string) *DescribeVsDomainConfigsResponseBodyDomainConfigs {
	s.FunctionName = &v
	return s
}

func (s *DescribeVsDomainConfigsResponseBodyDomainConfigs) SetStatus(v string) *DescribeVsDomainConfigsResponseBodyDomainConfigs {
	s.Status = &v
	return s
}

func (s *DescribeVsDomainConfigsResponseBodyDomainConfigs) Validate() error {
	if s.FunctionArgs != nil {
		for _, item := range s.FunctionArgs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVsDomainConfigsResponseBodyDomainConfigsFunctionArgs struct {
	// example:
	//
	// auth_type
	ArgName *string `json:"ArgName,omitempty" xml:"ArgName,omitempty"`
	// example:
	//
	// req_auth
	ArgValue *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty"`
}

func (s DescribeVsDomainConfigsResponseBodyDomainConfigsFunctionArgs) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainConfigsResponseBodyDomainConfigsFunctionArgs) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainConfigsResponseBodyDomainConfigsFunctionArgs) GetArgName() *string {
	return s.ArgName
}

func (s *DescribeVsDomainConfigsResponseBodyDomainConfigsFunctionArgs) GetArgValue() *string {
	return s.ArgValue
}

func (s *DescribeVsDomainConfigsResponseBodyDomainConfigsFunctionArgs) SetArgName(v string) *DescribeVsDomainConfigsResponseBodyDomainConfigsFunctionArgs {
	s.ArgName = &v
	return s
}

func (s *DescribeVsDomainConfigsResponseBodyDomainConfigsFunctionArgs) SetArgValue(v string) *DescribeVsDomainConfigsResponseBodyDomainConfigsFunctionArgs {
	s.ArgValue = &v
	return s
}

func (s *DescribeVsDomainConfigsResponseBodyDomainConfigsFunctionArgs) Validate() error {
	return dara.Validate(s)
}
