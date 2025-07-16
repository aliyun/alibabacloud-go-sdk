// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageTranslateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetImageTranslateTaskResponseBody
	GetCode() *int32
	SetData(v *GetImageTranslateTaskResponseBodyData) *GetImageTranslateTaskResponseBody
	GetData() *GetImageTranslateTaskResponseBodyData
	SetMessage(v string) *GetImageTranslateTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetImageTranslateTaskResponseBody
	GetRequestId() *string
}

type GetImageTranslateTaskResponseBody struct {
	// example:
	//
	// 200
	Code *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetImageTranslateTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// xxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetImageTranslateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageTranslateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageTranslateTaskResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetImageTranslateTaskResponseBody) GetData() *GetImageTranslateTaskResponseBodyData {
	return s.Data
}

func (s *GetImageTranslateTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetImageTranslateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageTranslateTaskResponseBody) SetCode(v int32) *GetImageTranslateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetImageTranslateTaskResponseBody) SetData(v *GetImageTranslateTaskResponseBodyData) *GetImageTranslateTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetImageTranslateTaskResponseBody) SetMessage(v string) *GetImageTranslateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetImageTranslateTaskResponseBody) SetRequestId(v string) *GetImageTranslateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageTranslateTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetImageTranslateTaskResponseBodyData struct {
	// example:
	//
	// [{xxxxxx}]
	ImageData *string `json:"ImageData,omitempty" xml:"ImageData,omitempty"`
}

func (s GetImageTranslateTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetImageTranslateTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetImageTranslateTaskResponseBodyData) GetImageData() *string {
	return s.ImageData
}

func (s *GetImageTranslateTaskResponseBodyData) SetImageData(v string) *GetImageTranslateTaskResponseBodyData {
	s.ImageData = &v
	return s
}

func (s *GetImageTranslateTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
