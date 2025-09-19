// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageRepoCriteriaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCriteriaList(v []*DescribeImageRepoCriteriaResponseBodyCriteriaList) *DescribeImageRepoCriteriaResponseBody
	GetCriteriaList() []*DescribeImageRepoCriteriaResponseBodyCriteriaList
	SetRequestId(v string) *DescribeImageRepoCriteriaResponseBody
	GetRequestId() *string
}

type DescribeImageRepoCriteriaResponseBody struct {
	// An array consisting of the filter conditions that are supported by the image repository.
	CriteriaList []*DescribeImageRepoCriteriaResponseBodyCriteriaList `json:"CriteriaList,omitempty" xml:"CriteriaList,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// F35F45B0-5D6B-4238-BE02-A62D0760E840
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageRepoCriteriaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageRepoCriteriaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageRepoCriteriaResponseBody) GetCriteriaList() []*DescribeImageRepoCriteriaResponseBodyCriteriaList {
	return s.CriteriaList
}

func (s *DescribeImageRepoCriteriaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageRepoCriteriaResponseBody) SetCriteriaList(v []*DescribeImageRepoCriteriaResponseBodyCriteriaList) *DescribeImageRepoCriteriaResponseBody {
	s.CriteriaList = v
	return s
}

func (s *DescribeImageRepoCriteriaResponseBody) SetRequestId(v string) *DescribeImageRepoCriteriaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageRepoCriteriaResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeImageRepoCriteriaResponseBodyCriteriaList struct {
	// The name of the search condition. Valid values:
	//
	// 	- **instanceId**: the ID of the image instance.
	//
	// 	- **repoName**: the name of the image repository.
	//
	// 	- **repoId**: the ID of the image repository.
	//
	// 	- **repoNamespace**: the namespace of the image repository.
	//
	// 	- **regionId**: the region in which the image resides.
	//
	// 	- **vulStatus**: indicates whether vulnerabilities exist.
	//
	// 	- **alarmStatus**: indicates whether security alerts exist.
	//
	// 	- **hcStatus**: indicates whether baseline risks exist.
	//
	// 	- **riskStatus**: indicates whether risks exist.
	//
	// 	- **registryType**: the type of the image repository.
	//
	// 	- **ImageId**: the image ID.
	//
	// 	- **tag**: the image tag.
	//
	// example:
	//
	// vulStatus
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the search condition. Valid values:
	//
	// 	- **input**: The search condition needs to be specified.
	//
	// 	- **select**: The search condition is an option that can be selected from the drop-down list.
	//
	// example:
	//
	// select
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The values of the search condition. This parameter is returned only if the value of **Type*	- is set to **select**.
	//
	// > If the value of **Type*	- is set to **input**, the return value of this parameter is empty.
	//
	// example:
	//
	// NO,YES
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeImageRepoCriteriaResponseBodyCriteriaList) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageRepoCriteriaResponseBodyCriteriaList) GoString() string {
	return s.String()
}

func (s *DescribeImageRepoCriteriaResponseBodyCriteriaList) GetName() *string {
	return s.Name
}

func (s *DescribeImageRepoCriteriaResponseBodyCriteriaList) GetType() *string {
	return s.Type
}

func (s *DescribeImageRepoCriteriaResponseBodyCriteriaList) GetValues() *string {
	return s.Values
}

func (s *DescribeImageRepoCriteriaResponseBodyCriteriaList) SetName(v string) *DescribeImageRepoCriteriaResponseBodyCriteriaList {
	s.Name = &v
	return s
}

func (s *DescribeImageRepoCriteriaResponseBodyCriteriaList) SetType(v string) *DescribeImageRepoCriteriaResponseBodyCriteriaList {
	s.Type = &v
	return s
}

func (s *DescribeImageRepoCriteriaResponseBodyCriteriaList) SetValues(v string) *DescribeImageRepoCriteriaResponseBodyCriteriaList {
	s.Values = &v
	return s
}

func (s *DescribeImageRepoCriteriaResponseBodyCriteriaList) Validate() error {
	return dara.Validate(s)
}
