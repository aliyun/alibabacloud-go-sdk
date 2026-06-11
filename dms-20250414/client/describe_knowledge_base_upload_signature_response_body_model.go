// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKnowledgeBaseUploadSignatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeKnowledgeBaseUploadSignatureResponseBodyData) *DescribeKnowledgeBaseUploadSignatureResponseBody
	GetData() *DescribeKnowledgeBaseUploadSignatureResponseBodyData
	SetErrorCode(v string) *DescribeKnowledgeBaseUploadSignatureResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DescribeKnowledgeBaseUploadSignatureResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DescribeKnowledgeBaseUploadSignatureResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeKnowledgeBaseUploadSignatureResponseBody
	GetSuccess() *bool
}

type DescribeKnowledgeBaseUploadSignatureResponseBody struct {
	// The upload signature details.
	Data *DescribeKnowledgeBaseUploadSignatureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code if the request fails.
	//
	// example:
	//
	// KnowledgeBaseNotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message if the request fails.
	//
	// example:
	//
	// Resource not found xxx
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The unique ID of the request. If an error occurs, use this ID to troubleshoot the issue.
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
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeKnowledgeBaseUploadSignatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseUploadSignatureResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBody) GetData() *DescribeKnowledgeBaseUploadSignatureResponseBodyData {
	return s.Data
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBody) SetData(v *DescribeKnowledgeBaseUploadSignatureResponseBodyData) *DescribeKnowledgeBaseUploadSignatureResponseBody {
	s.Data = v
	return s
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBody) SetErrorCode(v string) *DescribeKnowledgeBaseUploadSignatureResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBody) SetErrorMessage(v string) *DescribeKnowledgeBaseUploadSignatureResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBody) SetRequestId(v string) *DescribeKnowledgeBaseUploadSignatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBody) SetSuccess(v bool) *DescribeKnowledgeBaseUploadSignatureResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeKnowledgeBaseUploadSignatureResponseBodyData struct {
	// The credential scope string for the signature.
	OssCredential *string `json:"OssCredential,omitempty" xml:"OssCredential,omitempty"`
	// The request time in ISO 8601 format.
	//
	// example:
	//
	// 20260101T135341Z
	OssDate *string `json:"OssDate,omitempty" xml:"OssDate,omitempty"`
	// The STS token used for uploading to OSS. It is valid for one hour.
	//
	// example:
	//
	// CAIS4gJ1q6Ft5B2yfSjIr5vPHMj4p+lHx/utUUjg13ptZ+5u3oDzkzz2IHhMdXlrCOgYt/8xnG1V6f8flrJ/ToQAX0HfatZq5ZkS9AqnaoXM/te496IFg5D9y7dIs8GgjqHoeOzcYI73WJXEMiLp9EJaxb/9ak/RPTiMOoGIjphKd8keWhLCAxNNGNZRIHkJyqZYTwyzU8ygKRn3mGHdIVN1sw5n8wNF5L+439eX52i17jS46JdM/9ysesH5NpQxbMwkDYnk5oEsKPqdihw3wgNR6aJ7gJZD/Tr6pdyHCzFTmU7ea7uEqYw3clYiOPBnRvEd8eKPnPl5q/HVm4Hs0wxKNuxOSCXZS4yp3MLeH+ekJgOGwWFHz9qnOLmtQXqV22tMCRpzXIj6Zlmz+/reI6iNW+Ory74mxSFbrz3ZP4yv+o+Yv3QbMVumcySkKVbBbVvnv0R8GNsIC2lMUbp+rfShhfFuG2QagAECCyigwAlSAryrFmteD+EVuvxvi0NE7zDJLbUkhek6dcY+/u5V5jcmvL67CQ7bTNk+9lV8WDCvtoCD9ucqTaHweJEd8fS2DaFedAMDf8BfZa2C1CTLhVXdSgE2WORYbMqidelRm7dH3fTbZVvryWKDaveDRLt5J/Qfs**********
	OssSecurityToken *string `json:"OssSecurityToken,omitempty" xml:"OssSecurityToken,omitempty"`
	// The authentication signature.
	//
	// example:
	//
	// 9bebe0900716bdefaab899781c7bdfd614ec6ed711e0de5ddf6f5a**********
	OssSignature *string `json:"OssSignature,omitempty" xml:"OssSignature,omitempty"`
	// The signature version and algorithm.
	//
	// example:
	//
	// OSS4-HMAC-SHA256
	OssSignatureVersion *string `json:"OssSignatureVersion,omitempty" xml:"OssSignatureVersion,omitempty"`
	// The Base64-encoded POST policy that specifies the conditions for the file upload.
	//
	// example:
	//
	// eyJjb25kaXRpb25zIjpbeyJ4LW9zcy1jcmVkZW50aWFsIjoiU1RTLk5aRmR2RDJRWlFSeWlwZmpkS295NEYxb2ovMjAyNjA1MTQvY24taGFuZ3pob3Uvb3NzL2FsaXl1bl92NF9yZXF1ZXN0In0seyJ4LW9zcy1kYXRlIjoiMjAyNjA1MTRUMDMzMjI3WiJ9LHsieC1vc3Mtc2VjdXJpdHktdG9rZW4iOiJDQUlTMmdKMXE2RnQ1QjJ5ZlNqSXI1bnpMOHp3MzQ1NzVwQ1NhMWJYam1RZVkvWVlxZlRFaUR6MklIaE1**
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The path prefix for the file upload.
	UploadDir *string `json:"UploadDir,omitempty" xml:"UploadDir,omitempty"`
	// The destination URL for the file upload.
	//
	// example:
	//
	// https://onemeta-kb-staging.oss-cn-hangzhou.aliyuncs.com
	UploadHost *string `json:"UploadHost,omitempty" xml:"UploadHost,omitempty"`
}

