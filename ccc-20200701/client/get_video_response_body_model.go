// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetVideoResponseBody
	GetCode() *string
	SetData(v *GetVideoResponseBodyData) *GetVideoResponseBody
	GetData() *GetVideoResponseBodyData
	SetHttpStatusCode(v int32) *GetVideoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetVideoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetVideoResponseBody
	GetRequestId() *string
}

type GetVideoResponseBody struct {
	Code           *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetVideoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVideoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetVideoResponseBody) GetData() *GetVideoResponseBodyData {
	return s.Data
}

func (s *GetVideoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetVideoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVideoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoResponseBody) SetCode(v string) *GetVideoResponseBody {
	s.Code = &v
	return s
}

func (s *GetVideoResponseBody) SetData(v *GetVideoResponseBodyData) *GetVideoResponseBody {
	s.Data = v
	return s
}

func (s *GetVideoResponseBody) SetHttpStatusCode(v int32) *GetVideoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetVideoResponseBody) SetMessage(v string) *GetVideoResponseBody {
	s.Message = &v
	return s
}

func (s *GetVideoResponseBody) SetRequestId(v string) *GetVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetVideoResponseBodyData struct {
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileUrl  *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s GetVideoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetVideoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVideoResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *GetVideoResponseBodyData) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetVideoResponseBodyData) SetFileName(v string) *GetVideoResponseBodyData {
	s.FileName = &v
	return s
}

func (s *GetVideoResponseBodyData) SetFileUrl(v string) *GetVideoResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *GetVideoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
