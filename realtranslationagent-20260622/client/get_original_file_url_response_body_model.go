// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginalFileUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOriginalFileUrlResponseBody
	GetCode() *string
	SetData(v *GetOriginalFileUrlResponseBodyData) *GetOriginalFileUrlResponseBody
	GetData() *GetOriginalFileUrlResponseBodyData
	SetMessage(v string) *GetOriginalFileUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetOriginalFileUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOriginalFileUrlResponseBody
	GetSuccess() *bool
}

type GetOriginalFileUrlResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetOriginalFileUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOriginalFileUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOriginalFileUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetOriginalFileUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOriginalFileUrlResponseBody) GetData() *GetOriginalFileUrlResponseBodyData {
	return s.Data
}

func (s *GetOriginalFileUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOriginalFileUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOriginalFileUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOriginalFileUrlResponseBody) SetCode(v string) *GetOriginalFileUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetOriginalFileUrlResponseBody) SetData(v *GetOriginalFileUrlResponseBodyData) *GetOriginalFileUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetOriginalFileUrlResponseBody) SetMessage(v string) *GetOriginalFileUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetOriginalFileUrlResponseBody) SetRequestId(v string) *GetOriginalFileUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOriginalFileUrlResponseBody) SetSuccess(v bool) *GetOriginalFileUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetOriginalFileUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOriginalFileUrlResponseBodyData struct {
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	ExpireTime  *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	FileName    *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s GetOriginalFileUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOriginalFileUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOriginalFileUrlResponseBodyData) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetOriginalFileUrlResponseBodyData) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *GetOriginalFileUrlResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *GetOriginalFileUrlResponseBodyData) SetDownloadUrl(v string) *GetOriginalFileUrlResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *GetOriginalFileUrlResponseBodyData) SetExpireTime(v int64) *GetOriginalFileUrlResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetOriginalFileUrlResponseBodyData) SetFileName(v string) *GetOriginalFileUrlResponseBodyData {
	s.FileName = &v
	return s
}

func (s *GetOriginalFileUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
