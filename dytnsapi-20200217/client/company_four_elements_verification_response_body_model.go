// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompanyFourElementsVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CompanyFourElementsVerificationResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CompanyFourElementsVerificationResponseBody
	GetCode() *string
	SetData(v *CompanyFourElementsVerificationResponseBodyData) *CompanyFourElementsVerificationResponseBody
	GetData() *CompanyFourElementsVerificationResponseBodyData
	SetMessage(v string) *CompanyFourElementsVerificationResponseBody
	GetMessage() *string
	SetRequestId(v string) *CompanyFourElementsVerificationResponseBody
	GetRequestId() *string
}

type CompanyFourElementsVerificationResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// -
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The structure.
	Data *CompanyFourElementsVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the returned status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The common parameter. The ID returned for each request is unique and can be used to troubleshoot and locate issues.
	//
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CompanyFourElementsVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CompanyFourElementsVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *CompanyFourElementsVerificationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CompanyFourElementsVerificationResponseBody) GetCode() *string {
	return s.Code
}

func (s *CompanyFourElementsVerificationResponseBody) GetData() *CompanyFourElementsVerificationResponseBodyData {
	return s.Data
}

func (s *CompanyFourElementsVerificationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CompanyFourElementsVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CompanyFourElementsVerificationResponseBody) SetAccessDeniedDetail(v string) *CompanyFourElementsVerificationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CompanyFourElementsVerificationResponseBody) SetCode(v string) *CompanyFourElementsVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *CompanyFourElementsVerificationResponseBody) SetData(v *CompanyFourElementsVerificationResponseBodyData) *CompanyFourElementsVerificationResponseBody {
	s.Data = v
	return s
}

func (s *CompanyFourElementsVerificationResponseBody) SetMessage(v string) *CompanyFourElementsVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *CompanyFourElementsVerificationResponseBody) SetRequestId(v string) *CompanyFourElementsVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompanyFourElementsVerificationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CompanyFourElementsVerificationResponseBodyData struct {
	// The enterprise details.
	//
	// example:
	//
	// { "enterpriseStatus": "在营（开业）",
	//
	// "openTime": "2023-05-25/2053-05-24" }
	DetailInfo *CompanyFourElementsVerificationResponseBodyDataDetailInfo `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty" type:"Struct"`
	// The fields that failed verification.
	InconsistentData []*string `json:"InconsistentData,omitempty" xml:"InconsistentData,omitempty" type:"Repeated"`
	// The verification result code. Valid values:
	//
	// - 0: Verification passed.
	//
	// - 1: Verification passed, but the enterprise is not operating normally.
	//
	// - 2: The legal person and enterprise information are inconsistent.
	//
	// - 3: The enterprise four-element verification failed.
	//
	// - 4: The enterprise was not found.
	//
	// - 5: The legal person was not found in the database.
	//
	// example:
	//
	// 0
	ReasonCode *int64 `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// The verification result. Valid values:
	//
	// -   true: The information is consistent and the enterprise is operating normally.
	//
	// -   false: Verification failed.
	//
	// example:
	//
	// true
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s CompanyFourElementsVerificationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CompanyFourElementsVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CompanyFourElementsVerificationResponseBodyData) GetDetailInfo() *CompanyFourElementsVerificationResponseBodyDataDetailInfo {
	return s.DetailInfo
}

func (s *CompanyFourElementsVerificationResponseBodyData) GetInconsistentData() []*string {
	return s.InconsistentData
}

func (s *CompanyFourElementsVerificationResponseBodyData) GetReasonCode() *int64 {
	return s.ReasonCode
}

func (s *CompanyFourElementsVerificationResponseBodyData) GetVerifyResult() *string {
	return s.VerifyResult
}

func (s *CompanyFourElementsVerificationResponseBodyData) SetDetailInfo(v *CompanyFourElementsVerificationResponseBodyDataDetailInfo) *CompanyFourElementsVerificationResponseBodyData {
	s.DetailInfo = v
	return s
}

func (s *CompanyFourElementsVerificationResponseBodyData) SetInconsistentData(v []*string) *CompanyFourElementsVerificationResponseBodyData {
	s.InconsistentData = v
	return s
}

func (s *CompanyFourElementsVerificationResponseBodyData) SetReasonCode(v int64) *CompanyFourElementsVerificationResponseBodyData {
	s.ReasonCode = &v
	return s
}

func (s *CompanyFourElementsVerificationResponseBodyData) SetVerifyResult(v string) *CompanyFourElementsVerificationResponseBodyData {
	s.VerifyResult = &v
	return s
}

func (s *CompanyFourElementsVerificationResponseBodyData) Validate() error {
	if s.DetailInfo != nil {
		if err := s.DetailInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CompanyFourElementsVerificationResponseBodyDataDetailInfo struct {
	// The operating status of the enterprise.
	//
	// example:
	//
	// 在营（开业）
	EnterpriseStatus *string `json:"EnterpriseStatus,omitempty" xml:"EnterpriseStatus,omitempty"`
	// The business term of the enterprise.
	//
	// example:
	//
	// 2023-05-25/2053-05-24
	OpenTime *string `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
}

func (s CompanyFourElementsVerificationResponseBodyDataDetailInfo) String() string {
	return dara.Prettify(s)
}

func (s CompanyFourElementsVerificationResponseBodyDataDetailInfo) GoString() string {
	return s.String()
}

func (s *CompanyFourElementsVerificationResponseBodyDataDetailInfo) GetEnterpriseStatus() *string {
	return s.EnterpriseStatus
}

func (s *CompanyFourElementsVerificationResponseBodyDataDetailInfo) GetOpenTime() *string {
	return s.OpenTime
}

func (s *CompanyFourElementsVerificationResponseBodyDataDetailInfo) SetEnterpriseStatus(v string) *CompanyFourElementsVerificationResponseBodyDataDetailInfo {
	s.EnterpriseStatus = &v
	return s
}

func (s *CompanyFourElementsVerificationResponseBodyDataDetailInfo) SetOpenTime(v string) *CompanyFourElementsVerificationResponseBodyDataDetailInfo {
	s.OpenTime = &v
	return s
}

func (s *CompanyFourElementsVerificationResponseBodyDataDetailInfo) Validate() error {
	return dara.Validate(s)
}
