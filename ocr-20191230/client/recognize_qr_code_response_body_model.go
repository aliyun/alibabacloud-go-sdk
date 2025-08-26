// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeQrCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RecognizeQrCodeResponseBodyData) *RecognizeQrCodeResponseBody
	GetData() *RecognizeQrCodeResponseBodyData
	SetRequestId(v string) *RecognizeQrCodeResponseBody
	GetRequestId() *string
}

type RecognizeQrCodeResponseBody struct {
	Data *RecognizeQrCodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// A53DC437-F883-4968-86D5-EB21FB044692
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeQrCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecognizeQrCodeResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeQrCodeResponseBody) GetData() *RecognizeQrCodeResponseBodyData {
	return s.Data
}

func (s *RecognizeQrCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecognizeQrCodeResponseBody) SetData(v *RecognizeQrCodeResponseBodyData) *RecognizeQrCodeResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeQrCodeResponseBody) SetRequestId(v string) *RecognizeQrCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeQrCodeResponseBody) Validate() error {
	return dara.Validate(s)
}

type RecognizeQrCodeResponseBodyData struct {
	Elements []*RecognizeQrCodeResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s RecognizeQrCodeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecognizeQrCodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeQrCodeResponseBodyData) GetElements() []*RecognizeQrCodeResponseBodyDataElements {
	return s.Elements
}

func (s *RecognizeQrCodeResponseBodyData) SetElements(v []*RecognizeQrCodeResponseBodyDataElements) *RecognizeQrCodeResponseBodyData {
	s.Elements = v
	return s
}

func (s *RecognizeQrCodeResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type RecognizeQrCodeResponseBodyDataElements struct {
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeQrCode/RecognizeQrCode6.jpg
	ImageURL *string                                           `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Results  []*RecognizeQrCodeResponseBodyDataElementsResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// example:
	//
	// img5iGtwVIxQzc4Nqy$L84yHd-1v****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RecognizeQrCodeResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s RecognizeQrCodeResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *RecognizeQrCodeResponseBodyDataElements) GetImageURL() *string {
	return s.ImageURL
}

func (s *RecognizeQrCodeResponseBodyDataElements) GetResults() []*RecognizeQrCodeResponseBodyDataElementsResults {
	return s.Results
}

func (s *RecognizeQrCodeResponseBodyDataElements) GetTaskId() *string {
	return s.TaskId
}

func (s *RecognizeQrCodeResponseBodyDataElements) SetImageURL(v string) *RecognizeQrCodeResponseBodyDataElements {
	s.ImageURL = &v
	return s
}

func (s *RecognizeQrCodeResponseBodyDataElements) SetResults(v []*RecognizeQrCodeResponseBodyDataElementsResults) *RecognizeQrCodeResponseBodyDataElements {
	s.Results = v
	return s
}

func (s *RecognizeQrCodeResponseBodyDataElements) SetTaskId(v string) *RecognizeQrCodeResponseBodyDataElements {
	s.TaskId = &v
	return s
}

func (s *RecognizeQrCodeResponseBodyDataElements) Validate() error {
	return dara.Validate(s)
}

type RecognizeQrCodeResponseBodyDataElementsResults struct {
	// example:
	//
	// qrcode
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// 1
	QrCodesData []*string `json:"QrCodesData,omitempty" xml:"QrCodesData,omitempty" type:"Repeated"`
	// example:
	//
	// 99.91
	Rate *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// example:
	//
	// review
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s RecognizeQrCodeResponseBodyDataElementsResults) String() string {
	return dara.Prettify(s)
}

func (s RecognizeQrCodeResponseBodyDataElementsResults) GoString() string {
	return s.String()
}

func (s *RecognizeQrCodeResponseBodyDataElementsResults) GetLabel() *string {
	return s.Label
}

func (s *RecognizeQrCodeResponseBodyDataElementsResults) GetQrCodesData() []*string {
	return s.QrCodesData
}

func (s *RecognizeQrCodeResponseBodyDataElementsResults) GetRate() *float32 {
	return s.Rate
}

func (s *RecognizeQrCodeResponseBodyDataElementsResults) GetSuggestion() *string {
	return s.Suggestion
}

func (s *RecognizeQrCodeResponseBodyDataElementsResults) SetLabel(v string) *RecognizeQrCodeResponseBodyDataElementsResults {
	s.Label = &v
	return s
}

func (s *RecognizeQrCodeResponseBodyDataElementsResults) SetQrCodesData(v []*string) *RecognizeQrCodeResponseBodyDataElementsResults {
	s.QrCodesData = v
	return s
}

func (s *RecognizeQrCodeResponseBodyDataElementsResults) SetRate(v float32) *RecognizeQrCodeResponseBodyDataElementsResults {
	s.Rate = &v
	return s
}

func (s *RecognizeQrCodeResponseBodyDataElementsResults) SetSuggestion(v string) *RecognizeQrCodeResponseBodyDataElementsResults {
	s.Suggestion = &v
	return s
}

func (s *RecognizeQrCodeResponseBodyDataElementsResults) Validate() error {
	return dara.Validate(s)
}
