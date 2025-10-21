// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelSqlPreviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CancelSqlPreviewResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CancelSqlPreviewResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *CancelSqlPreviewResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *CancelSqlPreviewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelSqlPreviewResponseBody
	GetSuccess() *bool
}

type CancelSqlPreviewResponseBody struct {
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

func (s CancelSqlPreviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelSqlPreviewResponseBody) GoString() string {
	return s.String()
}

func (s *CancelSqlPreviewResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CancelSqlPreviewResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CancelSqlPreviewResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *CancelSqlPreviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelSqlPreviewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelSqlPreviewResponseBody) SetErrorCode(v string) *CancelSqlPreviewResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CancelSqlPreviewResponseBody) SetErrorMessage(v string) *CancelSqlPreviewResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CancelSqlPreviewResponseBody) SetHttpCode(v int32) *CancelSqlPreviewResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CancelSqlPreviewResponseBody) SetRequestId(v string) *CancelSqlPreviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelSqlPreviewResponseBody) SetSuccess(v bool) *CancelSqlPreviewResponseBody {
	s.Success = &v
	return s
}

func (s *CancelSqlPreviewResponseBody) Validate() error {
	return dara.Validate(s)
}
