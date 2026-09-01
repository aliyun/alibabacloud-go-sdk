// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSkillFileUploadSignatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeSkillFileUploadSignatureResponseBodyData) *DescribeSkillFileUploadSignatureResponseBody
	GetData() *DescribeSkillFileUploadSignatureResponseBodyData
	SetErrorCode(v string) *DescribeSkillFileUploadSignatureResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DescribeSkillFileUploadSignatureResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DescribeSkillFileUploadSignatureResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSkillFileUploadSignatureResponseBody
	GetSuccess() *bool
}

type DescribeSkillFileUploadSignatureResponseBody struct {
	// The response struct.
	Data *DescribeSkillFileUploadSignatureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// InvalidTid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the call fails.
	//
	// example:
	//
	// Specified parameter Tid is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E0D21075-CD3E-4D98-8264-FD8AD04A63B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSkillFileUploadSignatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSkillFileUploadSignatureResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSkillFileUploadSignatureResponseBody) GetData() *DescribeSkillFileUploadSignatureResponseBodyData {
	return s.Data
}

func (s *DescribeSkillFileUploadSignatureResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeSkillFileUploadSignatureResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSkillFileUploadSignatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSkillFileUploadSignatureResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSkillFileUploadSignatureResponseBody) SetData(v *DescribeSkillFileUploadSignatureResponseBodyData) *DescribeSkillFileUploadSignatureResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSkillFileUploadSignatureResponseBody) SetErrorCode(v string) *DescribeSkillFileUploadSignatureResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSkillFileUploadSignatureResponseBody) SetErrorMessage(v string) *DescribeSkillFileUploadSignatureResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSkillFileUploadSignatureResponseBody) SetRequestId(v string) *DescribeSkillFileUploadSignatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSkillFileUploadSignatureResponseBody) SetSuccess(v bool) *DescribeSkillFileUploadSignatureResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSkillFileUploadSignatureResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSkillFileUploadSignatureResponseBodyData struct {
	// The parameter set that specifies the derived key.
	OssCredential *string `json:"OssCredential,omitempty" xml:"OssCredential,omitempty"`
	// The time of the request. The format follows the ISO 8601 date and time standard.
	//
	// example:
	//
	// 20260101T135341Z
	OssDate *string `json:"OssDate,omitempty" xml:"OssDate,omitempty"`
	// The STS token used for uploading to OSS. The token is valid for 1 hour.
	//
	// example:
	//
	// CAIS4gJ1q6Ft5B2yfSjIr5vPHMj4p+lHx/utUUjg13ptZ+5u3oDzkzz2IHhMdXlrCOgYt/8xnG1V6f8flrJ/ToQAX0HfatZq5ZkS9AqnaoXM/te496IFg5D9y7dIs8GgjqHoeOzcYI73WJXEMiLp9EJaxb/9ak/RPTiMOoGIjphKd8keWhLCAxNNGNZRIHkJyqZYTwyzU8ygKRn3mGHdIVN1sw5n8wNF5L+439eX52i17jS46JdM/9ysesH5NpQxbMwkDYnk5oEsKPqdihw3wgNR6aJ7gJZD/Tr6pdyHCzFTmU7ea7uEqYw3clYiOPBnRvEd8eKPnPl5q/HVm4Hs0wxKNuxOSCXZS4yp3MLeH+ekJgOGwWFHz9qnOLmtQXqV22tMCRpzXIj6Zlmz+/reI6iNW+Ory74mxSFbrz3ZP4yv+o+Yv3QbMVumcySkKVbBbVvnv0R8GNsIC2lMUbp+rfShhfFuG2QagAECCyigwAlSAryrFmteD+EVuvxvi0NE7zDJLbUkhek6dcY+/u5V5jcmvL67CQ7bTNk+9lV8WDCvtoCD9ucqTaHweJEd8fS2DaFedAMDf8BfZa2C1CTLhVXdSgE2WORYbMqidelRm7dH3fTbZVvryWKDaveDRLt5J/Qfs**********
	OssSecurityToken *string `json:"OssSecurityToken,omitempty" xml:"OssSecurityToken,omitempty"`
	// The description information used for signature authentication.
	//
	// example:
	//
	// 9bebe0900716bdefaab899781c7bdfd614ec6ed711e0de5ddf6f5a**********
	OssSignature *string `json:"OssSignature,omitempty" xml:"OssSignature,omitempty"`
	// The version and algorithm of the signature.
	//
	// example:
	//
	// OSS4-HMAC-SHA256
	OssSignatureVersion *string `json:"OssSignatureVersion,omitempty" xml:"OssSignatureVersion,omitempty"`
	// The permission restrictions and constraints for file upload.
	//
	// example:
	//
	// eyJjb25kaXRpb25zIjpbeyJ4LW9zcy1jcmVkZW50aWFsIjoiU1RTLk5YeldyTEo2ZnA5RlNuUTN6OGthQjFFWH**********
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The path for file upload.
	UploadDir *string `json:"UploadDir,omitempty" xml:"UploadDir,omitempty"`
	// The destination address for file upload.
	//
	// example:
	//
	// https://**********.oss-cn-hangzhou.aliyuncs.com
	UploadHost *string `json:"UploadHost,omitempty" xml:"UploadHost,omitempty"`
}

func (s DescribeSkillFileUploadSignatureResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSkillFileUploadSignatureResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSkillFileUploadSignatureResponseBodyData) GetOssCredential() *string {
	return s.OssCredential
}

func (s *DescribeSkillFileUploadSignatureResponseBodyData) GetOssDate() *string {
	return s.OssDate
}

func (s *DescribeSkillFileUploadSignatureResponseBodyData) GetOssSecurityToken() *string {
	return s.OssSecurityToken
}

func (s *DescribeSkillFileUploadSignatureResponseBodyData) GetOssSignature() *string {
	return s.OssSignature
}

func (s *DescribeSkillFileUploadSignatureResponseBodyData) GetOssSignatureVersion() *string {
	return s.OssSignatureVersion
}

func (s *DescribeSkillFileUploadSignatureResponseBodyData) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeSkillFileUploadSignatureResponseBodyData) GetUploadDir() *string {
	return s.UploadDir
}

func (s *DescribeSkillFileUploadSignatureResponseBodyData) GetUploadHost() *string {
	return s.UploadHost
}

func (s *DescribeSkillFileUploadSignatureResponseBodyData) SetOssCredential(v string) *DescribeSkillFileUploadSignatureResponseBodyData {
	s.OssCredential = &v
	return s
}

func (s *DescribeSkillFileUploadSignatureResponseBodyData) SetOssDate(v string) *DescribeSkillFileUploadSignatureResponseBodyData {
	s.OssDate = &v
	return s
}

func (s *DescribeSkillFileUploadSignatureResponseBodyData) SetOssSecurityToken(v string) *DescribeSkillFileUploadSignatureResponseBodyData {
	s.OssSecurityToken = &v
	return s
}

func (s *DescribeSkillFileUploadSignatureResponseBodyData) SetOssSignature(v string) *DescribeSkillFileUploadSignatureResponseBodyData {
	s.OssSignature = &v
	return s
}

func (s *DescribeSkillFileUploadSignatureResponseBodyData) SetOssSignatureVersion(v string) *DescribeSkillFileUploadSignatureResponseBodyData {
	s.OssSignatureVersion = &v
	return s
}

func (s *DescribeSkillFileUploadSignatureResponseBodyData) SetPolicy(v string) *DescribeSkillFileUploadSignatureResponseBodyData {
	s.Policy = &v
	return s
}

func (s *DescribeSkillFileUploadSignatureResponseBodyData) SetUploadDir(v string) *DescribeSkillFileUploadSignatureResponseBodyData {
	s.UploadDir = &v
	return s
}

func (s *DescribeSkillFileUploadSignatureResponseBodyData) SetUploadHost(v string) *DescribeSkillFileUploadSignatureResponseBodyData {
	s.UploadHost = &v
	return s
}

func (s *DescribeSkillFileUploadSignatureResponseBodyData) Validate() error {
	return dara.Validate(s)
}
