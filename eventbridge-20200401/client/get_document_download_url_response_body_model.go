// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDocumentDownloadUrlResponseBody
	GetCode() *string
	SetData(v *GetDocumentDownloadUrlResponseBodyData) *GetDocumentDownloadUrlResponseBody
	GetData() *GetDocumentDownloadUrlResponseBodyData
	SetMessage(v string) *GetDocumentDownloadUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDocumentDownloadUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDocumentDownloadUrlResponseBody
	GetSuccess() *bool
}

type GetDocumentDownloadUrlResponseBody struct {
	// The response code. A value of Success indicates a successful call. If the call fails, a specific error code is returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The document download URL information.
	Data *GetDocumentDownloadUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned by the operation. A value of Operation success indicates a successful call. If the call fails, a specific error description is returned.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. A value of true indicates success.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDocumentDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentDownloadUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDocumentDownloadUrlResponseBody) GetData() *GetDocumentDownloadUrlResponseBodyData {
	return s.Data
}

func (s *GetDocumentDownloadUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDocumentDownloadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocumentDownloadUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDocumentDownloadUrlResponseBody) SetCode(v string) *GetDocumentDownloadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetDocumentDownloadUrlResponseBody) SetData(v *GetDocumentDownloadUrlResponseBodyData) *GetDocumentDownloadUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetDocumentDownloadUrlResponseBody) SetMessage(v string) *GetDocumentDownloadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentDownloadUrlResponseBody) SetRequestId(v string) *GetDocumentDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentDownloadUrlResponseBody) SetSuccess(v bool) *GetDocumentDownloadUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetDocumentDownloadUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDocumentDownloadUrlResponseBodyData struct {
	// A short-lived GET pre-signed URL that the client uses to download the original file.
	//
	// example:
	//
	// https://my-bucket.oss-cn-hangzhou.aliyuncs.com/kb/doc-bp1xxxxxxxxxxxx?Expires=1788000000&Signature=xxxx
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The expiration time of the pre-signed download URL, in ISO 8601 UTC format.
	//
	// example:
	//
	// 1756022400000
	ExpireAt *string `json:"ExpireAt,omitempty" xml:"ExpireAt,omitempty"`
}

func (s GetDocumentDownloadUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentDownloadUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocumentDownloadUrlResponseBodyData) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetDocumentDownloadUrlResponseBodyData) GetExpireAt() *string {
	return s.ExpireAt
}

func (s *GetDocumentDownloadUrlResponseBodyData) SetDownloadUrl(v string) *GetDocumentDownloadUrlResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *GetDocumentDownloadUrlResponseBodyData) SetExpireAt(v string) *GetDocumentDownloadUrlResponseBodyData {
	s.ExpireAt = &v
	return s
}

func (s *GetDocumentDownloadUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
