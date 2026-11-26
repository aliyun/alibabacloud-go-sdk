// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteHttpApiRoutesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchDeleteHttpApiRoutesResponseBody
	GetCode() *string
	SetData(v *BatchDeleteHttpApiRoutesResponseBodyData) *BatchDeleteHttpApiRoutesResponseBody
	GetData() *BatchDeleteHttpApiRoutesResponseBodyData
	SetMessage(v string) *BatchDeleteHttpApiRoutesResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchDeleteHttpApiRoutesResponseBody
	GetRequestId() *string
}

type BatchDeleteHttpApiRoutesResponseBody struct {
	Code      *string                                   `json:"code,omitempty" xml:"code,omitempty"`
	Data      *BatchDeleteHttpApiRoutesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message   *string                                   `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s BatchDeleteHttpApiRoutesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteHttpApiRoutesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteHttpApiRoutesResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchDeleteHttpApiRoutesResponseBody) GetData() *BatchDeleteHttpApiRoutesResponseBodyData {
	return s.Data
}

func (s *BatchDeleteHttpApiRoutesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchDeleteHttpApiRoutesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeleteHttpApiRoutesResponseBody) SetCode(v string) *BatchDeleteHttpApiRoutesResponseBody {
	s.Code = &v
	return s
}

func (s *BatchDeleteHttpApiRoutesResponseBody) SetData(v *BatchDeleteHttpApiRoutesResponseBodyData) *BatchDeleteHttpApiRoutesResponseBody {
	s.Data = v
	return s
}

func (s *BatchDeleteHttpApiRoutesResponseBody) SetMessage(v string) *BatchDeleteHttpApiRoutesResponseBody {
	s.Message = &v
	return s
}

func (s *BatchDeleteHttpApiRoutesResponseBody) SetRequestId(v string) *BatchDeleteHttpApiRoutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteHttpApiRoutesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchDeleteHttpApiRoutesResponseBodyData struct {
	FailedCount  *int32                       `json:"failedCount,omitempty" xml:"failedCount,omitempty"`
	Results      map[string]*DataResultsValue `json:"results,omitempty" xml:"results,omitempty"`
	SuccessCount *int32                       `json:"successCount,omitempty" xml:"successCount,omitempty"`
}

func (s BatchDeleteHttpApiRoutesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteHttpApiRoutesResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchDeleteHttpApiRoutesResponseBodyData) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *BatchDeleteHttpApiRoutesResponseBodyData) GetResults() map[string]*DataResultsValue {
	return s.Results
}

func (s *BatchDeleteHttpApiRoutesResponseBodyData) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *BatchDeleteHttpApiRoutesResponseBodyData) SetFailedCount(v int32) *BatchDeleteHttpApiRoutesResponseBodyData {
	s.FailedCount = &v
	return s
}

func (s *BatchDeleteHttpApiRoutesResponseBodyData) SetResults(v map[string]*DataResultsValue) *BatchDeleteHttpApiRoutesResponseBodyData {
	s.Results = v
	return s
}

func (s *BatchDeleteHttpApiRoutesResponseBodyData) SetSuccessCount(v int32) *BatchDeleteHttpApiRoutesResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *BatchDeleteHttpApiRoutesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
