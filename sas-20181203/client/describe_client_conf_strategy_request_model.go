// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientConfStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTag(v string) *DescribeClientConfStrategyRequest
	GetTag() *string
	SetTagValue(v string) *DescribeClientConfStrategyRequest
	GetTagValue() *string
}

type DescribeClientConfStrategyRequest struct {
	// The tag that is added to the server.
	//
	// This parameter is required.
	//
	// example:
	//
	// machineResource
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
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
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeClientConfStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientConfStrategyRequest) GoString() string {
	return s.String()
}

func (s *DescribeClientConfStrategyRequest) GetTag() *string {
	return s.Tag
}

func (s *DescribeClientConfStrategyRequest) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeClientConfStrategyRequest) SetTag(v string) *DescribeClientConfStrategyRequest {
	s.Tag = &v
	return s
}

func (s *DescribeClientConfStrategyRequest) SetTagValue(v string) *DescribeClientConfStrategyRequest {
	s.TagValue = &v
	return s
}

func (s *DescribeClientConfStrategyRequest) Validate() error {
	return dara.Validate(s)
}
