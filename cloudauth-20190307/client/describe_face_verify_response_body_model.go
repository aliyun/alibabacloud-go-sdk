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
	// The return code. A value of 200 indicates success. Other values indicate failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 130A2C10-B9EE-4D84-88E3-5384FF039795
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result information.
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
	// The device risk label.
	//
	// example:
	//
	// ROOT,VPN,HOOK
	DeviceRisk *string `json:"DeviceRisk,omitempty" xml:"DeviceRisk,omitempty"`
	// The device token.
	//
	// example:
	//
	// McozS1ZWRcRZStlERcZZo_QOytx5jcgZoZJEoRLOxxxxxxx
	DeviceToken *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	// The identity information of the verification subject. This field is empty in common verification scenarios.
	//
	// example:
	//
	// null
	IdentityInfo *string `json:"IdentityInfo,omitempty" xml:"IdentityInfo,omitempty"`
	// The supplementary material information of the verification subject, primarily image-based materials. The value is in JSON format. See the example below.
	//
	// example:
	//
	// {"faceAttack": "F","facialPictureFront": {"qualityScore": 88.3615493774414,"pictureUrl": "https://cn-shanghai-aliyun-cloudauth-xxxxxx.oss-cn-shanghai.aliyuncs.com/verify/xxxxx/xxxxx.jpeg","ossBucketName": "cn-shanghai-aliyun-cloudauth-1260051251634779","ossObjectName": "verify/1260051251634779/6ba7bcfccf33f56cdb44ed086f36ce3e0.jpeg"}}
	MaterialInfo *string `json:"MaterialInfo,omitempty" xml:"MaterialInfo,omitempty"`
	// Indicates whether the verification is passed. A value of T indicates passed. A value of F indicates not passed.
	//
	// example:
	//
	// T
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// The verification result description. For more information, see the SubCode description below.
	//
	// example:
	//
	// 200
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	// Indicates whether the response is successful.
	//
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The identity information and corresponding encoding entered by the user in rare character mode. The value is a JSON-formatted string. An empty string is returned if the name does not contain rare characters.
	//
	// - name: the name entered by the user.
	//
	// - verifyName: the final name encoding that passed verification. For example, if a rare character is verified through transcoding: "王先生", the actual verified name is "王先升".
	//
	// - number: the ID number entered by the user.
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
