// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextCorrectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TextCorrectResponseBody
	GetCode() *string
	SetData(v *TextCorrectResponseBodyData) *TextCorrectResponseBody
	GetData() *TextCorrectResponseBodyData
	SetMessage(v string) *TextCorrectResponseBody
	GetMessage() *string
	SetRequestId(v string) *TextCorrectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TextCorrectResponseBody
	GetSuccess() *bool
}

type TextCorrectResponseBody struct {
	// Response code. Returns "success" during normal calls.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Intelligent error correction result data.
	Data *TextCorrectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Error message. Returns "Success" during normal calls. Returns specific error information during exceptions, such as "The parameters contain sensitive information. Please try a different input."
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID, used to identify a unique request call.
	//
	// example:
	//
	// 32882AD0-50D1-1D90-A221-3987325EC03E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful. true indicates success, false indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TextCorrectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TextCorrectResponseBody) GoString() string {
	return s.String()
}

func (s *TextCorrectResponseBody) GetCode() *string {
	return s.Code
}

func (s *TextCorrectResponseBody) GetData() *TextCorrectResponseBodyData {
	return s.Data
}

func (s *TextCorrectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TextCorrectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TextCorrectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TextCorrectResponseBody) SetCode(v string) *TextCorrectResponseBody {
	s.Code = &v
	return s
}

func (s *TextCorrectResponseBody) SetData(v *TextCorrectResponseBodyData) *TextCorrectResponseBody {
	s.Data = v
	return s
}

func (s *TextCorrectResponseBody) SetMessage(v string) *TextCorrectResponseBody {
	s.Message = &v
	return s
}

func (s *TextCorrectResponseBody) SetRequestId(v string) *TextCorrectResponseBody {
	s.RequestId = &v
	return s
}

func (s *TextCorrectResponseBody) SetSuccess(v bool) *TextCorrectResponseBody {
	s.Success = &v
	return s
}

func (s *TextCorrectResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TextCorrectResponseBodyData struct {
	// The corrected text.
	//
	// example:
	//
	// Empfehlung
	CorrectedText *string `json:"CorrectedText,omitempty" xml:"CorrectedText,omitempty"`
	// Usage information, including the number of input characters.
	//
	// example:
	//
	// {"InputCharacterCount":9}
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s TextCorrectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TextCorrectResponseBodyData) GoString() string {
	return s.String()
}

func (s *TextCorrectResponseBodyData) GetCorrectedText() *string {
	return s.CorrectedText
}

func (s *TextCorrectResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *TextCorrectResponseBodyData) SetCorrectedText(v string) *TextCorrectResponseBodyData {
	s.CorrectedText = &v
	return s
}

func (s *TextCorrectResponseBodyData) SetUsageMap(v map[string]*int64) *TextCorrectResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *TextCorrectResponseBodyData) Validate() error {
	return dara.Validate(s)
}
