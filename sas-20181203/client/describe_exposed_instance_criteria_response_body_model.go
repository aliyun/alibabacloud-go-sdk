// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExposedInstanceCriteriaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCriteriaList(v []*DescribeExposedInstanceCriteriaResponseBodyCriteriaList) *DescribeExposedInstanceCriteriaResponseBody
	GetCriteriaList() []*DescribeExposedInstanceCriteriaResponseBodyCriteriaList
	SetRequestId(v string) *DescribeExposedInstanceCriteriaResponseBody
	GetRequestId() *string
}

type DescribeExposedInstanceCriteriaResponseBody struct {
	// The search conditions that are used to search for exposed assets.
	CriteriaList []*DescribeExposedInstanceCriteriaResponseBodyCriteriaList `json:"CriteriaList,omitempty" xml:"CriteriaList,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 6D9CDB47-6191-4415-BE63-7E8B12CD4FBE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeExposedInstanceCriteriaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedInstanceCriteriaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceCriteriaResponseBody) GetCriteriaList() []*DescribeExposedInstanceCriteriaResponseBodyCriteriaList {
	return s.CriteriaList
}

func (s *DescribeExposedInstanceCriteriaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExposedInstanceCriteriaResponseBody) SetCriteriaList(v []*DescribeExposedInstanceCriteriaResponseBodyCriteriaList) *DescribeExposedInstanceCriteriaResponseBody {
	s.CriteriaList = v
	return s
}

func (s *DescribeExposedInstanceCriteriaResponseBody) SetRequestId(v string) *DescribeExposedInstanceCriteriaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExposedInstanceCriteriaResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeExposedInstanceCriteriaResponseBodyCriteriaList struct {
	// The name of the search condition.
	//
	// example:
	//
	// instanceId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the search condition. Valid values:
	//
	// 	- **input**: You must configure the search condition.
	//
	// 	- **select**: You must select a search condition from the **Values*	- list.
	//
	// example:
	//
	// select
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the search condition. This parameter is returned only when the value of the **Type*	- parameter is **select**.
	//
	// >  If the value of the **Type*	- parameter is **input**, this parameter is empty.
	//
	// example:
	//
	// i-bp19r0fdd39idxhf****
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeExposedInstanceCriteriaResponseBodyCriteriaList) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedInstanceCriteriaResponseBodyCriteriaList) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceCriteriaResponseBodyCriteriaList) GetName() *string {
	return s.Name
}

func (s *DescribeExposedInstanceCriteriaResponseBodyCriteriaList) GetType() *string {
	return s.Type
}

func (s *DescribeExposedInstanceCriteriaResponseBodyCriteriaList) GetValues() *string {
	return s.Values
}

func (s *DescribeExposedInstanceCriteriaResponseBodyCriteriaList) SetName(v string) *DescribeExposedInstanceCriteriaResponseBodyCriteriaList {
	s.Name = &v
	return s
}

func (s *DescribeExposedInstanceCriteriaResponseBodyCriteriaList) SetType(v string) *DescribeExposedInstanceCriteriaResponseBodyCriteriaList {
	s.Type = &v
	return s
}

func (s *DescribeExposedInstanceCriteriaResponseBodyCriteriaList) SetValues(v string) *DescribeExposedInstanceCriteriaResponseBodyCriteriaList {
	s.Values = &v
	return s
}

func (s *DescribeExposedInstanceCriteriaResponseBodyCriteriaList) Validate() error {
	return dara.Validate(s)
}
