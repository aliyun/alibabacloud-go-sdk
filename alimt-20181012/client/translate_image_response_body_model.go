// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *TranslateImageResponseBody
	GetCode() *int32
	SetData(v *TranslateImageResponseBodyData) *TranslateImageResponseBody
	GetData() *TranslateImageResponseBodyData
	SetMessage(v string) *TranslateImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *TranslateImageResponseBody
	GetRequestId() *string
}

type TranslateImageResponseBody struct {
	// example:
	//
	// 200
	Code *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *TranslateImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Error Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D774D33D-F1CB-5A2C-A787-E0A2179239CE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TranslateImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TranslateImageResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateImageResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *TranslateImageResponseBody) GetData() *TranslateImageResponseBodyData {
	return s.Data
}

func (s *TranslateImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TranslateImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TranslateImageResponseBody) SetCode(v int32) *TranslateImageResponseBody {
	s.Code = &v
	return s
}

func (s *TranslateImageResponseBody) SetData(v *TranslateImageResponseBodyData) *TranslateImageResponseBody {
	s.Data = v
	return s
}

func (s *TranslateImageResponseBody) SetMessage(v string) *TranslateImageResponseBody {
	s.Message = &v
	return s
}

func (s *TranslateImageResponseBody) SetRequestId(v string) *TranslateImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *TranslateImageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TranslateImageResponseBodyData struct {
	// example:
	//
	// https://example.com/example.jpg
	FinalImageUrl *string `json:"FinalImageUrl,omitempty" xml:"FinalImageUrl,omitempty"`
	// example:
	//
	// https://example.com/example.jpg
	InPaintingUrl *string `json:"InPaintingUrl,omitempty" xml:"InPaintingUrl,omitempty"`
	// example:
	//
	// Editor Template Json String
	TemplateJson *string `json:"TemplateJson,omitempty" xml:"TemplateJson,omitempty"`
}

func (s TranslateImageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TranslateImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *TranslateImageResponseBodyData) GetFinalImageUrl() *string {
	return s.FinalImageUrl
}

func (s *TranslateImageResponseBodyData) GetInPaintingUrl() *string {
	return s.InPaintingUrl
}

func (s *TranslateImageResponseBodyData) GetTemplateJson() *string {
	return s.TemplateJson
}

func (s *TranslateImageResponseBodyData) SetFinalImageUrl(v string) *TranslateImageResponseBodyData {
	s.FinalImageUrl = &v
	return s
}

func (s *TranslateImageResponseBodyData) SetInPaintingUrl(v string) *TranslateImageResponseBodyData {
	s.InPaintingUrl = &v
	return s
}

func (s *TranslateImageResponseBodyData) SetTemplateJson(v string) *TranslateImageResponseBodyData {
	s.TemplateJson = &v
	return s
}

func (s *TranslateImageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
