// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchSqlPreviewResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *FetchResult) *FetchSqlPreviewResultsResponseBody
	GetData() *FetchResult
	SetErrorCode(v string) *FetchSqlPreviewResultsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *FetchSqlPreviewResultsResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *FetchSqlPreviewResultsResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *FetchSqlPreviewResultsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FetchSqlPreviewResultsResponseBody
	GetSuccess() *bool
}

type FetchSqlPreviewResultsResponseBody struct {
	Data *FetchResult `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s FetchSqlPreviewResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FetchSqlPreviewResultsResponseBody) GoString() string {
	return s.String()
}

func (s *FetchSqlPreviewResultsResponseBody) GetData() *FetchResult {
	return s.Data
}

func (s *FetchSqlPreviewResultsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *FetchSqlPreviewResultsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *FetchSqlPreviewResultsResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *FetchSqlPreviewResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FetchSqlPreviewResultsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FetchSqlPreviewResultsResponseBody) SetData(v *FetchResult) *FetchSqlPreviewResultsResponseBody {
	s.Data = v
	return s
}

func (s *FetchSqlPreviewResultsResponseBody) SetErrorCode(v string) *FetchSqlPreviewResultsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *FetchSqlPreviewResultsResponseBody) SetErrorMessage(v string) *FetchSqlPreviewResultsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *FetchSqlPreviewResultsResponseBody) SetHttpCode(v int32) *FetchSqlPreviewResultsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *FetchSqlPreviewResultsResponseBody) SetRequestId(v string) *FetchSqlPreviewResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *FetchSqlPreviewResultsResponseBody) SetSuccess(v bool) *FetchSqlPreviewResultsResponseBody {
	s.Success = &v
	return s
}

func (s *FetchSqlPreviewResultsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
