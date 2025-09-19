// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientConfSetupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStrategyTag(v string) *DescribeClientConfSetupRequest
	GetStrategyTag() *string
	SetStrategyTagValue(v string) *DescribeClientConfSetupRequest
	GetStrategyTagValue() *string
}

type DescribeClientConfSetupRequest struct {
	// The tag that is added to the server.
	//
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// major
	StrategyTagValue *string `json:"StrategyTagValue,omitempty" xml:"StrategyTagValue,omitempty"`
}

func (s DescribeClientConfSetupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientConfSetupRequest) GoString() string {
	return s.String()
}

func (s *DescribeClientConfSetupRequest) GetStrategyTag() *string {
	return s.StrategyTag
}

func (s *DescribeClientConfSetupRequest) GetStrategyTagValue() *string {
	return s.StrategyTagValue
}

func (s *DescribeClientConfSetupRequest) SetStrategyTag(v string) *DescribeClientConfSetupRequest {
	s.StrategyTag = &v
	return s
}

func (s *DescribeClientConfSetupRequest) SetStrategyTagValue(v string) *DescribeClientConfSetupRequest {
	s.StrategyTagValue = &v
	return s
}

func (s *DescribeClientConfSetupRequest) Validate() error {
	return dara.Validate(s)
}
