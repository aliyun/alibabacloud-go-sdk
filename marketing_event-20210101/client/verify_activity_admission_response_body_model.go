// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyActivityAdmissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VerifyActivityAdmissionResponseBody
	GetCode() *string
	SetData(v *VerifyActivityAdmissionResponseBodyData) *VerifyActivityAdmissionResponseBody
	GetData() *VerifyActivityAdmissionResponseBodyData
	SetMessage(v string) *VerifyActivityAdmissionResponseBody
	GetMessage() *string
	SetRequestId(v string) *VerifyActivityAdmissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *VerifyActivityAdmissionResponseBody
	GetSuccess() *bool
}

type VerifyActivityAdmissionResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *VerifyActivityAdmissionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s VerifyActivityAdmissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyActivityAdmissionResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyActivityAdmissionResponseBody) GetCode() *string {
	return s.Code
}

func (s *VerifyActivityAdmissionResponseBody) GetData() *VerifyActivityAdmissionResponseBodyData {
	return s.Data
}

func (s *VerifyActivityAdmissionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VerifyActivityAdmissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyActivityAdmissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VerifyActivityAdmissionResponseBody) SetCode(v string) *VerifyActivityAdmissionResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyActivityAdmissionResponseBody) SetData(v *VerifyActivityAdmissionResponseBodyData) *VerifyActivityAdmissionResponseBody {
	s.Data = v
	return s
}

func (s *VerifyActivityAdmissionResponseBody) SetMessage(v string) *VerifyActivityAdmissionResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyActivityAdmissionResponseBody) SetRequestId(v string) *VerifyActivityAdmissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyActivityAdmissionResponseBody) SetSuccess(v bool) *VerifyActivityAdmissionResponseBody {
	s.Success = &v
	return s
}

func (s *VerifyActivityAdmissionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VerifyActivityAdmissionResponseBodyData struct {
	ActivityId         *int64    `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	ActivityName       *string   `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	Address            *string   `json:"Address,omitempty" xml:"Address,omitempty"`
	AdmissionDate      *string   `json:"AdmissionDate,omitempty" xml:"AdmissionDate,omitempty"`
	AllowedDates       []*string `json:"AllowedDates,omitempty" xml:"AllowedDates,omitempty" type:"Repeated"`
	Decision           *string   `json:"Decision,omitempty" xml:"Decision,omitempty"`
	EndTime            *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MaskedIdCardNo     *string   `json:"MaskedIdCardNo,omitempty" xml:"MaskedIdCardNo,omitempty"`
	Name               *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	RegistrationStatus *int32    `json:"RegistrationStatus,omitempty" xml:"RegistrationStatus,omitempty"`
	SignedIn           *bool     `json:"SignedIn,omitempty" xml:"SignedIn,omitempty"`
	SigninTime         *string   `json:"SigninTime,omitempty" xml:"SigninTime,omitempty"`
	StartTime          *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s VerifyActivityAdmissionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s VerifyActivityAdmissionResponseBodyData) GoString() string {
	return s.String()
}

func (s *VerifyActivityAdmissionResponseBodyData) GetActivityId() *int64 {
	return s.ActivityId
}

func (s *VerifyActivityAdmissionResponseBodyData) GetActivityName() *string {
	return s.ActivityName
}

func (s *VerifyActivityAdmissionResponseBodyData) GetAddress() *string {
	return s.Address
}

func (s *VerifyActivityAdmissionResponseBodyData) GetAdmissionDate() *string {
	return s.AdmissionDate
}

func (s *VerifyActivityAdmissionResponseBodyData) GetAllowedDates() []*string {
	return s.AllowedDates
}

func (s *VerifyActivityAdmissionResponseBodyData) GetDecision() *string {
	return s.Decision
}

func (s *VerifyActivityAdmissionResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *VerifyActivityAdmissionResponseBodyData) GetMaskedIdCardNo() *string {
	return s.MaskedIdCardNo
}

func (s *VerifyActivityAdmissionResponseBodyData) GetName() *string {
	return s.Name
}

func (s *VerifyActivityAdmissionResponseBodyData) GetRegistrationStatus() *int32 {
	return s.RegistrationStatus
}

func (s *VerifyActivityAdmissionResponseBodyData) GetSignedIn() *bool {
	return s.SignedIn
}

func (s *VerifyActivityAdmissionResponseBodyData) GetSigninTime() *string {
	return s.SigninTime
}

func (s *VerifyActivityAdmissionResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *VerifyActivityAdmissionResponseBodyData) SetActivityId(v int64) *VerifyActivityAdmissionResponseBodyData {
	s.ActivityId = &v
	return s
}

func (s *VerifyActivityAdmissionResponseBodyData) SetActivityName(v string) *VerifyActivityAdmissionResponseBodyData {
	s.ActivityName = &v
	return s
}

func (s *VerifyActivityAdmissionResponseBodyData) SetAddress(v string) *VerifyActivityAdmissionResponseBodyData {
	s.Address = &v
	return s
}

func (s *VerifyActivityAdmissionResponseBodyData) SetAdmissionDate(v string) *VerifyActivityAdmissionResponseBodyData {
	s.AdmissionDate = &v
	return s
}

func (s *VerifyActivityAdmissionResponseBodyData) SetAllowedDates(v []*string) *VerifyActivityAdmissionResponseBodyData {
	s.AllowedDates = v
	return s
}

func (s *VerifyActivityAdmissionResponseBodyData) SetDecision(v string) *VerifyActivityAdmissionResponseBodyData {
	s.Decision = &v
	return s
}

func (s *VerifyActivityAdmissionResponseBodyData) SetEndTime(v string) *VerifyActivityAdmissionResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *VerifyActivityAdmissionResponseBodyData) SetMaskedIdCardNo(v string) *VerifyActivityAdmissionResponseBodyData {
	s.MaskedIdCardNo = &v
	return s
}

func (s *VerifyActivityAdmissionResponseBodyData) SetName(v string) *VerifyActivityAdmissionResponseBodyData {
	s.Name = &v
	return s
}

func (s *VerifyActivityAdmissionResponseBodyData) SetRegistrationStatus(v int32) *VerifyActivityAdmissionResponseBodyData {
	s.RegistrationStatus = &v
	return s
}

func (s *VerifyActivityAdmissionResponseBodyData) SetSignedIn(v bool) *VerifyActivityAdmissionResponseBodyData {
	s.SignedIn = &v
	return s
}

func (s *VerifyActivityAdmissionResponseBodyData) SetSigninTime(v string) *VerifyActivityAdmissionResponseBodyData {
	s.SigninTime = &v
	return s
}

func (s *VerifyActivityAdmissionResponseBodyData) SetStartTime(v string) *VerifyActivityAdmissionResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *VerifyActivityAdmissionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
