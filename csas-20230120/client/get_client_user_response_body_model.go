// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetClientUserResponseBodyData) *GetClientUserResponseBody
	GetData() *GetClientUserResponseBodyData
	SetRequestId(v string) *GetClientUserResponseBody
	GetRequestId() *string
}

type GetClientUserResponseBody struct {
	Data *GetClientUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetClientUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClientUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetClientUserResponseBody) GetData() *GetClientUserResponseBodyData {
	return s.Data
}

func (s *GetClientUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClientUserResponseBody) SetData(v *GetClientUserResponseBodyData) *GetClientUserResponseBody {
	s.Data = v
	return s
}

func (s *GetClientUserResponseBody) SetRequestId(v string) *GetClientUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClientUserResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetClientUserResponseBodyData struct {
	Department *GetClientUserResponseBodyDataDepartment `json:"Department,omitempty" xml:"Department,omitempty" type:"Struct"`
	// example:
	//
	// 10713
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// johndoe@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 83
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 598
	IdpConfigId *string `json:"IdpConfigId,omitempty" xml:"IdpConfigId,omitempty"`
	// example:
	//
	// 13641966835
	MobileNumber *string `json:"MobileNumber,omitempty" xml:"MobileNumber,omitempty"`
	// example:
	//
	// Disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// su_abcd7215****
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetClientUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetClientUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetClientUserResponseBodyData) GetDepartment() *GetClientUserResponseBodyDataDepartment {
	return s.Department
}

func (s *GetClientUserResponseBodyData) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *GetClientUserResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetClientUserResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *GetClientUserResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetClientUserResponseBodyData) GetIdpConfigId() *string {
	return s.IdpConfigId
}

func (s *GetClientUserResponseBodyData) GetMobileNumber() *string {
	return s.MobileNumber
}

func (s *GetClientUserResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetClientUserResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *GetClientUserResponseBodyData) GetUsername() *string {
	return s.Username
}

func (s *GetClientUserResponseBodyData) SetDepartment(v *GetClientUserResponseBodyDataDepartment) *GetClientUserResponseBodyData {
	s.Department = v
	return s
}

func (s *GetClientUserResponseBodyData) SetDepartmentId(v string) *GetClientUserResponseBodyData {
	s.DepartmentId = &v
	return s
}

func (s *GetClientUserResponseBodyData) SetDescription(v string) *GetClientUserResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetClientUserResponseBodyData) SetEmail(v string) *GetClientUserResponseBodyData {
	s.Email = &v
	return s
}

func (s *GetClientUserResponseBodyData) SetId(v string) *GetClientUserResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetClientUserResponseBodyData) SetIdpConfigId(v string) *GetClientUserResponseBodyData {
	s.IdpConfigId = &v
	return s
}

func (s *GetClientUserResponseBodyData) SetMobileNumber(v string) *GetClientUserResponseBodyData {
	s.MobileNumber = &v
	return s
}

func (s *GetClientUserResponseBodyData) SetStatus(v string) *GetClientUserResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetClientUserResponseBodyData) SetUserId(v string) *GetClientUserResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetClientUserResponseBodyData) SetUsername(v string) *GetClientUserResponseBodyData {
	s.Username = &v
	return s
}

func (s *GetClientUserResponseBodyData) Validate() error {
	if s.Department != nil {
		if err := s.Department.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetClientUserResponseBodyDataDepartment struct {
	// example:
	//
	// 107
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetClientUserResponseBodyDataDepartment) String() string {
	return dara.Prettify(s)
}

func (s GetClientUserResponseBodyDataDepartment) GoString() string {
	return s.String()
}

func (s *GetClientUserResponseBodyDataDepartment) GetId() *string {
	return s.Id
}

func (s *GetClientUserResponseBodyDataDepartment) GetName() *string {
	return s.Name
}

func (s *GetClientUserResponseBodyDataDepartment) SetId(v string) *GetClientUserResponseBodyDataDepartment {
	s.Id = &v
	return s
}

func (s *GetClientUserResponseBodyDataDepartment) SetName(v string) *GetClientUserResponseBodyDataDepartment {
	s.Name = &v
	return s
}

func (s *GetClientUserResponseBodyDataDepartment) Validate() error {
	return dara.Validate(s)
}
