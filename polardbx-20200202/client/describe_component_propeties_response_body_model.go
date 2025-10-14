// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentPropetiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProperties(v []*DescribeComponentPropetiesResponseBodyProperties) *DescribeComponentPropetiesResponseBody
	GetProperties() []*DescribeComponentPropetiesResponseBodyProperties
	SetRequestId(v string) *DescribeComponentPropetiesResponseBody
	GetRequestId() *string
}

type DescribeComponentPropetiesResponseBody struct {
	Properties []*DescribeComponentPropetiesResponseBodyProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// A501A191-BD70-5E50-98A9-C2A486A82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeComponentPropetiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentPropetiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComponentPropetiesResponseBody) GetProperties() []*DescribeComponentPropetiesResponseBodyProperties {
	return s.Properties
}

func (s *DescribeComponentPropetiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeComponentPropetiesResponseBody) SetProperties(v []*DescribeComponentPropetiesResponseBodyProperties) *DescribeComponentPropetiesResponseBody {
	s.Properties = v
	return s
}

func (s *DescribeComponentPropetiesResponseBody) SetRequestId(v string) *DescribeComponentPropetiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeComponentPropetiesResponseBody) Validate() error {
	if s.Properties != nil {
		for _, item := range s.Properties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeComponentPropetiesResponseBodyProperties struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	OrderIndex *string `json:"OrderIndex,omitempty" xml:"OrderIndex,omitempty"`
	// example:
	//
	// dn_node_spec
	PropertyCode *string `json:"PropertyCode,omitempty" xml:"PropertyCode,omitempty"`
	// example:
	//
	// 172.27.35.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeComponentPropetiesResponseBodyProperties) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentPropetiesResponseBodyProperties) GoString() string {
	return s.String()
}

func (s *DescribeComponentPropetiesResponseBodyProperties) GetName() *string {
	return s.Name
}

func (s *DescribeComponentPropetiesResponseBodyProperties) GetOrderIndex() *string {
	return s.OrderIndex
}

func (s *DescribeComponentPropetiesResponseBodyProperties) GetPropertyCode() *string {
	return s.PropertyCode
}

func (s *DescribeComponentPropetiesResponseBodyProperties) GetValue() *string {
	return s.Value
}

func (s *DescribeComponentPropetiesResponseBodyProperties) SetName(v string) *DescribeComponentPropetiesResponseBodyProperties {
	s.Name = &v
	return s
}

func (s *DescribeComponentPropetiesResponseBodyProperties) SetOrderIndex(v string) *DescribeComponentPropetiesResponseBodyProperties {
	s.OrderIndex = &v
	return s
}

func (s *DescribeComponentPropetiesResponseBodyProperties) SetPropertyCode(v string) *DescribeComponentPropetiesResponseBodyProperties {
	s.PropertyCode = &v
	return s
}

func (s *DescribeComponentPropetiesResponseBodyProperties) SetValue(v string) *DescribeComponentPropetiesResponseBodyProperties {
	s.Value = &v
	return s
}

func (s *DescribeComponentPropetiesResponseBodyProperties) Validate() error {
	return dara.Validate(s)
}
