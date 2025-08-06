// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainStagingConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainConfigs(v []*DescribeVodDomainStagingConfigResponseBodyDomainConfigs) *DescribeVodDomainStagingConfigResponseBody
	GetDomainConfigs() []*DescribeVodDomainStagingConfigResponseBodyDomainConfigs
	SetDomainName(v string) *DescribeVodDomainStagingConfigResponseBody
	GetDomainName() *string
	SetRequestId(v string) *DescribeVodDomainStagingConfigResponseBody
	GetRequestId() *string
}

type DescribeVodDomainStagingConfigResponseBody struct {
	DomainConfigs []*DescribeVodDomainStagingConfigResponseBodyDomainConfigs `json:"DomainConfigs,omitempty" xml:"DomainConfigs,omitempty" type:"Repeated"`
	DomainName    *string                                                    `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	RequestId     *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodDomainStagingConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainStagingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainStagingConfigResponseBody) GetDomainConfigs() []*DescribeVodDomainStagingConfigResponseBodyDomainConfigs {
	return s.DomainConfigs
}

func (s *DescribeVodDomainStagingConfigResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainStagingConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainStagingConfigResponseBody) SetDomainConfigs(v []*DescribeVodDomainStagingConfigResponseBodyDomainConfigs) *DescribeVodDomainStagingConfigResponseBody {
	s.DomainConfigs = v
	return s
}

func (s *DescribeVodDomainStagingConfigResponseBody) SetDomainName(v string) *DescribeVodDomainStagingConfigResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainStagingConfigResponseBody) SetRequestId(v string) *DescribeVodDomainStagingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainStagingConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainStagingConfigResponseBodyDomainConfigs struct {
	ConfigId     *string                                                                `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	FunctionArgs []*DescribeVodDomainStagingConfigResponseBodyDomainConfigsFunctionArgs `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty" type:"Repeated"`
	FunctionName *string                                                                `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	ParentId     *string                                                                `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Status       *string                                                                `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeVodDomainStagingConfigResponseBodyDomainConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainStagingConfigResponseBodyDomainConfigs) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainStagingConfigResponseBodyDomainConfigs) GetConfigId() *string {
	return s.ConfigId
}

func (s *DescribeVodDomainStagingConfigResponseBodyDomainConfigs) GetFunctionArgs() []*DescribeVodDomainStagingConfigResponseBodyDomainConfigsFunctionArgs {
	return s.FunctionArgs
}

func (s *DescribeVodDomainStagingConfigResponseBodyDomainConfigs) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DescribeVodDomainStagingConfigResponseBodyDomainConfigs) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeVodDomainStagingConfigResponseBodyDomainConfigs) GetStatus() *string {
	return s.Status
}

func (s *DescribeVodDomainStagingConfigResponseBodyDomainConfigs) SetConfigId(v string) *DescribeVodDomainStagingConfigResponseBodyDomainConfigs {
	s.ConfigId = &v
	return s
}

func (s *DescribeVodDomainStagingConfigResponseBodyDomainConfigs) SetFunctionArgs(v []*DescribeVodDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) *DescribeVodDomainStagingConfigResponseBodyDomainConfigs {
	s.FunctionArgs = v
	return s
}

func (s *DescribeVodDomainStagingConfigResponseBodyDomainConfigs) SetFunctionName(v string) *DescribeVodDomainStagingConfigResponseBodyDomainConfigs {
	s.FunctionName = &v
	return s
}

func (s *DescribeVodDomainStagingConfigResponseBodyDomainConfigs) SetParentId(v string) *DescribeVodDomainStagingConfigResponseBodyDomainConfigs {
	s.ParentId = &v
	return s
}

func (s *DescribeVodDomainStagingConfigResponseBodyDomainConfigs) SetStatus(v string) *DescribeVodDomainStagingConfigResponseBodyDomainConfigs {
	s.Status = &v
	return s
}

func (s *DescribeVodDomainStagingConfigResponseBodyDomainConfigs) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainStagingConfigResponseBodyDomainConfigsFunctionArgs struct {
	ArgName  *string `json:"ArgName,omitempty" xml:"ArgName,omitempty"`
	ArgValue *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty"`
}

func (s DescribeVodDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) GetArgName() *string {
	return s.ArgName
}

func (s *DescribeVodDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) GetArgValue() *string {
	return s.ArgValue
}

func (s *DescribeVodDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) SetArgName(v string) *DescribeVodDomainStagingConfigResponseBodyDomainConfigsFunctionArgs {
	s.ArgName = &v
	return s
}

func (s *DescribeVodDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) SetArgValue(v string) *DescribeVodDomainStagingConfigResponseBodyDomainConfigsFunctionArgs {
	s.ArgValue = &v
	return s
}

func (s *DescribeVodDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) Validate() error {
	return dara.Validate(s)
}
