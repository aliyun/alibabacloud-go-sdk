// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDigitalSignByNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryDigitalSignByNameResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryDigitalSignByNameResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *QueryDigitalSignByNameResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *QueryDigitalSignByNameResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryDigitalSignByNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryDigitalSignByNameResponseBody
	GetSuccess() *bool
}

type QueryDigitalSignByNameResponseBody struct {
	// Details of the access denial. This parameter is returned only if Resource Access Management (RAM) authentication fails.
	//
	// example:
	//
	// 无
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code of the request. A value of `OK` indicates that the request was successful. Other values indicate error codes.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// - `gmtModified`: The time when the signature was last modified.
	//
	// - `creator`: The ID of the user who created the signature.
	//
	// - `signName`: The name of the digital SMS signature.
	//
	// - `qualificationId`: The ID of the qualification. You can create qualifications in the console or by calling an API operation.
	//
	// - `signIndustry`: The industry type. Valid values: `0` (General) and `1` (E-commerce).
	//
	// - `signVersion`: The version of the signature. This value is updated for each new version. The current version is 1.
	//
	// - `telecomRegisterStatus`: The filing status with China Telecom. Valid values: `0` (Filing Failed), `3` (Filing Successful), `-1` (Filing in Progress), and `-2` (Not Filed).
	//
	// - `signCode`: The code of the digital SMS signature.
	//
	// - `gmtCreate`: The time when the signature was created.
	//
	// - `signId`: The ID of the signature. This is a unique identifier.
	//
	// - `mobileRegisterStatus`: The filing status with China Mobile.
	//
	// - `SignSource`: The source of the signature. Valid values:
	//
	// - `mobileAvailableStatus`: The availability status with China Mobile. Valid values: `0` (Unavailable) and `1` (Available). We recommend that you select an available signature when you create a template or send a digital SMS message.
	//
	// - `unicomRegisterStatus`: The filing status with China Unicom. Valid values: `0` (Filing Failed), `3` (Filing Successful), `-1` (Filing in Progress), and `-2` (Not Filed).
	//
	// - `unicomAvailableStatus`: The availability status with China Unicom. Valid values: `0` (Unavailable) and `1` (Available). We recommend that you select an available signature when you create a template or send a digital SMS message.
	//
	// - `telecomAvailableStatus`: The availability status with China Telecom. Valid values: `0` (Unavailable) and `1` (Available). We recommend that you select an available signature when you create a template or send a digital SMS message.
	//
	// example:
	//
	// "Data": {
	//
	//     "gmtModified": 1770005896000,
	//
	//     "creator": "UID:1498275912899720",
	//
	//     "signName": "驳回理由",
	//
	//     "qualificationId": 2757722,
	//
	//     "qualificationVersion": 1768974751432,
	//
	//     "signIndustry": 0,
	//
	//     "registerStatueReason": {},
	//
	//     "signVersion": 1,
	//
	//     "remark": "",
	//
	//     "telecomRegisterStatus": -1,
	//
	//     "signCode": "DIGSIGN_100000168688001_1769050485148_pYZu1",
	//
	//     "gmtCreate": 1769050485000,
	//
	//     "signId": 22784769,
	//
	//     "mobileRegisterStatus": -1,
	//
	//     "signSource": 0,
	//
	//     "mobileAvailableStatus": 0,
	//
	//     "attachmentList": [],
	//
	//     "unicomRegisterStatus": 3,
	//
	//     "unicomAvailableStatus": 1,
	//
	//     "signType": 1,
	//
	//     "telecomAvailableStatus": 0,
	//
	//     "extendMessage": "{}"
	//
	//   },
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDigitalSignByNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDigitalSignByNameResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDigitalSignByNameResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryDigitalSignByNameResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryDigitalSignByNameResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *QueryDigitalSignByNameResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryDigitalSignByNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDigitalSignByNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryDigitalSignByNameResponseBody) SetAccessDeniedDetail(v string) *QueryDigitalSignByNameResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryDigitalSignByNameResponseBody) SetCode(v string) *QueryDigitalSignByNameResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDigitalSignByNameResponseBody) SetData(v map[string]interface{}) *QueryDigitalSignByNameResponseBody {
	s.Data = v
	return s
}

func (s *QueryDigitalSignByNameResponseBody) SetMessage(v string) *QueryDigitalSignByNameResponseBody {
	s.Message = &v
	return s
}

func (s *QueryDigitalSignByNameResponseBody) SetRequestId(v string) *QueryDigitalSignByNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDigitalSignByNameResponseBody) SetSuccess(v bool) *QueryDigitalSignByNameResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDigitalSignByNameResponseBody) Validate() error {
	return dara.Validate(s)
}
