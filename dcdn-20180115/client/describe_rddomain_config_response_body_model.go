// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRDDomainConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainConfigs(v []*DescribeRDDomainConfigResponseBodyDomainConfigs) *DescribeRDDomainConfigResponseBody
	GetDomainConfigs() []*DescribeRDDomainConfigResponseBodyDomainConfigs
	SetRequestId(v string) *DescribeRDDomainConfigResponseBody
	GetRequestId() *string
}

type DescribeRDDomainConfigResponseBody struct {
	// The configuration of the domain name.
	DomainConfigs []*DescribeRDDomainConfigResponseBodyDomainConfigs `json:"DomainConfigs,omitempty" xml:"DomainConfigs,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// C80705BF-0F76-41FA-BAD1-5B59296A4E59
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRDDomainConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRDDomainConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRDDomainConfigResponseBody) GetDomainConfigs() []*DescribeRDDomainConfigResponseBodyDomainConfigs {
	return s.DomainConfigs
}

func (s *DescribeRDDomainConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRDDomainConfigResponseBody) SetDomainConfigs(v []*DescribeRDDomainConfigResponseBodyDomainConfigs) *DescribeRDDomainConfigResponseBody {
	s.DomainConfigs = v
	return s
}

func (s *DescribeRDDomainConfigResponseBody) SetRequestId(v string) *DescribeRDDomainConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRDDomainConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRDDomainConfigResponseBodyDomainConfigs struct {
	// The ID of the configuration.
	//
	// example:
	//
	// 6295
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configurations of the features.
	FunctionArgs []*DescribeRDDomainConfigResponseBodyDomainConfigsFunctionArgs `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty" type:"Repeated"`
	// The name of the feature.
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
	// The status. Valid values:
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

func (s DescribeRDDomainConfigResponseBodyDomainConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeRDDomainConfigResponseBodyDomainConfigs) GoString() string {
	return s.String()
}

func (s *DescribeRDDomainConfigResponseBodyDomainConfigs) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *DescribeRDDomainConfigResponseBodyDomainConfigs) GetFunctionArgs() []*DescribeRDDomainConfigResponseBodyDomainConfigsFunctionArgs {
	return s.FunctionArgs
}

func (s *DescribeRDDomainConfigResponseBodyDomainConfigs) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DescribeRDDomainConfigResponseBodyDomainConfigs) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeRDDomainConfigResponseBodyDomainConfigs) GetStatus() *string {
	return s.Status
}

func (s *DescribeRDDomainConfigResponseBodyDomainConfigs) SetConfigId(v int64) *DescribeRDDomainConfigResponseBodyDomainConfigs {
	s.ConfigId = &v
	return s
}

func (s *DescribeRDDomainConfigResponseBodyDomainConfigs) SetFunctionArgs(v []*DescribeRDDomainConfigResponseBodyDomainConfigsFunctionArgs) *DescribeRDDomainConfigResponseBodyDomainConfigs {
	s.FunctionArgs = v
	return s
}

func (s *DescribeRDDomainConfigResponseBodyDomainConfigs) SetFunctionName(v string) *DescribeRDDomainConfigResponseBodyDomainConfigs {
	s.FunctionName = &v
	return s
}

func (s *DescribeRDDomainConfigResponseBodyDomainConfigs) SetParentId(v string) *DescribeRDDomainConfigResponseBodyDomainConfigs {
	s.ParentId = &v
	return s
}

func (s *DescribeRDDomainConfigResponseBodyDomainConfigs) SetStatus(v string) *DescribeRDDomainConfigResponseBodyDomainConfigs {
	s.Status = &v
	return s
}

func (s *DescribeRDDomainConfigResponseBodyDomainConfigs) Validate() error {
	return dara.Validate(s)
}

type DescribeRDDomainConfigResponseBodyDomainConfigsFunctionArgs struct {
	// The name of the configuration.
	//
	// example:
	//
	// source_group_name
	ArgName *string `json:"ArgName,omitempty" xml:"ArgName,omitempty"`
	// The value of the configuration.
	//
	// example:
	//
	// 123
	ArgValue *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty"`
}

func (s DescribeRDDomainConfigResponseBodyDomainConfigsFunctionArgs) String() string {
	return dara.Prettify(s)
}

func (s DescribeRDDomainConfigResponseBodyDomainConfigsFunctionArgs) GoString() string {
	return s.String()
}

func (s *DescribeRDDomainConfigResponseBodyDomainConfigsFunctionArgs) GetArgName() *string {
	return s.ArgName
}

func (s *DescribeRDDomainConfigResponseBodyDomainConfigsFunctionArgs) GetArgValue() *string {
	return s.ArgValue
}

func (s *DescribeRDDomainConfigResponseBodyDomainConfigsFunctionArgs) SetArgName(v string) *DescribeRDDomainConfigResponseBodyDomainConfigsFunctionArgs {
	s.ArgName = &v
	return s
}

func (s *DescribeRDDomainConfigResponseBodyDomainConfigsFunctionArgs) SetArgValue(v string) *DescribeRDDomainConfigResponseBodyDomainConfigsFunctionArgs {
	s.ArgValue = &v
	return s
}

func (s *DescribeRDDomainConfigResponseBodyDomainConfigsFunctionArgs) Validate() error {
	return dara.Validate(s)
}
