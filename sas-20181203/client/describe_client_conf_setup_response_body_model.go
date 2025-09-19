// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientConfSetupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClientConf(v *DescribeClientConfSetupResponseBodyClientConf) *DescribeClientConfSetupResponseBody
	GetClientConf() *DescribeClientConfSetupResponseBodyClientConf
	SetRequestId(v string) *DescribeClientConfSetupResponseBody
	GetRequestId() *string
}

type DescribeClientConfSetupResponseBody struct {
	// The configurations of the Security Center agent.
	ClientConf *DescribeClientConfSetupResponseBodyClientConf `json:"ClientConf,omitempty" xml:"ClientConf,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 151F6EB6-D5F3-417A-AF7B-4D84975D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClientConfSetupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientConfSetupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClientConfSetupResponseBody) GetClientConf() *DescribeClientConfSetupResponseBodyClientConf {
	return s.ClientConf
}

func (s *DescribeClientConfSetupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClientConfSetupResponseBody) SetClientConf(v *DescribeClientConfSetupResponseBodyClientConf) *DescribeClientConfSetupResponseBody {
	s.ClientConf = v
	return s
}

func (s *DescribeClientConfSetupResponseBody) SetRequestId(v string) *DescribeClientConfSetupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClientConfSetupResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeClientConfSetupResponseBodyClientConf struct {
	// The configurations of the usage for the Security Center agent.
	//
	// example:
	//
	// {"mem":"200","cpu":"10","cpu_all":"0"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The tag that is added to the configuration.
	//
	// example:
	//
	// machineResource
	StrategyTag *string `json:"StrategyTag,omitempty" xml:"StrategyTag,omitempty"`
	// The value of the tag. Valid values:
	//
	// 	- major
	//
	// 	- advanced
	//
	// 	- basic
	//
	// example:
	//
	// major
	StrategyTagValue *string `json:"StrategyTagValue,omitempty" xml:"StrategyTagValue,omitempty"`
}

func (s DescribeClientConfSetupResponseBodyClientConf) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientConfSetupResponseBodyClientConf) GoString() string {
	return s.String()
}

func (s *DescribeClientConfSetupResponseBodyClientConf) GetConfig() *string {
	return s.Config
}

func (s *DescribeClientConfSetupResponseBodyClientConf) GetStrategyTag() *string {
	return s.StrategyTag
}

func (s *DescribeClientConfSetupResponseBodyClientConf) GetStrategyTagValue() *string {
	return s.StrategyTagValue
}

func (s *DescribeClientConfSetupResponseBodyClientConf) SetConfig(v string) *DescribeClientConfSetupResponseBodyClientConf {
	s.Config = &v
	return s
}

func (s *DescribeClientConfSetupResponseBodyClientConf) SetStrategyTag(v string) *DescribeClientConfSetupResponseBodyClientConf {
	s.StrategyTag = &v
	return s
}

func (s *DescribeClientConfSetupResponseBodyClientConf) SetStrategyTagValue(v string) *DescribeClientConfSetupResponseBodyClientConf {
	s.StrategyTagValue = &v
	return s
}

func (s *DescribeClientConfSetupResponseBodyClientConf) Validate() error {
	return dara.Validate(s)
}
