// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TranslateSearchResponseBody
	GetCode() *string
	SetData(v *TranslateSearchResponseBodyData) *TranslateSearchResponseBody
	GetData() *TranslateSearchResponseBodyData
	SetMessage(v string) *TranslateSearchResponseBody
	GetMessage() *string
	SetRequestId(v string) *TranslateSearchResponseBody
	GetRequestId() *string
}

type TranslateSearchResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *TranslateSearchResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TranslateSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TranslateSearchResponseBody) GoString() string {
	return s.String()
}

func (s *TranslateSearchResponseBody) GetCode() *string {
	return s.Code
}

func (s *TranslateSearchResponseBody) GetData() *TranslateSearchResponseBodyData {
	return s.Data
}

func (s *TranslateSearchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TranslateSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TranslateSearchResponseBody) SetCode(v string) *TranslateSearchResponseBody {
	s.Code = &v
	return s
}

func (s *TranslateSearchResponseBody) SetData(v *TranslateSearchResponseBodyData) *TranslateSearchResponseBody {
	s.Data = v
	return s
}

func (s *TranslateSearchResponseBody) SetMessage(v string) *TranslateSearchResponseBody {
	s.Message = &v
	return s
}

func (s *TranslateSearchResponseBody) SetRequestId(v string) *TranslateSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *TranslateSearchResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TranslateSearchResponseBodyData struct {
	Translated *string `json:"Translated,omitempty" xml:"Translated,omitempty"`
}

func (s TranslateSearchResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TranslateSearchResponseBodyData) GoString() string {
	return s.String()
}

func (s *TranslateSearchResponseBodyData) GetTranslated() *string {
	return s.Translated
}

func (s *TranslateSearchResponseBodyData) SetTranslated(v string) *TranslateSearchResponseBodyData {
	s.Translated = &v
	return s
}

func (s *TranslateSearchResponseBodyData) Validate() error {
	return dara.Validate(s)
}
