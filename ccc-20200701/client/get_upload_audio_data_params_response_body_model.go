// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadAudioDataParamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetUploadAudioDataParamsResponseBody
	GetCode() *string
	SetData(v *GetUploadAudioDataParamsResponseBodyData) *GetUploadAudioDataParamsResponseBody
	GetData() *GetUploadAudioDataParamsResponseBodyData
	SetHttpStatusCode(v int32) *GetUploadAudioDataParamsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetUploadAudioDataParamsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUploadAudioDataParamsResponseBody
	GetRequestId() *string
}

type GetUploadAudioDataParamsResponseBody struct {
	// example:
	//
	// OK
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetUploadAudioDataParamsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 03C67DAD-EB26-41D8-949D-9B0C470FB716
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUploadAudioDataParamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUploadAudioDataParamsResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadAudioDataParamsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUploadAudioDataParamsResponseBody) GetData() *GetUploadAudioDataParamsResponseBodyData {
	return s.Data
}

func (s *GetUploadAudioDataParamsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetUploadAudioDataParamsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUploadAudioDataParamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUploadAudioDataParamsResponseBody) SetCode(v string) *GetUploadAudioDataParamsResponseBody {
	s.Code = &v
	return s
}

func (s *GetUploadAudioDataParamsResponseBody) SetData(v *GetUploadAudioDataParamsResponseBodyData) *GetUploadAudioDataParamsResponseBody {
	s.Data = v
	return s
}

func (s *GetUploadAudioDataParamsResponseBody) SetHttpStatusCode(v int32) *GetUploadAudioDataParamsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUploadAudioDataParamsResponseBody) SetMessage(v string) *GetUploadAudioDataParamsResponseBody {
	s.Message = &v
	return s
}

func (s *GetUploadAudioDataParamsResponseBody) SetRequestId(v string) *GetUploadAudioDataParamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUploadAudioDataParamsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUploadAudioDataParamsResponseBodyData struct {
	ParamsStr *string `json:"ParamsStr,omitempty" xml:"ParamsStr,omitempty"`
}

func (s GetUploadAudioDataParamsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetUploadAudioDataParamsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUploadAudioDataParamsResponseBodyData) GetParamsStr() *string {
	return s.ParamsStr
}

func (s *GetUploadAudioDataParamsResponseBodyData) SetParamsStr(v string) *GetUploadAudioDataParamsResponseBodyData {
	s.ParamsStr = &v
	return s
}

func (s *GetUploadAudioDataParamsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
