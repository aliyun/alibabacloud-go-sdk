// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQueryConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*DescribeQueryConfigsResponseBodyConfigs) *DescribeQueryConfigsResponseBody
	GetConfigs() []*DescribeQueryConfigsResponseBodyConfigs
	SetRequestId(v string) *DescribeQueryConfigsResponseBody
	GetRequestId() *string
}

type DescribeQueryConfigsResponseBody struct {
	Configs   []*DescribeQueryConfigsResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeQueryConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeQueryConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQueryConfigsResponseBody) GetConfigs() []*DescribeQueryConfigsResponseBodyConfigs {
	return s.Configs
}

func (s *DescribeQueryConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeQueryConfigsResponseBody) SetConfigs(v []*DescribeQueryConfigsResponseBodyConfigs) *DescribeQueryConfigsResponseBody {
	s.Configs = v
	return s
}

func (s *DescribeQueryConfigsResponseBody) SetRequestId(v string) *DescribeQueryConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeQueryConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeQueryConfigsResponseBodyConfigs struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeQueryConfigsResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeQueryConfigsResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *DescribeQueryConfigsResponseBodyConfigs) GetValue() *string {
	return s.Value
}

func (s *DescribeQueryConfigsResponseBodyConfigs) SetValue(v string) *DescribeQueryConfigsResponseBodyConfigs {
	s.Value = &v
	return s
}

func (s *DescribeQueryConfigsResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}
