// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigList(v []*DescribeConfigsResponseBodyConfigList) *DescribeConfigsResponseBody
	GetConfigList() []*DescribeConfigsResponseBodyConfigList
	SetRequestId(v string) *DescribeConfigsResponseBody
	GetRequestId() *string
}

type DescribeConfigsResponseBody struct {
	// An array that consists of common configuration items for alerts.
	ConfigList []*DescribeConfigsResponseBodyConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConfigsResponseBody) GetConfigList() []*DescribeConfigsResponseBodyConfigList {
	return s.ConfigList
}

func (s *DescribeConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeConfigsResponseBody) SetConfigList(v []*DescribeConfigsResponseBodyConfigList) *DescribeConfigsResponseBody {
	s.ConfigList = v
	return s
}

func (s *DescribeConfigsResponseBody) SetRequestId(v string) *DescribeConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConfigsResponseBody) Validate() error {
	if s.ConfigList != nil {
		for _, item := range s.ConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeConfigsResponseBodyConfigList struct {
	// The code of the common configuration item.
	//
	// example:
	//
	// 1
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description of the default value for the common configuration item.
	//
	// example:
	//
	// The volume of logs of a specific type that are generated on the current day is less than 30% of the average volume of logs generated in the previous 10 days.
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The description of the common configuration item.
	//
	// example:
	//
	// Anomalous log output
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique ID of the common configuration item.
	//
	// example:
	//
	// 2133
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The value of the common configuration item.
	//
	// example:
	//
	// 30
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeConfigsResponseBodyConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigsResponseBodyConfigList) GoString() string {
	return s.String()
}

func (s *DescribeConfigsResponseBodyConfigList) GetCode() *string {
	return s.Code
}

func (s *DescribeConfigsResponseBodyConfigList) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *DescribeConfigsResponseBodyConfigList) GetDescription() *string {
	return s.Description
}

func (s *DescribeConfigsResponseBodyConfigList) GetId() *int64 {
	return s.Id
}

func (s *DescribeConfigsResponseBodyConfigList) GetValue() *string {
	return s.Value
}

func (s *DescribeConfigsResponseBodyConfigList) SetCode(v string) *DescribeConfigsResponseBodyConfigList {
	s.Code = &v
	return s
}

func (s *DescribeConfigsResponseBodyConfigList) SetDefaultValue(v string) *DescribeConfigsResponseBodyConfigList {
	s.DefaultValue = &v
	return s
}

func (s *DescribeConfigsResponseBodyConfigList) SetDescription(v string) *DescribeConfigsResponseBodyConfigList {
	s.Description = &v
	return s
}

func (s *DescribeConfigsResponseBodyConfigList) SetId(v int64) *DescribeConfigsResponseBodyConfigList {
	s.Id = &v
	return s
}

func (s *DescribeConfigsResponseBodyConfigList) SetValue(v string) *DescribeConfigsResponseBodyConfigList {
	s.Value = &v
	return s
}

func (s *DescribeConfigsResponseBodyConfigList) Validate() error {
	return dara.Validate(s)
}
