// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranslatedFileUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTranslatedFileUrlResponseBody
	GetCode() *string
	SetData(v *GetTranslatedFileUrlResponseBodyData) *GetTranslatedFileUrlResponseBody
	GetData() *GetTranslatedFileUrlResponseBodyData
	SetMessage(v string) *GetTranslatedFileUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTranslatedFileUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTranslatedFileUrlResponseBody
	GetSuccess() *bool
}

type GetTranslatedFileUrlResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetTranslatedFileUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTranslatedFileUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTranslatedFileUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetTranslatedFileUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTranslatedFileUrlResponseBody) GetData() *GetTranslatedFileUrlResponseBodyData {
	return s.Data
}

func (s *GetTranslatedFileUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTranslatedFileUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTranslatedFileUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTranslatedFileUrlResponseBody) SetCode(v string) *GetTranslatedFileUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetTranslatedFileUrlResponseBody) SetData(v *GetTranslatedFileUrlResponseBodyData) *GetTranslatedFileUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetTranslatedFileUrlResponseBody) SetMessage(v string) *GetTranslatedFileUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetTranslatedFileUrlResponseBody) SetRequestId(v string) *GetTranslatedFileUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTranslatedFileUrlResponseBody) SetSuccess(v bool) *GetTranslatedFileUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetTranslatedFileUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranslatedFileUrlResponseBodyData struct {
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	ExpireTime  *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	FileName    *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s GetTranslatedFileUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTranslatedFileUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTranslatedFileUrlResponseBodyData) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetTranslatedFileUrlResponseBodyData) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *GetTranslatedFileUrlResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *GetTranslatedFileUrlResponseBodyData) SetDownloadUrl(v string) *GetTranslatedFileUrlResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *GetTranslatedFileUrlResponseBodyData) SetExpireTime(v int64) *GetTranslatedFileUrlResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetTranslatedFileUrlResponseBodyData) SetFileName(v string) *GetTranslatedFileUrlResponseBodyData {
	s.FileName = &v
	return s
}

func (s *GetTranslatedFileUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
