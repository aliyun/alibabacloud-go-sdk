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
	// The list of the search conditions.
	CriteriaList []*DescribeImageCriteriaResponseBodyCriteriaList `json:"CriteriaList,omitempty" xml:"CriteriaList,omitempty" type:"Repeated"`
	// The request ID.
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
	// The name of the search condition.
	//
	// - **tag**: the tag of the image
	//
	// - **digest**: the digest of the image
	//
	// - **vulStatus**: the status of the vulnerability
	//
	// - **alarmStatus**: the status of the alert
	//
	// - **riskStatus**: the status of the risk
	//
	// - **registryType**: the type of the image repository
	//
	// example:
	//
	// vulStatus
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the search condition. Valid values:
	//
	// - **input**: The search condition needs to be specified.
	//
	// - **select**: The search condition is an option that can be selected from the drop-down list.
	//
	// example:
	//
	// input
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The values of the search condition. This parameter is returned only if the value of Type is select.
	//
	// > If the value of **Type*	- is **input**, the value of this parameter is an empty string.
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
