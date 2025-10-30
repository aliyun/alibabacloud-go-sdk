// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListClientUsersResponseBodyData) *ListClientUsersResponseBody
	GetData() *ListClientUsersResponseBodyData
	SetRequestId(v string) *ListClientUsersResponseBody
	GetRequestId() *string
}

type ListClientUsersResponseBody struct {
	Data *ListClientUsersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// FEF1144C-95D1-5F7C-81EF-9DB70EA49FCE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClientUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClientUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClientUsersResponseBody) GetData() *ListClientUsersResponseBodyData {
	return s.Data
}

func (s *ListClientUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClientUsersResponseBody) SetData(v *ListClientUsersResponseBodyData) *ListClientUsersResponseBody {
	s.Data = v
	return s
}

func (s *ListClientUsersResponseBody) SetRequestId(v string) *ListClientUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClientUsersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListClientUsersResponseBodyData struct {
	DataList []*ListClientUsersResponseBodyDataDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListClientUsersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListClientUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClientUsersResponseBodyData) GetDataList() []*ListClientUsersResponseBodyDataDataList {
	return s.DataList
}

func (s *ListClientUsersResponseBodyData) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *ListClientUsersResponseBodyData) SetDataList(v []*ListClientUsersResponseBodyDataDataList) *ListClientUsersResponseBodyData {
	s.DataList = v
	return s
}

func (s *ListClientUsersResponseBodyData) SetTotalNum(v int64) *ListClientUsersResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *ListClientUsersResponseBodyData) Validate() error {
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

type ListClientUsersResponseBodyDataDataList struct {
	Department *ListClientUsersResponseBodyDataDataListDepartment `json:"Department,omitempty" xml:"Department,omitempty" type:"Struct"`
	// example:
	//
	// 10800
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// johndoe@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 1970
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1026
	IdpConfigId *string `json:"IdpConfigId,omitempty" xml:"IdpConfigId,omitempty"`
	// example:
	//
	// 15800820468
	MobileNumber *string `json:"MobileNumber,omitempty" xml:"MobileNumber,omitempty"`
	// example:
	//
	// Disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// su_dead7216****
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListClientUsersResponseBodyDataDataList) String() string {
	return dara.Prettify(s)
}

func (s ListClientUsersResponseBodyDataDataList) GoString() string {
	return s.String()
}

func (s *ListClientUsersResponseBodyDataDataList) GetDepartment() *ListClientUsersResponseBodyDataDataListDepartment {
	return s.Department
}

func (s *ListClientUsersResponseBodyDataDataList) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *ListClientUsersResponseBodyDataDataList) GetDescription() *string {
	return s.Description
}

func (s *ListClientUsersResponseBodyDataDataList) GetEmail() *string {
	return s.Email
}

func (s *ListClientUsersResponseBodyDataDataList) GetId() *string {
	return s.Id
}

func (s *ListClientUsersResponseBodyDataDataList) GetIdpConfigId() *string {
	return s.IdpConfigId
}

func (s *ListClientUsersResponseBodyDataDataList) GetMobileNumber() *string {
	return s.MobileNumber
}

func (s *ListClientUsersResponseBodyDataDataList) GetStatus() *string {
	return s.Status
}

func (s *ListClientUsersResponseBodyDataDataList) GetUserId() *string {
	return s.UserId
}

func (s *ListClientUsersResponseBodyDataDataList) GetUsername() *string {
	return s.Username
}

func (s *ListClientUsersResponseBodyDataDataList) SetDepartment(v *ListClientUsersResponseBodyDataDataListDepartment) *ListClientUsersResponseBodyDataDataList {
	s.Department = v
	return s
}

func (s *ListClientUsersResponseBodyDataDataList) SetDepartmentId(v string) *ListClientUsersResponseBodyDataDataList {
	s.DepartmentId = &v
	return s
}

func (s *ListClientUsersResponseBodyDataDataList) SetDescription(v string) *ListClientUsersResponseBodyDataDataList {
	s.Description = &v
	return s
}

func (s *ListClientUsersResponseBodyDataDataList) SetEmail(v string) *ListClientUsersResponseBodyDataDataList {
	s.Email = &v
	return s
}

func (s *ListClientUsersResponseBodyDataDataList) SetId(v string) *ListClientUsersResponseBodyDataDataList {
	s.Id = &v
	return s
}

func (s *ListClientUsersResponseBodyDataDataList) SetIdpConfigId(v string) *ListClientUsersResponseBodyDataDataList {
	s.IdpConfigId = &v
	return s
}

func (s *ListClientUsersResponseBodyDataDataList) SetMobileNumber(v string) *ListClientUsersResponseBodyDataDataList {
	s.MobileNumber = &v
	return s
}

func (s *ListClientUsersResponseBodyDataDataList) SetStatus(v string) *ListClientUsersResponseBodyDataDataList {
	s.Status = &v
	return s
}

func (s *ListClientUsersResponseBodyDataDataList) SetUserId(v string) *ListClientUsersResponseBodyDataDataList {
	s.UserId = &v
	return s
}

func (s *ListClientUsersResponseBodyDataDataList) SetUsername(v string) *ListClientUsersResponseBodyDataDataList {
	s.Username = &v
	return s
}

func (s *ListClientUsersResponseBodyDataDataList) Validate() error {
	if s.Department != nil {
		if err := s.Department.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListClientUsersResponseBodyDataDataListDepartment struct {
	// example:
	//
	// 105
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListClientUsersResponseBodyDataDataListDepartment) String() string {
	return dara.Prettify(s)
}

func (s ListClientUsersResponseBodyDataDataListDepartment) GoString() string {
	return s.String()
}

func (s *ListClientUsersResponseBodyDataDataListDepartment) GetId() *string {
	return s.Id
}

func (s *ListClientUsersResponseBodyDataDataListDepartment) GetName() *string {
	return s.Name
}

func (s *ListClientUsersResponseBodyDataDataListDepartment) SetId(v string) *ListClientUsersResponseBodyDataDataListDepartment {
	s.Id = &v
	return s
}

func (s *ListClientUsersResponseBodyDataDataListDepartment) SetName(v string) *ListClientUsersResponseBodyDataDataListDepartment {
	s.Name = &v
	return s
}

func (s *ListClientUsersResponseBodyDataDataListDepartment) Validate() error {
	return dara.Validate(s)
}
