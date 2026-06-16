// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLanguageDetectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *LanguageDetectResponseBody
	GetCode() *string
	SetData(v *LanguageDetectResponseBodyData) *LanguageDetectResponseBody
	GetData() *LanguageDetectResponseBodyData
	SetMessage(v string) *LanguageDetectResponseBody
	GetMessage() *string
	SetRequestId(v string) *LanguageDetectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *LanguageDetectResponseBody
	GetSuccess() *bool
}

type LanguageDetectResponseBody struct {
	// The response code. A value of 200 indicates success. For other response codes, refer to the error code documentation.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The language detection result data, including the detected language and usage information.
	Data *LanguageDetectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. Returns "Success" for successful calls. Returns a specific error message for failed calls, such as "The parameters contain sensitive information. Try other input.".
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which uniquely identifies the request.
	//
	// example:
	//
	// 42542C6C-F2A4-1B2B-8EFF-130C8FD06F54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s LanguageDetectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LanguageDetectResponseBody) GoString() string {
	return s.String()
}

func (s *LanguageDetectResponseBody) GetCode() *string {
	return s.Code
}

func (s *LanguageDetectResponseBody) GetData() *LanguageDetectResponseBodyData {
	return s.Data
}

func (s *LanguageDetectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *LanguageDetectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LanguageDetectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *LanguageDetectResponseBody) SetCode(v string) *LanguageDetectResponseBody {
	s.Code = &v
	return s
}

func (s *LanguageDetectResponseBody) SetData(v *LanguageDetectResponseBodyData) *LanguageDetectResponseBody {
	s.Data = v
	return s
}

func (s *LanguageDetectResponseBody) SetMessage(v string) *LanguageDetectResponseBody {
	s.Message = &v
	return s
}

func (s *LanguageDetectResponseBody) SetRequestId(v string) *LanguageDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *LanguageDetectResponseBody) SetSuccess(v bool) *LanguageDetectResponseBody {
	s.Success = &v
	return s
}

func (s *LanguageDetectResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LanguageDetectResponseBodyData struct {
	// The detected language code.
	//
	// example:
	//
	// zh
	DetectedLanguage *string `json:"DetectedLanguage,omitempty" xml:"DetectedLanguage,omitempty"`
	// The usage information, including the number of input characters.
	//
	// example:
	//
	// {"InputCharacterCount":4}
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s LanguageDetectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s LanguageDetectResponseBodyData) GoString() string {
	return s.String()
}

func (s *LanguageDetectResponseBodyData) GetDetectedLanguage() *string {
	return s.DetectedLanguage
}

func (s *LanguageDetectResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *LanguageDetectResponseBodyData) SetDetectedLanguage(v string) *LanguageDetectResponseBodyData {
	s.DetectedLanguage = &v
	return s
}

func (s *LanguageDetectResponseBodyData) SetUsageMap(v map[string]*int64) *LanguageDetectResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *LanguageDetectResponseBodyData) Validate() error {
	return dara.Validate(s)
}
