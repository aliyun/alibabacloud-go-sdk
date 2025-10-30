// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySecretAPhoneNoToCustResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QuerySecretAPhoneNoToCustResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QuerySecretAPhoneNoToCustResponseBody
	GetCode() *string
	SetData(v *QuerySecretAPhoneNoToCustResponseBodyData) *QuerySecretAPhoneNoToCustResponseBody
	GetData() *QuerySecretAPhoneNoToCustResponseBodyData
	SetMessage(v string) *QuerySecretAPhoneNoToCustResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySecretAPhoneNoToCustResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySecretAPhoneNoToCustResponseBody
	GetSuccess() *bool
}

type QuerySecretAPhoneNoToCustResponseBody struct {
	// example:
	//
	// 0
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 请求状态码
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// A号码报备状态查询结构体
	Data *QuerySecretAPhoneNoToCustResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 失败错误提示
	//
	// example:
	//
	// 号码组不存在
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 返回id
	//
	// example:
	//
	// 1D73E648-0978-18A5-B089-3BB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求是否成功
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySecretAPhoneNoToCustResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySecretAPhoneNoToCustResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySecretAPhoneNoToCustResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QuerySecretAPhoneNoToCustResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySecretAPhoneNoToCustResponseBody) GetData() *QuerySecretAPhoneNoToCustResponseBodyData {
	return s.Data
}

func (s *QuerySecretAPhoneNoToCustResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySecretAPhoneNoToCustResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySecretAPhoneNoToCustResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySecretAPhoneNoToCustResponseBody) SetAccessDeniedDetail(v string) *QuerySecretAPhoneNoToCustResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponseBody) SetCode(v string) *QuerySecretAPhoneNoToCustResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponseBody) SetData(v *QuerySecretAPhoneNoToCustResponseBodyData) *QuerySecretAPhoneNoToCustResponseBody {
	s.Data = v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponseBody) SetMessage(v string) *QuerySecretAPhoneNoToCustResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponseBody) SetRequestId(v string) *QuerySecretAPhoneNoToCustResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponseBody) SetSuccess(v bool) *QuerySecretAPhoneNoToCustResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySecretAPhoneNoToCustResponseBodyData struct {
	// 所属a号码组id
	//
	// example:
	//
	// 19
	ANoWhiteGroupId *string `json:"ANoWhiteGroupId,omitempty" xml:"ANoWhiteGroupId,omitempty"`
	// 固话报备的经办人/法人电话
	//
	// example:
	//
	// 130*****8888
	CustPhoneNo *string `json:"CustPhoneNo,omitempty" xml:"CustPhoneNo,omitempty"`
	// 号码类型
	//
	// example:
	//
	// Mobile
	NoType *string `json:"NoType,omitempty" xml:"NoType,omitempty"`
	// A号码
	//
	// example:
	//
	// 130*****1234
	PhoneNoA *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	// 备注（客户自己的业务备注，可编辑）
	//
	// example:
	//
	// ***报备
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 报备失败原因
	//
	// example:
	//
	// ["系统判断为不同人"]
	ReportResult *string `json:"ReportResult,omitempty" xml:"ReportResult,omitempty"`
	// 报备结果
	//
	// example:
	//
	// REVIEWING
	ReportStatus *string `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
}

func (s QuerySecretAPhoneNoToCustResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QuerySecretAPhoneNoToCustResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySecretAPhoneNoToCustResponseBodyData) GetANoWhiteGroupId() *string {
	return s.ANoWhiteGroupId
}

func (s *QuerySecretAPhoneNoToCustResponseBodyData) GetCustPhoneNo() *string {
	return s.CustPhoneNo
}

func (s *QuerySecretAPhoneNoToCustResponseBodyData) GetNoType() *string {
	return s.NoType
}

func (s *QuerySecretAPhoneNoToCustResponseBodyData) GetPhoneNoA() *string {
	return s.PhoneNoA
}

func (s *QuerySecretAPhoneNoToCustResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *QuerySecretAPhoneNoToCustResponseBodyData) GetReportResult() *string {
	return s.ReportResult
}

func (s *QuerySecretAPhoneNoToCustResponseBodyData) GetReportStatus() *string {
	return s.ReportStatus
}

func (s *QuerySecretAPhoneNoToCustResponseBodyData) SetANoWhiteGroupId(v string) *QuerySecretAPhoneNoToCustResponseBodyData {
	s.ANoWhiteGroupId = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponseBodyData) SetCustPhoneNo(v string) *QuerySecretAPhoneNoToCustResponseBodyData {
	s.CustPhoneNo = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponseBodyData) SetNoType(v string) *QuerySecretAPhoneNoToCustResponseBodyData {
	s.NoType = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponseBodyData) SetPhoneNoA(v string) *QuerySecretAPhoneNoToCustResponseBodyData {
	s.PhoneNoA = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponseBodyData) SetRemark(v string) *QuerySecretAPhoneNoToCustResponseBodyData {
	s.Remark = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponseBodyData) SetReportResult(v string) *QuerySecretAPhoneNoToCustResponseBodyData {
	s.ReportResult = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponseBodyData) SetReportStatus(v string) *QuerySecretAPhoneNoToCustResponseBodyData {
	s.ReportStatus = &v
	return s
}

func (s *QuerySecretAPhoneNoToCustResponseBodyData) Validate() error {
	return dara.Validate(s)
}
