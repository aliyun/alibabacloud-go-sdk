// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompanyTwoElementsVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CompanyTwoElementsVerificationResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CompanyTwoElementsVerificationResponseBody
	GetCode() *string
	SetData(v *CompanyTwoElementsVerificationResponseBodyData) *CompanyTwoElementsVerificationResponseBody
	GetData() *CompanyTwoElementsVerificationResponseBodyData
	SetMessage(v string) *CompanyTwoElementsVerificationResponseBody
	GetMessage() *string
	SetRequestId(v string) *CompanyTwoElementsVerificationResponseBody
	GetRequestId() *string
}

type CompanyTwoElementsVerificationResponseBody struct {
	// Details about the access denial.
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
	// The struct.
	Data *CompanyTwoElementsVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CompanyTwoElementsVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CompanyTwoElementsVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *CompanyTwoElementsVerificationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CompanyTwoElementsVerificationResponseBody) GetCode() *string {
	return s.Code
}

func (s *CompanyTwoElementsVerificationResponseBody) GetData() *CompanyTwoElementsVerificationResponseBodyData {
	return s.Data
}

func (s *CompanyTwoElementsVerificationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CompanyTwoElementsVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CompanyTwoElementsVerificationResponseBody) SetAccessDeniedDetail(v string) *CompanyTwoElementsVerificationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CompanyTwoElementsVerificationResponseBody) SetCode(v string) *CompanyTwoElementsVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *CompanyTwoElementsVerificationResponseBody) SetData(v *CompanyTwoElementsVerificationResponseBodyData) *CompanyTwoElementsVerificationResponseBody {
	s.Data = v
	return s
}

func (s *CompanyTwoElementsVerificationResponseBody) SetMessage(v string) *CompanyTwoElementsVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *CompanyTwoElementsVerificationResponseBody) SetRequestId(v string) *CompanyTwoElementsVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompanyTwoElementsVerificationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CompanyTwoElementsVerificationResponseBodyData struct {
	// Enterprise details.
	//
	// > This field is not returned for public institutions or civil groups.
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
	DetailInfo *CompanyTwoElementsVerificationResponseBodyDataDetailInfo `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty" type:"Struct"`
	// The fields that are inconsistent in the verification.
	InconsistentData []*string `json:"InconsistentData,omitempty" xml:"InconsistentData,omitempty" type:"Repeated"`
	// The verification result code. Valid values:
	//
	// - 0: Verification consistent
	//
	// - 1: Verification consistent, the enterprise is not in normal operation
	//
	// - 3: The two enterprise elements failed verification
	//
	// - 4: No such enterprise found
	//
	// example:
	//
	// 0
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// The verification result. Valid values:
	//
	// - true: The information is verified to be consistent, and the enterprise is operating normally.
	//
	// - false: The verification failed.
	//
	// > The operating status verification is not supported for public institutions or civil groups. Only the consistency between the enterprise name and the enterprise certificate number is verified.
	//
	// example:
	//
	// true
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s CompanyTwoElementsVerificationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CompanyTwoElementsVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CompanyTwoElementsVerificationResponseBodyData) GetDetailInfo() *CompanyTwoElementsVerificationResponseBodyDataDetailInfo {
	return s.DetailInfo
}

func (s *CompanyTwoElementsVerificationResponseBodyData) GetInconsistentData() []*string {
	return s.InconsistentData
}

func (s *CompanyTwoElementsVerificationResponseBodyData) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *CompanyTwoElementsVerificationResponseBodyData) GetVerifyResult() *string {
	return s.VerifyResult
}

func (s *CompanyTwoElementsVerificationResponseBodyData) SetDetailInfo(v *CompanyTwoElementsVerificationResponseBodyDataDetailInfo) *CompanyTwoElementsVerificationResponseBodyData {
	s.DetailInfo = v
	return s
}

func (s *CompanyTwoElementsVerificationResponseBodyData) SetInconsistentData(v []*string) *CompanyTwoElementsVerificationResponseBodyData {
	s.InconsistentData = v
	return s
}

func (s *CompanyTwoElementsVerificationResponseBodyData) SetReasonCode(v string) *CompanyTwoElementsVerificationResponseBodyData {
	s.ReasonCode = &v
	return s
}

func (s *CompanyTwoElementsVerificationResponseBodyData) SetVerifyResult(v string) *CompanyTwoElementsVerificationResponseBodyData {
	s.VerifyResult = &v
	return s
}

func (s *CompanyTwoElementsVerificationResponseBodyData) Validate() error {
	if s.DetailInfo != nil {
		if err := s.DetailInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CompanyTwoElementsVerificationResponseBodyDataDetailInfo struct {
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

func (s CompanyTwoElementsVerificationResponseBodyDataDetailInfo) String() string {
	return dara.Prettify(s)
}

func (s CompanyTwoElementsVerificationResponseBodyDataDetailInfo) GoString() string {
	return s.String()
}

func (s *CompanyTwoElementsVerificationResponseBodyDataDetailInfo) GetEnterpriseStatus() *string {
	return s.EnterpriseStatus
}

func (s *CompanyTwoElementsVerificationResponseBodyDataDetailInfo) GetOpenTime() *string {
	return s.OpenTime
}

func (s *CompanyTwoElementsVerificationResponseBodyDataDetailInfo) SetEnterpriseStatus(v string) *CompanyTwoElementsVerificationResponseBodyDataDetailInfo {
	s.EnterpriseStatus = &v
	return s
}

func (s *CompanyTwoElementsVerificationResponseBodyDataDetailInfo) SetOpenTime(v string) *CompanyTwoElementsVerificationResponseBodyDataDetailInfo {
	s.OpenTime = &v
	return s
}

func (s *CompanyTwoElementsVerificationResponseBodyDataDetailInfo) Validate() error {
	return dara.Validate(s)
}
