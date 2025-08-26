// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeVINCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RecognizeVINCodeResponseBodyData) *RecognizeVINCodeResponseBody
	GetData() *RecognizeVINCodeResponseBodyData
	SetRequestId(v string) *RecognizeVINCodeResponseBody
	GetRequestId() *string
}

type RecognizeVINCodeResponseBody struct {
	Data *RecognizeVINCodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 911FC8CF-CC27-477E-BE3B-7ED77DF4DFE0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeVINCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVINCodeResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVINCodeResponseBody) GetData() *RecognizeVINCodeResponseBodyData {
	return s.Data
}

func (s *RecognizeVINCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecognizeVINCodeResponseBody) SetData(v *RecognizeVINCodeResponseBodyData) *RecognizeVINCodeResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeVINCodeResponseBody) SetRequestId(v string) *RecognizeVINCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeVINCodeResponseBody) Validate() error {
	return dara.Validate(s)
}

type RecognizeVINCodeResponseBodyData struct {
	// example:
	//
	// LVBB2FAF777999888
	VinCode *string `json:"VinCode,omitempty" xml:"VinCode,omitempty"`
}

func (s RecognizeVINCodeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVINCodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeVINCodeResponseBodyData) GetVinCode() *string {
	return s.VinCode
}

func (s *RecognizeVINCodeResponseBodyData) SetVinCode(v string) *RecognizeVINCodeResponseBodyData {
	s.VinCode = &v
	return s
}

func (s *RecognizeVINCodeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
