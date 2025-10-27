// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIdcAssetCriteriaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCriteriaList(v []*DescribeIdcAssetCriteriaResponseBodyCriteriaList) *DescribeIdcAssetCriteriaResponseBody
	GetCriteriaList() []*DescribeIdcAssetCriteriaResponseBodyCriteriaList
	SetRequestId(v string) *DescribeIdcAssetCriteriaResponseBody
	GetRequestId() *string
}

type DescribeIdcAssetCriteriaResponseBody struct {
	// The information about the asset search conditions.
	CriteriaList []*DescribeIdcAssetCriteriaResponseBodyCriteriaList `json:"CriteriaList,omitempty" xml:"CriteriaList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 11C96623-E106-59C9-866D-A6C82911****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIdcAssetCriteriaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIdcAssetCriteriaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIdcAssetCriteriaResponseBody) GetCriteriaList() []*DescribeIdcAssetCriteriaResponseBodyCriteriaList {
	return s.CriteriaList
}

func (s *DescribeIdcAssetCriteriaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIdcAssetCriteriaResponseBody) SetCriteriaList(v []*DescribeIdcAssetCriteriaResponseBodyCriteriaList) *DescribeIdcAssetCriteriaResponseBody {
	s.CriteriaList = v
	return s
}

func (s *DescribeIdcAssetCriteriaResponseBody) SetRequestId(v string) *DescribeIdcAssetCriteriaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIdcAssetCriteriaResponseBody) Validate() error {
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

type DescribeIdcAssetCriteriaResponseBodyCriteriaList struct {
	// The name of the search condition.
	//
	// example:
	//
	// scannedIp
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
	// The attribute values of the assets that match the keyword.
	//
	// example:
	//
	// 1.1.1.*
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeIdcAssetCriteriaResponseBodyCriteriaList) String() string {
	return dara.Prettify(s)
}

func (s DescribeIdcAssetCriteriaResponseBodyCriteriaList) GoString() string {
	return s.String()
}

func (s *DescribeIdcAssetCriteriaResponseBodyCriteriaList) GetName() *string {
	return s.Name
}

func (s *DescribeIdcAssetCriteriaResponseBodyCriteriaList) GetType() *string {
	return s.Type
}

func (s *DescribeIdcAssetCriteriaResponseBodyCriteriaList) GetValues() *string {
	return s.Values
}

func (s *DescribeIdcAssetCriteriaResponseBodyCriteriaList) SetName(v string) *DescribeIdcAssetCriteriaResponseBodyCriteriaList {
	s.Name = &v
	return s
}

func (s *DescribeIdcAssetCriteriaResponseBodyCriteriaList) SetType(v string) *DescribeIdcAssetCriteriaResponseBodyCriteriaList {
	s.Type = &v
	return s
}

func (s *DescribeIdcAssetCriteriaResponseBodyCriteriaList) SetValues(v string) *DescribeIdcAssetCriteriaResponseBodyCriteriaList {
	s.Values = &v
	return s
}

func (s *DescribeIdcAssetCriteriaResponseBodyCriteriaList) Validate() error {
	return dara.Validate(s)
}
