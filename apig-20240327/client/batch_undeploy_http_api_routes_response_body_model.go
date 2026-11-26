// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUndeployHttpApiRoutesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchUndeployHttpApiRoutesResponseBody
	GetCode() *string
	SetData(v *BatchUndeployHttpApiRoutesResponseBodyData) *BatchUndeployHttpApiRoutesResponseBody
	GetData() *BatchUndeployHttpApiRoutesResponseBodyData
	SetMessage(v string) *BatchUndeployHttpApiRoutesResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchUndeployHttpApiRoutesResponseBody
	GetRequestId() *string
}

type BatchUndeployHttpApiRoutesResponseBody struct {
	Code      *string                                     `json:"code,omitempty" xml:"code,omitempty"`
	Data      *BatchUndeployHttpApiRoutesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message   *string                                     `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s BatchUndeployHttpApiRoutesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchUndeployHttpApiRoutesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUndeployHttpApiRoutesResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchUndeployHttpApiRoutesResponseBody) GetData() *BatchUndeployHttpApiRoutesResponseBodyData {
	return s.Data
}

func (s *BatchUndeployHttpApiRoutesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchUndeployHttpApiRoutesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchUndeployHttpApiRoutesResponseBody) SetCode(v string) *BatchUndeployHttpApiRoutesResponseBody {
	s.Code = &v
	return s
}

func (s *BatchUndeployHttpApiRoutesResponseBody) SetData(v *BatchUndeployHttpApiRoutesResponseBodyData) *BatchUndeployHttpApiRoutesResponseBody {
	s.Data = v
	return s
}

func (s *BatchUndeployHttpApiRoutesResponseBody) SetMessage(v string) *BatchUndeployHttpApiRoutesResponseBody {
	s.Message = &v
	return s
}

func (s *BatchUndeployHttpApiRoutesResponseBody) SetRequestId(v string) *BatchUndeployHttpApiRoutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUndeployHttpApiRoutesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchUndeployHttpApiRoutesResponseBodyData struct {
	FailedCount  *int32                       `json:"failedCount,omitempty" xml:"failedCount,omitempty"`
	Results      map[string]*DataResultsValue `json:"results,omitempty" xml:"results,omitempty"`
	SkippedCount *int32                       `json:"skippedCount,omitempty" xml:"skippedCount,omitempty"`
	SuccessCount *int32                       `json:"successCount,omitempty" xml:"successCount,omitempty"`
}

func (s BatchUndeployHttpApiRoutesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchUndeployHttpApiRoutesResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchUndeployHttpApiRoutesResponseBodyData) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *BatchUndeployHttpApiRoutesResponseBodyData) GetResults() map[string]*DataResultsValue {
	return s.Results
}

func (s *BatchUndeployHttpApiRoutesResponseBodyData) GetSkippedCount() *int32 {
	return s.SkippedCount
}

func (s *BatchUndeployHttpApiRoutesResponseBodyData) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *BatchUndeployHttpApiRoutesResponseBodyData) SetFailedCount(v int32) *BatchUndeployHttpApiRoutesResponseBodyData {
	s.FailedCount = &v
	return s
}

func (s *BatchUndeployHttpApiRoutesResponseBodyData) SetResults(v map[string]*DataResultsValue) *BatchUndeployHttpApiRoutesResponseBodyData {
	s.Results = v
	return s
}

func (s *BatchUndeployHttpApiRoutesResponseBodyData) SetSkippedCount(v int32) *BatchUndeployHttpApiRoutesResponseBodyData {
	s.SkippedCount = &v
	return s
}

func (s *BatchUndeployHttpApiRoutesResponseBodyData) SetSuccessCount(v int32) *BatchUndeployHttpApiRoutesResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *BatchUndeployHttpApiRoutesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
