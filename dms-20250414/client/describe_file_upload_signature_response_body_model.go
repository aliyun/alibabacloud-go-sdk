// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileUploadSignatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeFileUploadSignatureResponseBodyData) *DescribeFileUploadSignatureResponseBody
	GetData() *DescribeFileUploadSignatureResponseBodyData
	SetErrorCode(v string) *DescribeFileUploadSignatureResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DescribeFileUploadSignatureResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DescribeFileUploadSignatureResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeFileUploadSignatureResponseBody
	GetSuccess() *bool
}

type DescribeFileUploadSignatureResponseBody struct {
	Data *DescribeFileUploadSignatureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Specified parameter Tid is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0FEE5834-C55A-5995-A6A3-B443304965BD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFileUploadSignatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileUploadSignatureResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFileUploadSignatureResponseBody) GetData() *DescribeFileUploadSignatureResponseBodyData {
	return s.Data
}

func (s *DescribeFileUploadSignatureResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeFileUploadSignatureResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeFileUploadSignatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFileUploadSignatureResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeFileUploadSignatureResponseBody) SetData(v *DescribeFileUploadSignatureResponseBodyData) *DescribeFileUploadSignatureResponseBody {
	s.Data = v
	return s
}

func (s *DescribeFileUploadSignatureResponseBody) SetErrorCode(v string) *DescribeFileUploadSignatureResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFileUploadSignatureResponseBody) SetErrorMessage(v string) *DescribeFileUploadSignatureResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeFileUploadSignatureResponseBody) SetRequestId(v string) *DescribeFileUploadSignatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFileUploadSignatureResponseBody) SetSuccess(v bool) *DescribeFileUploadSignatureResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFileUploadSignatureResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeFileUploadSignatureResponseBodyData struct {
	OssCredential *string `json:"OssCredential,omitempty" xml:"OssCredential,omitempty"`
	// example:
	//
	// 20260101T135341Z
	OssDate *string `json:"OssDate,omitempty" xml:"OssDate,omitempty"`
	// example:
	//
	// CAIS4gJ1q6Ft5B2yfSjIr5vPHMj4p+lHx/utUUjg13ptZ+5u3oDzkzz2IHhMdXlrCOgYt/8xnG1V6f8flrJ/ToQAX0HfatZq5ZkS9AqnaoXM/te496IFg5D9y7dIs8GgjqHoeOzcYI73WJXEMiLp9EJaxb/9ak/RPTiMOoGIjphKd8keWhLCAxNNGNZRIHkJyqZYTwyzU8ygKRn3mGHdIVN1sw5n8wNF5L+439eX52i17jS46JdM/9ysesH5NpQxbMwkDYnk5oEsKPqdihw3wgNR6aJ7gJZD/Tr6pdyHCzFTmU7ea7uEqYw3clYiOPBnRvEd8eKPnPl5q/HVm4Hs0wxKNuxOSCXZS4yp3MLeH+ekJgOGwWFHz9qnOLmtQXqV22tMCRpzXIj6Zlmz+/reI6iNW+Ory74mxSFbrz3ZP4yv+o+Yv3QbMVumcySkKVbBbVvnv0R8GNsIC2lMUbp+rfShhfFuG2QagAECCyigwAlSAryrFmteD+EVuvxvi0NE7zDJLbUkhek6dcY+/u5V5jcmvL67CQ7bTNk+9lV8WDCvtoCD9ucqTaHweJEd8fS2DaFedAMDf8BfZa2C1CTLhVXdSgE2WORYbMqidelRm7dH3fTbZVvryWKDaveDRLt5J/Qfs**********
	OssSecurityToken *string `json:"OssSecurityToken,omitempty" xml:"OssSecurityToken,omitempty"`
	// example:
	//
	// 9bebe0900716bdefaab899781c7bdfd614ec6ed711e0de5ddf6f5a**********
	OssSignature *string `json:"OssSignature,omitempty" xml:"OssSignature,omitempty"`
	// example:
	//
	// OSS4-HMAC-SHA256
	OssSignatureVersion *string `json:"OssSignatureVersion,omitempty" xml:"OssSignatureVersion,omitempty"`
	// example:
	//
	// eyJjb25kaXRpb25zIjpbeyJ4LW9zcy1jcmVkZW50aWFsIjoiU1RTLk5YeldyTEo2ZnA5RlNuUTN6OGthQjFFWH**********
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	UploadDir *string `json:"UploadDir,omitempty" xml:"UploadDir,omitempty"`
	// example:
	//
	// https://**********.oss-cn-hangzhou.aliyuncs.com
	UploadHost *string `json:"UploadHost,omitempty" xml:"UploadHost,omitempty"`
}

func (s DescribeFileUploadSignatureResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileUploadSignatureResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeFileUploadSignatureResponseBodyData) GetOssCredential() *string {
	return s.OssCredential
}

func (s *DescribeFileUploadSignatureResponseBodyData) GetOssDate() *string {
	return s.OssDate
}

func (s *DescribeFileUploadSignatureResponseBodyData) GetOssSecurityToken() *string {
	return s.OssSecurityToken
}

func (s *DescribeFileUploadSignatureResponseBodyData) GetOssSignature() *string {
	return s.OssSignature
}

func (s *DescribeFileUploadSignatureResponseBodyData) GetOssSignatureVersion() *string {
	return s.OssSignatureVersion
}

func (s *DescribeFileUploadSignatureResponseBodyData) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeFileUploadSignatureResponseBodyData) GetUploadDir() *string {
	return s.UploadDir
}

func (s *DescribeFileUploadSignatureResponseBodyData) GetUploadHost() *string {
	return s.UploadHost
}

func (s *DescribeFileUploadSignatureResponseBodyData) SetOssCredential(v string) *DescribeFileUploadSignatureResponseBodyData {
	s.OssCredential = &v
	return s
}

func (s *DescribeFileUploadSignatureResponseBodyData) SetOssDate(v string) *DescribeFileUploadSignatureResponseBodyData {
	s.OssDate = &v
	return s
}

func (s *DescribeFileUploadSignatureResponseBodyData) SetOssSecurityToken(v string) *DescribeFileUploadSignatureResponseBodyData {
	s.OssSecurityToken = &v
	return s
}

func (s *DescribeFileUploadSignatureResponseBodyData) SetOssSignature(v string) *DescribeFileUploadSignatureResponseBodyData {
	s.OssSignature = &v
	return s
}

func (s *DescribeFileUploadSignatureResponseBodyData) SetOssSignatureVersion(v string) *DescribeFileUploadSignatureResponseBodyData {
	s.OssSignatureVersion = &v
	return s
}

func (s *DescribeFileUploadSignatureResponseBodyData) SetPolicy(v string) *DescribeFileUploadSignatureResponseBodyData {
	s.Policy = &v
	return s
}

func (s *DescribeFileUploadSignatureResponseBodyData) SetUploadDir(v string) *DescribeFileUploadSignatureResponseBodyData {
	s.UploadDir = &v
	return s
}

func (s *DescribeFileUploadSignatureResponseBodyData) SetUploadHost(v string) *DescribeFileUploadSignatureResponseBodyData {
	s.UploadHost = &v
	return s
}

func (s *DescribeFileUploadSignatureResponseBodyData) Validate() error {
	return dara.Validate(s)
}
