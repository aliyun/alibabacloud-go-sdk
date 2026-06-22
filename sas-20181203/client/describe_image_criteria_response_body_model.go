// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageCriteriaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCriteriaList(v []*DescribeImageCriteriaResponseBodyCriteriaList) *DescribeImageCriteriaResponseBody
	GetCriteriaList() []*DescribeImageCriteriaResponseBodyCriteriaList
	SetRequestId(v string) *DescribeImageCriteriaResponseBody
	GetRequestId() *string
}

type DescribeImageCriteriaResponseBody struct {
	// The list of image query criteria.
	CriteriaList []*DescribeImageCriteriaResponseBodyCriteriaList `json:"CriteriaList,omitempty" xml:"CriteriaList,omitempty" type:"Repeated"`
	// The request ID. Alibaba Cloud generates a unique identifier for each API request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 24A20733-10A0-4AF6-BE6B-E3322413BB68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageCriteriaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageCriteriaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageCriteriaResponseBody) GetCriteriaList() []*DescribeImageCriteriaResponseBodyCriteriaList {
	return s.CriteriaList
}

func (s *DescribeImageCriteriaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageCriteriaResponseBody) SetCriteriaList(v []*DescribeImageCriteriaResponseBodyCriteriaList) *DescribeImageCriteriaResponseBody {
	s.CriteriaList = v
	return s
}

func (s *DescribeImageCriteriaResponseBody) SetRequestId(v string) *DescribeImageCriteriaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageCriteriaResponseBody) Validate() error {
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

type DescribeImageCriteriaResponseBodyCriteriaList struct {
	// The name of the query criterion.
	//
	// - **tag**: image tag.
	//
	// - **digest**: image digest.
	//
	// - **vulStatus**: vulnerability status.
	//
	// - **alarmStatus**: security alert status.
	//
	// - **riskStatus**: risk status.
	//
	// - **registryType**: image repository type.
	//
	// example:
	//
	// vulStatus
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the query criterion. Valid values:
	//
	// - **input**: requires manual input of the query field.
	//
	// - **select**: requires selecting a subtype from a drop-down list.
	//
	// example:
	//
	// input
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The available option values when **Type*	- (the type of the query criterion) is **select**.
	//
	// > When **Type*	- (the type of the query criterion) is **input**, this parameter returns an empty value.
	//
	// example:
	//
	// NO,YES
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeImageCriteriaResponseBodyCriteriaList) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageCriteriaResponseBodyCriteriaList) GoString() string {
	return s.String()
}

func (s *DescribeImageCriteriaResponseBodyCriteriaList) GetName() *string {
	return s.Name
}

func (s *DescribeImageCriteriaResponseBodyCriteriaList) GetType() *string {
	return s.Type
}

func (s *DescribeImageCriteriaResponseBodyCriteriaList) GetValues() *string {
	return s.Values
}

func (s *DescribeImageCriteriaResponseBodyCriteriaList) SetName(v string) *DescribeImageCriteriaResponseBodyCriteriaList {
	s.Name = &v
	return s
}

func (s *DescribeImageCriteriaResponseBodyCriteriaList) SetType(v string) *DescribeImageCriteriaResponseBodyCriteriaList {
	s.Type = &v
	return s
}

func (s *DescribeImageCriteriaResponseBodyCriteriaList) SetValues(v string) *DescribeImageCriteriaResponseBodyCriteriaList {
	s.Values = &v
	return s
}

func (s *DescribeImageCriteriaResponseBodyCriteriaList) Validate() error {
	return dara.Validate(s)
}
