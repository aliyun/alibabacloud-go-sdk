// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeObjectUploadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateKnowledgeObjectUploadResponseBody
	GetCode() *string
	SetData(v *CreateKnowledgeObjectUploadResponseBodyData) *CreateKnowledgeObjectUploadResponseBody
	GetData() *CreateKnowledgeObjectUploadResponseBodyData
	SetMessage(v string) *CreateKnowledgeObjectUploadResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateKnowledgeObjectUploadResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateKnowledgeObjectUploadResponseBody
	GetSuccess() *bool
}

type CreateKnowledgeObjectUploadResponseBody struct {
	// The business error code, which is used to determine the error type when the request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned when the request is successful.
	Data *CreateKnowledgeObjectUploadResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned when the request fails.
	//
	// example:
	//
	// The request parameters are invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateKnowledgeObjectUploadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectUploadResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectUploadResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateKnowledgeObjectUploadResponseBody) GetData() *CreateKnowledgeObjectUploadResponseBodyData {
	return s.Data
}

func (s *CreateKnowledgeObjectUploadResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateKnowledgeObjectUploadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKnowledgeObjectUploadResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateKnowledgeObjectUploadResponseBody) SetCode(v string) *CreateKnowledgeObjectUploadResponseBody {
	s.Code = &v
	return s
}

func (s *CreateKnowledgeObjectUploadResponseBody) SetData(v *CreateKnowledgeObjectUploadResponseBodyData) *CreateKnowledgeObjectUploadResponseBody {
	s.Data = v
	return s
}

func (s *CreateKnowledgeObjectUploadResponseBody) SetMessage(v string) *CreateKnowledgeObjectUploadResponseBody {
	s.Message = &v
	return s
}

func (s *CreateKnowledgeObjectUploadResponseBody) SetRequestId(v string) *CreateKnowledgeObjectUploadResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKnowledgeObjectUploadResponseBody) SetSuccess(v bool) *CreateKnowledgeObjectUploadResponseBody {
	s.Success = &v
	return s
}

func (s *CreateKnowledgeObjectUploadResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateKnowledgeObjectUploadResponseBodyData struct {
	// The expiration time of the UploadUrl, represented as a UNIX timestamp in milliseconds. The maximum validity period is 30 minutes, subject to the upper limit of the OSS signature configuration and the remaining time of the session. This field prevails.
	//
	// example:
	//
	// 1792456200000
	ExpiresAt *int64 `json:"ExpiresAt,omitempty" xml:"ExpiresAt,omitempty"`
	// The HTTP headers that must be included in the direct upload request. The key is the actual header name, and the value is the header value.
	//
	// example:
	//
	// {"Content-Type":"text/plain","Content-MD5":"1B2M2Y8AsgTpgAmY7PhCfg=="}
	RequiredHeaders map[string]*string `json:"RequiredHeaders,omitempty" xml:"RequiredHeaders,omitempty"`
	// The identifier of the upload session to be consumed by document creation, model update, or model import tasks. This is not a KnowledgeObjectId or TaskId.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abd
	UploadId *string `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
	// The temporary direct upload URL.
	//
	// example:
	//
	// https://example.invalid/short-lived-upload
	UploadUrl *string `json:"UploadUrl,omitempty" xml:"UploadUrl,omitempty"`
}

func (s CreateKnowledgeObjectUploadResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectUploadResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectUploadResponseBodyData) GetExpiresAt() *int64 {
	return s.ExpiresAt
}

func (s *CreateKnowledgeObjectUploadResponseBodyData) GetRequiredHeaders() map[string]*string {
	return s.RequiredHeaders
}

func (s *CreateKnowledgeObjectUploadResponseBodyData) GetUploadId() *string {
	return s.UploadId
}

func (s *CreateKnowledgeObjectUploadResponseBodyData) GetUploadUrl() *string {
	return s.UploadUrl
}

func (s *CreateKnowledgeObjectUploadResponseBodyData) SetExpiresAt(v int64) *CreateKnowledgeObjectUploadResponseBodyData {
	s.ExpiresAt = &v
	return s
}

func (s *CreateKnowledgeObjectUploadResponseBodyData) SetRequiredHeaders(v map[string]*string) *CreateKnowledgeObjectUploadResponseBodyData {
	s.RequiredHeaders = v
	return s
}

func (s *CreateKnowledgeObjectUploadResponseBodyData) SetUploadId(v string) *CreateKnowledgeObjectUploadResponseBodyData {
	s.UploadId = &v
	return s
}

func (s *CreateKnowledgeObjectUploadResponseBodyData) SetUploadUrl(v string) *CreateKnowledgeObjectUploadResponseBodyData {
	s.UploadUrl = &v
	return s
}

func (s *CreateKnowledgeObjectUploadResponseBodyData) Validate() error {
	return dara.Validate(s)
}
