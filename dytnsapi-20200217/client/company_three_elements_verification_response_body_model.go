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
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *CompanyThreeElementsVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID. It is a common parameter and can be used to troubleshoot issues.
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
	// The information about the enterprise.
	DetailInfo *CompanyThreeElementsVerificationResponseBodyDataDetailInfo `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty" type:"Struct"`
	// The fields to be verified.
	InconsistentData []*string `json:"InconsistentData,omitempty" xml:"InconsistentData,omitempty" type:"Repeated"`
	// The code of the verification result. Valid values:
	//
	// 	- 0: The three elements belong to the same enterprise.
	//
	// 	- 1: The three elements belong to the same enterprise, and the business status of the enterprise is abnormal.
	//
	// 	- 2: The legal representative information cannot match the enterprise information.
	//
	// 	- 3: The three elements do not belong to the same enterprise.
	//
	// 	- 4: No information about the enterprise is found.
	//
	// 	- 5: No information about the legal representative is found.
	//
	// example:
	//
	// 0
	ReasonCode *int64 `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// The verification result. Valid values:
	//
	// 	- true: The three elements belong to the same enterprise and the business status of the enterprise is Active.
	//
	// 	- false: The three elements do not belong to the same enterprise.
	//
	// example:
	//
	// true
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
	// The business status of the enterprise.
	//
	// example:
	//
	// Active
	EnterpriseStatus *string `json:"EnterpriseStatus,omitempty" xml:"EnterpriseStatus,omitempty"`
	// The business term of the enterprise.
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
