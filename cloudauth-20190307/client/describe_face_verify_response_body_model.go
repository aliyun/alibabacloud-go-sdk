// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFaceVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeFaceVerifyResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeFaceVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeFaceVerifyResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeFaceVerifyResponseBodyResultObject) *DescribeFaceVerifyResponseBody
	GetResultObject() *DescribeFaceVerifyResponseBodyResultObject
}

type DescribeFaceVerifyResponseBody struct {
	// Return code: 200 indicates success, other values indicate failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Error message
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 130A2C10-B9EE-4D84-88E3-5384FF039795
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result information
	ResultObject *DescribeFaceVerifyResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s DescribeFaceVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaceVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFaceVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeFaceVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeFaceVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFaceVerifyResponseBody) GetResultObject() *DescribeFaceVerifyResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeFaceVerifyResponseBody) SetCode(v string) *DescribeFaceVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeFaceVerifyResponseBody) SetMessage(v string) *DescribeFaceVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeFaceVerifyResponseBody) SetRequestId(v string) *DescribeFaceVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFaceVerifyResponseBody) SetResultObject(v *DescribeFaceVerifyResponseBodyResultObject) *DescribeFaceVerifyResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeFaceVerifyResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeFaceVerifyResponseBodyResultObject struct {
	// Device risk label.
	//
	// example:
	//
	// ROOT,VPN,HOOK
	DeviceRisk *string `json:"DeviceRisk,omitempty" xml:"DeviceRisk,omitempty"`
	// Device token.
	//
	// example:
	//
	// McozS1ZWRcRZStlERcZZo_QOytx5jcgZoZJEoRLOxxxxxxx
	DeviceToken *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	// Information about the authenticated subject, usually empty in general authentication scenarios.
	//
	// example:
	//
	// null
	IdentityInfo *string `json:"IdentityInfo,omitempty" xml:"IdentityInfo,omitempty"`
	// Attachment information of the authenticated subject, mainly image materials. JSON format, see example below.
	//
	// example:
	//
	// {"faceAttack": "F","facialPictureFront": {"qualityScore": 88.3615493774414,"pictureUrl": "https://cn-shanghai-aliyun-cloudauth-xxxxxx.oss-cn-shanghai.aliyuncs.com/verify/xxxxx/xxxxx.jpeg","ossBucketName": "cn-shanghai-aliyun-cloudauth-1260051251634779","ossObjectName": "verify/1260051251634779/6ba7bcfccf33f56cdb44ed086f36ce3e0.jpeg"}}
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty"`
	// Whether it passed, T for pass, F for fail.
	//
	// example:
	//
	// T
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// Description of the authentication result. For details, see the SubCode explanation below.
	//
	// example:
	//
	// 200
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	// Whether the response was successful.
	//
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// Records the identity information and corresponding encoding entered by the user under the rare character mode. The returned data is a JSON formatted string, which will be an empty string if there are no rare characters in the name.
	//
	// - name: Refers to the name entered by the user.
	//
	// - verifyName: Refers to the final name encoding after verification. For example, if a rare character is verified through transcoding: “Mr. Wang”, the actual verified name is “Wang Xiansheng”.
	//
	// - number: Refers to the identification number entered by the user.
	//
	// example:
	//
	// {
	//
	//   "number": "610***********1110",
	//
	//   "name": "王先生",
	//
	//   "verifyName": "王先"
	//
	// }
	UserInfo *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s DescribeFaceVerifyResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaceVerifyResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeFaceVerifyResponseBodyResultObject) GetDeviceRisk() *string {
	return s.DeviceRisk
}

func (s *DescribeFaceVerifyResponseBodyResultObject) GetDeviceToken() *string {
	return s.DeviceToken
}

func (s *DescribeFaceVerifyResponseBodyResultObject) GetIdentityInfo() *string {
	return s.IdentityInfo
}

func (s *DescribeFaceVerifyResponseBodyResultObject) GetMaterialInfo() *string {
	return s.MaterialInfo
}

func (s *DescribeFaceVerifyResponseBodyResultObject) GetPassed() *string {
	return s.Passed
}

func (s *DescribeFaceVerifyResponseBodyResultObject) GetSubCode() *string {
	return s.SubCode
}

func (s *DescribeFaceVerifyResponseBodyResultObject) GetSuccess() *string {
	return s.Success
}

func (s *DescribeFaceVerifyResponseBodyResultObject) GetUserInfo() *string {
	return s.UserInfo
}

func (s *DescribeFaceVerifyResponseBodyResultObject) SetDeviceRisk(v string) *DescribeFaceVerifyResponseBodyResultObject {
	s.DeviceRisk = &v
	return s
}

func (s *DescribeFaceVerifyResponseBodyResultObject) SetDeviceToken(v string) *DescribeFaceVerifyResponseBodyResultObject {
	s.DeviceToken = &v
	return s
}

func (s *DescribeFaceVerifyResponseBodyResultObject) SetIdentityInfo(v string) *DescribeFaceVerifyResponseBodyResultObject {
	s.IdentityInfo = &v
	return s
}

func (s *DescribeFaceVerifyResponseBodyResultObject) SetMaterialInfo(v string) *DescribeFaceVerifyResponseBodyResultObject {
	s.MaterialInfo = &v
	return s
}

func (s *DescribeFaceVerifyResponseBodyResultObject) SetPassed(v string) *DescribeFaceVerifyResponseBodyResultObject {
	s.Passed = &v
	return s
}

func (s *DescribeFaceVerifyResponseBodyResultObject) SetSubCode(v string) *DescribeFaceVerifyResponseBodyResultObject {
	s.SubCode = &v
	return s
}

func (s *DescribeFaceVerifyResponseBodyResultObject) SetSuccess(v string) *DescribeFaceVerifyResponseBodyResultObject {
	s.Success = &v
	return s
}

func (s *DescribeFaceVerifyResponseBodyResultObject) SetUserInfo(v string) *DescribeFaceVerifyResponseBodyResultObject {
	s.UserInfo = &v
	return s
}

func (s *DescribeFaceVerifyResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