func (s DescribeKnowledgeBaseUploadSignatureResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseUploadSignatureResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBodyData) GetOssCredential() *string {
	return s.OssCredential
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBodyData) GetOssDate() *string {
	return s.OssDate
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBodyData) GetOssSecurityToken() *string {
	return s.OssSecurityToken
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBodyData) GetOssSignature() *string {
	return s.OssSignature
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBodyData) GetOssSignatureVersion() *string {
	return s.OssSignatureVersion
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBodyData) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBodyData) GetUploadDir() *string {
	return s.UploadDir
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBodyData) GetUploadHost() *string {
	return s.UploadHost
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBodyData) SetOssCredential(v string) *DescribeKnowledgeBaseUploadSignatureResponseBodyData {
	s.OssCredential = &v
	return s
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBodyData) SetOssDate(v string) *DescribeKnowledgeBaseUploadSignatureResponseBodyData {
	s.OssDate = &v
	return s
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBodyData) SetOssSecurityToken(v string) *DescribeKnowledgeBaseUploadSignatureResponseBodyData {
	s.OssSecurityToken = &v
	return s
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBodyData) SetOssSignature(v string) *DescribeKnowledgeBaseUploadSignatureResponseBodyData {
	s.OssSignature = &v
	return s
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBodyData) SetOssSignatureVersion(v string) *DescribeKnowledgeBaseUploadSignatureResponseBodyData {
	s.OssSignatureVersion = &v
	return s
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBodyData) SetPolicy(v string) *DescribeKnowledgeBaseUploadSignatureResponseBodyData {
	s.Policy = &v
	return s
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBodyData) SetUploadDir(v string) *DescribeKnowledgeBaseUploadSignatureResponseBodyData {
	s.UploadDir = &v
	return s
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBodyData) SetUploadHost(v string) *DescribeKnowledgeBaseUploadSignatureResponseBodyData {
	s.UploadHost = &v
	return s
}

func (s *DescribeKnowledgeBaseUploadSignatureResponseBodyData) Validate() error {
	return dara.Validate(s)
}
