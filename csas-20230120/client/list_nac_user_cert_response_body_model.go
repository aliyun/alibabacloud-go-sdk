// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNacUserCertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *ListNacUserCertResponseBody
	GetCode() *int64
	SetDataList(v []*ListNacUserCertResponseBodyDataList) *ListNacUserCertResponseBody
	GetDataList() []*ListNacUserCertResponseBodyDataList
	SetMessage(v string) *ListNacUserCertResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListNacUserCertResponseBody
	GetRequestId() *string
	SetTotalNum(v int64) *ListNacUserCertResponseBody
	GetTotalNum() *int64
}

type ListNacUserCertResponseBody struct {
	// example:
	//
	// 200
	Code     *int64                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	DataList []*ListNacUserCertResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListNacUserCertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNacUserCertResponseBody) GoString() string {
	return s.String()
}

func (s *ListNacUserCertResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ListNacUserCertResponseBody) GetDataList() []*ListNacUserCertResponseBodyDataList {
	return s.DataList
}

func (s *ListNacUserCertResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListNacUserCertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNacUserCertResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *ListNacUserCertResponseBody) SetCode(v int64) *ListNacUserCertResponseBody {
	s.Code = &v
	return s
}

func (s *ListNacUserCertResponseBody) SetDataList(v []*ListNacUserCertResponseBodyDataList) *ListNacUserCertResponseBody {
	s.DataList = v
	return s
}

func (s *ListNacUserCertResponseBody) SetMessage(v string) *ListNacUserCertResponseBody {
	s.Message = &v
	return s
}

func (s *ListNacUserCertResponseBody) SetRequestId(v string) *ListNacUserCertResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNacUserCertResponseBody) SetTotalNum(v int64) *ListNacUserCertResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListNacUserCertResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListNacUserCertResponseBodyDataList struct {
	// example:
	//
	// 1
	Aliuid     *string `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DevTag *string `json:"DevTag,omitempty" xml:"DevTag,omitempty"`
	// example:
	//
	// windows
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// 2029-06-30 09:31:54
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// example:
	//
	// MS-XU****
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// 08:f8:**:**:**:5e
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// example:
	//
	// Disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// zhang**
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListNacUserCertResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListNacUserCertResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListNacUserCertResponseBodyDataList) GetAliuid() *string {
	return s.Aliuid
}

func (s *ListNacUserCertResponseBodyDataList) GetDepartment() *string {
	return s.Department
}

func (s *ListNacUserCertResponseBodyDataList) GetDevTag() *string {
	return s.DevTag
}

func (s *ListNacUserCertResponseBodyDataList) GetDeviceType() *string {
	return s.DeviceType
}

func (s *ListNacUserCertResponseBodyDataList) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *ListNacUserCertResponseBodyDataList) GetHostname() *string {
	return s.Hostname
}

func (s *ListNacUserCertResponseBodyDataList) GetMac() *string {
	return s.Mac
}

func (s *ListNacUserCertResponseBodyDataList) GetStatus() *string {
	return s.Status
}

func (s *ListNacUserCertResponseBodyDataList) GetUserId() *string {
	return s.UserId
}

func (s *ListNacUserCertResponseBodyDataList) GetUsername() *string {
	return s.Username
}

func (s *ListNacUserCertResponseBodyDataList) SetAliuid(v string) *ListNacUserCertResponseBodyDataList {
	s.Aliuid = &v
	return s
}

func (s *ListNacUserCertResponseBodyDataList) SetDepartment(v string) *ListNacUserCertResponseBodyDataList {
	s.Department = &v
	return s
}

func (s *ListNacUserCertResponseBodyDataList) SetDevTag(v string) *ListNacUserCertResponseBodyDataList {
	s.DevTag = &v
	return s
}

func (s *ListNacUserCertResponseBodyDataList) SetDeviceType(v string) *ListNacUserCertResponseBodyDataList {
	s.DeviceType = &v
	return s
}

func (s *ListNacUserCertResponseBodyDataList) SetExpiredTime(v string) *ListNacUserCertResponseBodyDataList {
	s.ExpiredTime = &v
	return s
}

func (s *ListNacUserCertResponseBodyDataList) SetHostname(v string) *ListNacUserCertResponseBodyDataList {
	s.Hostname = &v
	return s
}

func (s *ListNacUserCertResponseBodyDataList) SetMac(v string) *ListNacUserCertResponseBodyDataList {
	s.Mac = &v
	return s
}

func (s *ListNacUserCertResponseBodyDataList) SetStatus(v string) *ListNacUserCertResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListNacUserCertResponseBodyDataList) SetUserId(v string) *ListNacUserCertResponseBodyDataList {
	s.UserId = &v
	return s
}

func (s *ListNacUserCertResponseBodyDataList) SetUsername(v string) *ListNacUserCertResponseBodyDataList {
	s.Username = &v
	return s
}

func (s *ListNacUserCertResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
