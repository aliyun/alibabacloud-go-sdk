// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadPrivacyKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DownloadPrivacyKeyResponseBody
	GetCode() *string
	SetData(v string) *DownloadPrivacyKeyResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DownloadPrivacyKeyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DownloadPrivacyKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DownloadPrivacyKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DownloadPrivacyKeyResponseBody
	GetSuccess() *bool
}

type DownloadPrivacyKeyResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DownloadPrivacyKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DownloadPrivacyKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadPrivacyKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DownloadPrivacyKeyResponseBody) GetData() *string {
	return s.Data
}

func (s *DownloadPrivacyKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DownloadPrivacyKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DownloadPrivacyKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DownloadPrivacyKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DownloadPrivacyKeyResponseBody) SetCode(v string) *DownloadPrivacyKeyResponseBody {
	s.Code = &v
	return s
}

func (s *DownloadPrivacyKeyResponseBody) SetData(v string) *DownloadPrivacyKeyResponseBody {
	s.Data = &v
	return s
}

func (s *DownloadPrivacyKeyResponseBody) SetHttpStatusCode(v int32) *DownloadPrivacyKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DownloadPrivacyKeyResponseBody) SetMessage(v string) *DownloadPrivacyKeyResponseBody {
	s.Message = &v
	return s
}

func (s *DownloadPrivacyKeyResponseBody) SetRequestId(v string) *DownloadPrivacyKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadPrivacyKeyResponseBody) SetSuccess(v bool) *DownloadPrivacyKeyResponseBody {
	s.Success = &v
	return s
}

func (s *DownloadPrivacyKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
