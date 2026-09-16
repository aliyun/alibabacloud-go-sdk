// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLumaDocumentDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetLumaDocumentDownloadUrlResponseBody
	GetCode() *string
	SetData(v *GetLumaDocumentDownloadUrlResponseBodyData) *GetLumaDocumentDownloadUrlResponseBody
	GetData() *GetLumaDocumentDownloadUrlResponseBodyData
	SetMessage(v string) *GetLumaDocumentDownloadUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLumaDocumentDownloadUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLumaDocumentDownloadUrlResponseBody
	GetSuccess() *bool
}

type GetLumaDocumentDownloadUrlResponseBody struct {
	// The response code. A value of Success indicates that the call succeeds. If the call fails, a specific error code is returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The pre-signed download URL information for the original document.
	Data *GetLumaDocumentDownloadUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned by the operation. The value is Operation success if the call succeeds, or a specific error description if the call fails.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request. Use this ID for troubleshooting or when submitting a ticket.
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates that the call succeeds.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLumaDocumentDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLumaDocumentDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetLumaDocumentDownloadUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetLumaDocumentDownloadUrlResponseBody) GetData() *GetLumaDocumentDownloadUrlResponseBodyData {
	return s.Data
}

func (s *GetLumaDocumentDownloadUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLumaDocumentDownloadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLumaDocumentDownloadUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLumaDocumentDownloadUrlResponseBody) SetCode(v string) *GetLumaDocumentDownloadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetLumaDocumentDownloadUrlResponseBody) SetData(v *GetLumaDocumentDownloadUrlResponseBodyData) *GetLumaDocumentDownloadUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetLumaDocumentDownloadUrlResponseBody) SetMessage(v string) *GetLumaDocumentDownloadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetLumaDocumentDownloadUrlResponseBody) SetRequestId(v string) *GetLumaDocumentDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLumaDocumentDownloadUrlResponseBody) SetSuccess(v bool) *GetLumaDocumentDownloadUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetLumaDocumentDownloadUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLumaDocumentDownloadUrlResponseBodyData struct {
	// The pre-signed download URL for the original document. The URL is valid for a limited period of time.
	//
	// example:
	//
	// https://my-bucket.oss-cn-hangzhou.aliyuncs.com/kb/doc-bp1xxxxxxxxxxxx?Expires=1788000000&Signature=xxxx
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The expiration time of the download URL in UTC.
	//
	// example:
	//
	// 2026-08-24T12:00:00Z
	ExpireAt *string `json:"ExpireAt,omitempty" xml:"ExpireAt,omitempty"`
}

func (s GetLumaDocumentDownloadUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetLumaDocumentDownloadUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLumaDocumentDownloadUrlResponseBodyData) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetLumaDocumentDownloadUrlResponseBodyData) GetExpireAt() *string {
	return s.ExpireAt
}

func (s *GetLumaDocumentDownloadUrlResponseBodyData) SetDownloadUrl(v string) *GetLumaDocumentDownloadUrlResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *GetLumaDocumentDownloadUrlResponseBodyData) SetExpireAt(v string) *GetLumaDocumentDownloadUrlResponseBodyData {
	s.ExpireAt = &v
	return s
}

func (s *GetLumaDocumentDownloadUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
