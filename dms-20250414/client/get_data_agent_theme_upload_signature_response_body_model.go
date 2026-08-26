// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataAgentThemeUploadSignatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDataAgentThemeUploadSignatureResponseBodyData) *GetDataAgentThemeUploadSignatureResponseBody
	GetData() *GetDataAgentThemeUploadSignatureResponseBodyData
	SetErrorCode(v string) *GetDataAgentThemeUploadSignatureResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataAgentThemeUploadSignatureResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetDataAgentThemeUploadSignatureResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataAgentThemeUploadSignatureResponseBody
	GetSuccess() *bool
}

type GetDataAgentThemeUploadSignatureResponseBody struct {
	// The response struct.
	Data *GetDataAgentThemeUploadSignatureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned when the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the call fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID, which is used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataAgentThemeUploadSignatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataAgentThemeUploadSignatureResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataAgentThemeUploadSignatureResponseBody) GetData() *GetDataAgentThemeUploadSignatureResponseBodyData {
	return s.Data
}

func (s *GetDataAgentThemeUploadSignatureResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataAgentThemeUploadSignatureResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataAgentThemeUploadSignatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataAgentThemeUploadSignatureResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataAgentThemeUploadSignatureResponseBody) SetData(v *GetDataAgentThemeUploadSignatureResponseBodyData) *GetDataAgentThemeUploadSignatureResponseBody {
	s.Data = v
	return s
}

func (s *GetDataAgentThemeUploadSignatureResponseBody) SetErrorCode(v string) *GetDataAgentThemeUploadSignatureResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataAgentThemeUploadSignatureResponseBody) SetErrorMessage(v string) *GetDataAgentThemeUploadSignatureResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataAgentThemeUploadSignatureResponseBody) SetRequestId(v string) *GetDataAgentThemeUploadSignatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataAgentThemeUploadSignatureResponseBody) SetSuccess(v bool) *GetDataAgentThemeUploadSignatureResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataAgentThemeUploadSignatureResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataAgentThemeUploadSignatureResponseBodyData struct {
	// The policy expiration time in ISO 8601 UTC format.
	//
	// example:
	//
	// 2026-01-01T14:53:41Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The target object key, which is exactly locked by the policy.
	ObjectKey *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	// The parameter set that specifies the derived key.
	OssCredential *string `json:"OssCredential,omitempty" xml:"OssCredential,omitempty"`
	// The signature time in the format of yyyyMMdd\\"T\\"HHmmss\\"Z\\".
	//
	// example:
	//
	// 20260101T135341Z
	OssDate *string `json:"OssDate,omitempty" xml:"OssDate,omitempty"`
	// The STS token used to upload files to OSS. The token is valid for 1 hour.
	//
	// example:
	//
	// CAIS4gJ1q6Ft5B2yfSjIr5vPHMj4p+lHx/utUUjg13ptZ+5u3oDzkzz2IHhMdXlrCOgYt/8xnG1V6f8flrJ/ToQAX0HfatZq5ZkS9AqnaoXM/te496IFg5D9y7dIs8GgjqHoeOzcYI73WJXEMiLp9EJaxb/9ak/RPTiMOoGIjphKd8keWhLCAxNNGNZRIHkJyqZYTwyzU8ygKRn3mGHdIVN1sw5n8wNF5L+439eX52i17jS46JdM/9ysesH5NpQxbMwkDYnk5oEsKPqdihw3wgNR6aJ7gJZD/Tr6pdyHCzFTmU7ea7uEqYw3clYiOPBnRvEd8eKPnPl5q/HVm2Hs0wxKNuxOSCXZS4yp3MLeH+ekJgOGwWFHz9qnOLmtQXqV22tMCRpzXIj6Zlmz+/reI6iNW+Ory74mxSFbrz3ZP4yv+o+Yv3QbMVumcySkKVbBbVvnv0R8GNsIC2lMUbp+rfShhfFuG2QagAECCyigwAlSAryrFmteD+EVuvxvi0NE7zDJLbUkhek6dcY+/u5V5jcmvL67cQ7bTNk+9lV8WDCvtoCD9ucqTaHweJEd8fS2DaFedAMDf8BfZa2C1CTLhVXdSgE2WORYbMqidelRm7dH3fTbZVvryWKDaveDRLt5J/Qfs**********
	OssSecurityToken *string `json:"OssSecurityToken,omitempty" xml:"OssSecurityToken,omitempty"`
	// The V4 signature value.
	//
	// example:
	//
	// 9bebe0900716bdefaab899781c7bdfd614ec6ed711e0de5ddf6f5a**********
	OssSignature *string `json:"OssSignature,omitempty" xml:"OssSignature,omitempty"`
	// The signature version. The value is fixed as OSS4-HMAC-SHA256.
	//
	// example:
	//
	// OSS4-HMAC-SHA256
	OssSignatureVersion *string `json:"OssSignatureVersion,omitempty" xml:"OssSignatureVersion,omitempty"`
	// The Base64-encoded value of the policy JSON.
	//
	// example:
	//
	// eyJjb25kaXRpb25zIjpbeyJ4LW9zcy1jcmVkZW50aWFsIjoiU1RTLk5YeldyTEo2ZnA5RlNuUTN6OGthQjFFWH**********
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The theme business identifier generated or reused for this request. Pass this identifier to the CreateDataAgentTheme operation after the upload is complete to register the metadata.
	//
	// example:
	//
	// 0f8b2c1d****a3e5f7b1c2d
	ThemeId *string `json:"ThemeId,omitempty" xml:"ThemeId,omitempty"`
	// The upload directory prefix.
	UploadDir *string `json:"UploadDir,omitempty" xml:"UploadDir,omitempty"`
	// The PostObject destination address over the public network.
	//
	// example:
	//
	// https://**********.oss-cn-hangzhou.aliyuncs.com
	UploadHost *string `json:"UploadHost,omitempty" xml:"UploadHost,omitempty"`
}

func (s GetDataAgentThemeUploadSignatureResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataAgentThemeUploadSignatureResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataAgentThemeUploadSignatureResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GetDataAgentThemeUploadSignatureResponseBodyData) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *GetDataAgentThemeUploadSignatureResponseBodyData) GetOssCredential() *string {
	return s.OssCredential
}

func (s *GetDataAgentThemeUploadSignatureResponseBodyData) GetOssDate() *string {
	return s.OssDate
}

func (s *GetDataAgentThemeUploadSignatureResponseBodyData) GetOssSecurityToken() *string {
	return s.OssSecurityToken
}

func (s *GetDataAgentThemeUploadSignatureResponseBodyData) GetOssSignature() *string {
	return s.OssSignature
}

func (s *GetDataAgentThemeUploadSignatureResponseBodyData) GetOssSignatureVersion() *string {
	return s.OssSignatureVersion
}

func (s *GetDataAgentThemeUploadSignatureResponseBodyData) GetPolicy() *string {
	return s.Policy
}

func (s *GetDataAgentThemeUploadSignatureResponseBodyData) GetThemeId() *string {
	return s.ThemeId
}

func (s *GetDataAgentThemeUploadSignatureResponseBodyData) GetUploadDir() *string {
	return s.UploadDir
}

func (s *GetDataAgentThemeUploadSignatureResponseBodyData) GetUploadHost() *string {
	return s.UploadHost
}

func (s *GetDataAgentThemeUploadSignatureResponseBodyData) SetExpireTime(v string) *GetDataAgentThemeUploadSignatureResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetDataAgentThemeUploadSignatureResponseBodyData) SetObjectKey(v string) *GetDataAgentThemeUploadSignatureResponseBodyData {
	s.ObjectKey = &v
	return s
}

func (s *GetDataAgentThemeUploadSignatureResponseBodyData) SetOssCredential(v string) *GetDataAgentThemeUploadSignatureResponseBodyData {
	s.OssCredential = &v
	return s
}

func (s *GetDataAgentThemeUploadSignatureResponseBodyData) SetOssDate(v string) *GetDataAgentThemeUploadSignatureResponseBodyData {
	s.OssDate = &v
	return s
}

func (s *GetDataAgentThemeUploadSignatureResponseBodyData) SetOssSecurityToken(v string) *GetDataAgentThemeUploadSignatureResponseBodyData {
	s.OssSecurityToken = &v
	return s
}

func (s *GetDataAgentThemeUploadSignatureResponseBodyData) SetOssSignature(v string) *GetDataAgentThemeUploadSignatureResponseBodyData {
	s.OssSignature = &v
	return s
}

func (s *GetDataAgentThemeUploadSignatureResponseBodyData) SetOssSignatureVersion(v string) *GetDataAgentThemeUploadSignatureResponseBodyData {
	s.OssSignatureVersion = &v
	return s
}

func (s *GetDataAgentThemeUploadSignatureResponseBodyData) SetPolicy(v string) *GetDataAgentThemeUploadSignatureResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GetDataAgentThemeUploadSignatureResponseBodyData) SetThemeId(v string) *GetDataAgentThemeUploadSignatureResponseBodyData {
	s.ThemeId = &v
	return s
}

func (s *GetDataAgentThemeUploadSignatureResponseBodyData) SetUploadDir(v string) *GetDataAgentThemeUploadSignatureResponseBodyData {
	s.UploadDir = &v
	return s
}

func (s *GetDataAgentThemeUploadSignatureResponseBodyData) SetUploadHost(v string) *GetDataAgentThemeUploadSignatureResponseBodyData {
	s.UploadHost = &v
	return s
}

func (s *GetDataAgentThemeUploadSignatureResponseBodyData) Validate() error {
	return dara.Validate(s)
}
