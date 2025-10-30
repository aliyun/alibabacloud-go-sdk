// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdpDepartmentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListIdpDepartmentsResponseBodyData) *ListIdpDepartmentsResponseBody
	GetData() *ListIdpDepartmentsResponseBodyData
	SetRequestId(v string) *ListIdpDepartmentsResponseBody
	GetRequestId() *string
}

type ListIdpDepartmentsResponseBody struct {
	Data *ListIdpDepartmentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 102350E7-1A20-58F5-9D63-ABEA820AE6E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListIdpDepartmentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIdpDepartmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIdpDepartmentsResponseBody) GetData() *ListIdpDepartmentsResponseBodyData {
	return s.Data
}

func (s *ListIdpDepartmentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIdpDepartmentsResponseBody) SetData(v *ListIdpDepartmentsResponseBodyData) *ListIdpDepartmentsResponseBody {
	s.Data = v
	return s
}

func (s *ListIdpDepartmentsResponseBody) SetRequestId(v string) *ListIdpDepartmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIdpDepartmentsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListIdpDepartmentsResponseBodyData struct {
	DataList []*ListIdpDepartmentsResponseBodyDataDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListIdpDepartmentsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListIdpDepartmentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIdpDepartmentsResponseBodyData) GetDataList() []*ListIdpDepartmentsResponseBodyDataDataList {
	return s.DataList
}

func (s *ListIdpDepartmentsResponseBodyData) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *ListIdpDepartmentsResponseBodyData) SetDataList(v []*ListIdpDepartmentsResponseBodyDataDataList) *ListIdpDepartmentsResponseBodyData {
	s.DataList = v
	return s
}

func (s *ListIdpDepartmentsResponseBodyData) SetTotalNum(v int64) *ListIdpDepartmentsResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *ListIdpDepartmentsResponseBodyData) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIdpDepartmentsResponseBodyDataDataList struct {
	// example:
	//
	// 30520
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1440
	IdpConfigId *string `json:"IdpConfigId,omitempty" xml:"IdpConfigId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListIdpDepartmentsResponseBodyDataDataList) String() string {
	return dara.Prettify(s)
}

func (s ListIdpDepartmentsResponseBodyDataDataList) GoString() string {
	return s.String()
}

func (s *ListIdpDepartmentsResponseBodyDataDataList) GetId() *string {
	return s.Id
}

func (s *ListIdpDepartmentsResponseBodyDataDataList) GetIdpConfigId() *string {
	return s.IdpConfigId
}

func (s *ListIdpDepartmentsResponseBodyDataDataList) GetName() *string {
	return s.Name
}

func (s *ListIdpDepartmentsResponseBodyDataDataList) SetId(v string) *ListIdpDepartmentsResponseBodyDataDataList {
	s.Id = &v
	return s
}

func (s *ListIdpDepartmentsResponseBodyDataDataList) SetIdpConfigId(v string) *ListIdpDepartmentsResponseBodyDataDataList {
	s.IdpConfigId = &v
	return s
}

func (s *ListIdpDepartmentsResponseBodyDataDataList) SetName(v string) *ListIdpDepartmentsResponseBodyDataDataList {
	s.Name = &v
	return s
}

func (s *ListIdpDepartmentsResponseBodyDataDataList) Validate() error {
	return dara.Validate(s)
}
