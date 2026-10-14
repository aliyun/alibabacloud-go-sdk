// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRiskLevelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListRiskLevelsResponseBody
	GetRequestId() *string
	SetRiskLevelList(v []*ListRiskLevelsResponseBodyRiskLevelList) *ListRiskLevelsResponseBody
	GetRiskLevelList() []*ListRiskLevelsResponseBodyRiskLevelList
}

type ListRiskLevelsResponseBody struct {
	// The request ID. You can use this ID to locate the API call.
	//
	// example:
	//
	// 7C6D8E9F-1234-5678-ABCD-0123456789AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of sensitivity levels.
	RiskLevelList []*ListRiskLevelsResponseBodyRiskLevelList `json:"RiskLevelList,omitempty" xml:"RiskLevelList,omitempty" type:"Repeated"`
}

func (s ListRiskLevelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRiskLevelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRiskLevelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRiskLevelsResponseBody) GetRiskLevelList() []*ListRiskLevelsResponseBodyRiskLevelList {
	return s.RiskLevelList
}

func (s *ListRiskLevelsResponseBody) SetRequestId(v string) *ListRiskLevelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRiskLevelsResponseBody) SetRiskLevelList(v []*ListRiskLevelsResponseBodyRiskLevelList) *ListRiskLevelsResponseBody {
	s.RiskLevelList = v
	return s
}

func (s *ListRiskLevelsResponseBody) Validate() error {
	if s.RiskLevelList != nil {
		for _, item := range s.RiskLevelList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRiskLevelsResponseBodyRiskLevelList struct {
	// The description of the sensitivity level.
	//
	// example:
	//
	// Personal information data
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the sensitivity level.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the sensitivity level.
	//
	// example:
	//
	// S2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of references to the sensitivity level.
	//
	// example:
	//
	// 10
	ReferenceNum *int32 `json:"ReferenceNum,omitempty" xml:"ReferenceNum,omitempty"`
}

func (s ListRiskLevelsResponseBodyRiskLevelList) String() string {
	return dara.Prettify(s)
}

func (s ListRiskLevelsResponseBodyRiskLevelList) GoString() string {
	return s.String()
}

func (s *ListRiskLevelsResponseBodyRiskLevelList) GetDescription() *string {
	return s.Description
}

func (s *ListRiskLevelsResponseBodyRiskLevelList) GetId() *int64 {
	return s.Id
}

func (s *ListRiskLevelsResponseBodyRiskLevelList) GetName() *string {
	return s.Name
}

func (s *ListRiskLevelsResponseBodyRiskLevelList) GetReferenceNum() *int32 {
	return s.ReferenceNum
}

func (s *ListRiskLevelsResponseBodyRiskLevelList) SetDescription(v string) *ListRiskLevelsResponseBodyRiskLevelList {
	s.Description = &v
	return s
}

func (s *ListRiskLevelsResponseBodyRiskLevelList) SetId(v int64) *ListRiskLevelsResponseBodyRiskLevelList {
	s.Id = &v
	return s
}

func (s *ListRiskLevelsResponseBodyRiskLevelList) SetName(v string) *ListRiskLevelsResponseBodyRiskLevelList {
	s.Name = &v
	return s
}

func (s *ListRiskLevelsResponseBodyRiskLevelList) SetReferenceNum(v int32) *ListRiskLevelsResponseBodyRiskLevelList {
	s.ReferenceNum = &v
	return s
}

func (s *ListRiskLevelsResponseBodyRiskLevelList) Validate() error {
	return dara.Validate(s)
}
