// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcHoneyPotCriteriaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCriteriaList(v []*DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList) *DescribeVpcHoneyPotCriteriaResponseBody
	GetCriteriaList() []*DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList
	SetRequestId(v string) *DescribeVpcHoneyPotCriteriaResponseBody
	GetRequestId() *string
}

type DescribeVpcHoneyPotCriteriaResponseBody struct {
	// An array that consists of the search conditions.
	CriteriaList []*DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList `json:"CriteriaList,omitempty" xml:"CriteriaList,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// FCE38ADB-7361-4212-AD87-A4514E4DF925
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVpcHoneyPotCriteriaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcHoneyPotCriteriaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcHoneyPotCriteriaResponseBody) GetCriteriaList() []*DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList {
	return s.CriteriaList
}

func (s *DescribeVpcHoneyPotCriteriaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcHoneyPotCriteriaResponseBody) SetCriteriaList(v []*DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList) *DescribeVpcHoneyPotCriteriaResponseBody {
	s.CriteriaList = v
	return s
}

func (s *DescribeVpcHoneyPotCriteriaResponseBody) SetRequestId(v string) *DescribeVpcHoneyPotCriteriaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcHoneyPotCriteriaResponseBody) Validate() error {
	if s.CriteriaList != nil {
		for _, item := range s.CriteriaList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList struct {
	// The name of the search condition.
	//
	// example:
	//
	// vpcRegionId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the search condition. Valid values:
	//
	// 	- **input**: You must manually enter the search condition.
	//
	// 	- **select**: You must select a search condition from the **Values*	- drop-down list.
	//
	// example:
	//
	// select
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The values of the search condition. This parameter is returned only if the value of **Type*	- is **select**.
	//
	// > If the value of **Type*	- is **input**, the value of this parameter is an empty string.
	//
	// example:
	//
	// ap-southeast-2,eu-west-1
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList) GoString() string {
	return s.String()
}

func (s *DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList) GetName() *string {
	return s.Name
}

func (s *DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList) GetType() *string {
	return s.Type
}

func (s *DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList) GetValues() *string {
	return s.Values
}

func (s *DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList) SetName(v string) *DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList {
	s.Name = &v
	return s
}

func (s *DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList) SetType(v string) *DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList {
	s.Type = &v
	return s
}

func (s *DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList) SetValues(v string) *DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList {
	s.Values = &v
	return s
}

func (s *DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList) Validate() error {
	return dara.Validate(s)
}
