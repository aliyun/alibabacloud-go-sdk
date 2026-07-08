// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompanyThreeElementsVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CompanyThreeElementsVerificationResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CompanyThreeElementsVerificationResponseBody
	GetCode() *string
	SetData(v *CompanyThreeElementsVerificationResponseBodyData) *CompanyThreeElementsVerificationResponseBody
	GetData() *CompanyThreeElementsVerificationResponseBodyData
	SetMessage(v string) *CompanyThreeElementsVerificationResponseBody
	GetMessage() *string
	SetRequestId(v string) *CompanyThreeElementsVerificationResponseBody
	GetRequestId() *string
}

type CompanyThreeElementsVerificationResponseBody struct {
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
	// The returned data.
	Data *CompanyThreeElementsVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the returned status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The common parameter. Each request returns a unique ID, which can be used to troubleshoot and locate issues.
	//
	// example:
	//
	// 68A40250-50CD-034C-B728-0BD135850177
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CompanyThreeElementsVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CompanyThreeElementsVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *CompanyThreeElementsVerificationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CompanyThreeElementsVerificationResponseBody) GetCode() *string {
	return s.Code
}

func (s *CompanyThreeElementsVerificationResponseBody) GetData() *CompanyThreeElementsVerificationResponseBodyData {
	return s.Data
}

func (s *CompanyThreeElementsVerificationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CompanyThreeElementsVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CompanyThreeElementsVerificationResponseBody) SetAccessDeniedDetail(v string) *CompanyThreeElementsVerificationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CompanyThreeElementsVerificationResponseBody) SetCode(v string) *CompanyThreeElementsVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *CompanyThreeElementsVerificationResponseBody) SetData(v *CompanyThreeElementsVerificationResponseBodyData) *CompanyThreeElementsVerificationResponseBody {
	s.Data = v
	return s
}

func (s *CompanyThreeElementsVerificationResponseBody) SetMessage(v string) *CompanyThreeElementsVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *CompanyThreeElementsVerificationResponseBody) SetRequestId(v string) *CompanyThreeElementsVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompanyThreeElementsVerificationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CompanyThreeElementsVerificationResponseBodyData struct {
	// The company details.
	//
	// example:
	//
	// {
	//
	//       "enterpriseStatus": "在营（开业）",
	//
	//       "openTime": "2023-05-25/2053-05-24"
	//
	// }
	DetailInfo *CompanyThreeElementsVerificationResponseBodyDataDetailInfo `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty" type:"Struct"`
	// The fields whose verification results are inconsistent.
	InconsistentData []*string `json:"InconsistentData,omitempty" xml:"InconsistentData,omitempty" type:"Repeated"`
	// The verification result code. Valid values:
	//
	// - 0: The verification is consistent.
	//
	// - 1: The verification is consistent, but the company is not operating normally.
	//
	// - 2: The person-company verification is inconsistent.
	//
	// - 3: The two-element company verification failed.
	//
	// - 4: The company is not found.
	//
	// - 5: The person does not exist in the database.
	//
	// example:
	//
	// 2
	ReasonCode *int64 `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// The verification result. Valid values:
	//
	// - true: The information is consistent and the company is operating normally.
	//
	// - false: The verification failed.
	//
	// example:
	//
	// false
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s CompanyThreeElementsVerificationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CompanyThreeElementsVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CompanyThreeElementsVerificationResponseBodyData) GetDetailInfo() *CompanyThreeElementsVerificationResponseBodyDataDetailInfo {
	return s.DetailInfo
}

func (s *CompanyThreeElementsVerificationResponseBodyData) GetInconsistentData() []*string {
	return s.InconsistentData
}

func (s *CompanyThreeElementsVerificationResponseBodyData) GetReasonCode() *int64 {
	return s.ReasonCode
}

func (s *CompanyThreeElementsVerificationResponseBodyData) GetVerifyResult() *string {
	return s.VerifyResult
}

func (s *CompanyThreeElementsVerificationResponseBodyData) SetDetailInfo(v *CompanyThreeElementsVerificationResponseBodyDataDetailInfo) *CompanyThreeElementsVerificationResponseBodyData {
	s.DetailInfo = v
	return s
}

func (s *CompanyThreeElementsVerificationResponseBodyData) SetInconsistentData(v []*string) *CompanyThreeElementsVerificationResponseBodyData {
	s.InconsistentData = v
	return s
}

func (s *CompanyThreeElementsVerificationResponseBodyData) SetReasonCode(v int64) *CompanyThreeElementsVerificationResponseBodyData {
	s.ReasonCode = &v
	return s
}

func (s *CompanyThreeElementsVerificationResponseBodyData) SetVerifyResult(v string) *CompanyThreeElementsVerificationResponseBodyData {
	s.VerifyResult = &v
	return s
}

func (s *CompanyThreeElementsVerificationResponseBodyData) Validate() error {
	if s.DetailInfo != nil {
		if err := s.DetailInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CompanyThreeElementsVerificationResponseBodyDataDetailInfo struct {
	// The operating status of the company.
	//
	// example:
	//
	// 在营（开业）
	EnterpriseStatus *string `json:"EnterpriseStatus,omitempty" xml:"EnterpriseStatus,omitempty"`
	// The business term of the company.
	//
	// example:
	//
	// 2023-05-25/2053-05-24
	OpenTime *string `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
}

func (s CompanyThreeElementsVerificationResponseBodyDataDetailInfo) String() string {
	return dara.Prettify(s)
}

func (s CompanyThreeElementsVerificationResponseBodyDataDetailInfo) GoString() string {
	return s.String()
}

func (s *CompanyThreeElementsVerificationResponseBodyDataDetailInfo) GetEnterpriseStatus() *string {
	return s.EnterpriseStatus
}

func (s *CompanyThreeElementsVerificationResponseBodyDataDetailInfo) GetOpenTime() *string {
	return s.OpenTime
}

func (s *CompanyThreeElementsVerificationResponseBodyDataDetailInfo) SetEnterpriseStatus(v string) *CompanyThreeElementsVerificationResponseBodyDataDetailInfo {
	s.EnterpriseStatus = &v
	return s
}

func (s *CompanyThreeElementsVerificationResponseBodyDataDetailInfo) SetOpenTime(v string) *CompanyThreeElementsVerificationResponseBodyDataDetailInfo {
	s.OpenTime = &v
	return s
}

func (s *CompanyThreeElementsVerificationResponseBodyDataDetailInfo) Validate() error {
	return dara.Validate(s)
}
