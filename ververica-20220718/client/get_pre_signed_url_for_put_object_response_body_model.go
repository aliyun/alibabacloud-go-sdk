// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPreSignedUrlForPutObjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetPreSignedUrlForObjectResult) *GetPreSignedUrlForPutObjectResponseBody
	GetData() *GetPreSignedUrlForObjectResult
	SetErrorCode(v string) *GetPreSignedUrlForPutObjectResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetPreSignedUrlForPutObjectResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetPreSignedUrlForPutObjectResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetPreSignedUrlForPutObjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPreSignedUrlForPutObjectResponseBody
	GetSuccess() *bool
}

type GetPreSignedUrlForPutObjectResponseBody struct {
	Data *GetPreSignedUrlForObjectResult `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 990301
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// 文件不存在
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// 1234567890
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetPreSignedUrlForPutObjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPreSignedUrlForPutObjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetPreSignedUrlForPutObjectResponseBody) GetData() *GetPreSignedUrlForObjectResult {
	return s.Data
}

func (s *GetPreSignedUrlForPutObjectResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetPreSignedUrlForPutObjectResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetPreSignedUrlForPutObjectResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetPreSignedUrlForPutObjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPreSignedUrlForPutObjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPreSignedUrlForPutObjectResponseBody) SetData(v *GetPreSignedUrlForObjectResult) *GetPreSignedUrlForPutObjectResponseBody {
	s.Data = v
	return s
}

func (s *GetPreSignedUrlForPutObjectResponseBody) SetErrorCode(v string) *GetPreSignedUrlForPutObjectResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPreSignedUrlForPutObjectResponseBody) SetErrorMessage(v string) *GetPreSignedUrlForPutObjectResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPreSignedUrlForPutObjectResponseBody) SetHttpCode(v int32) *GetPreSignedUrlForPutObjectResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetPreSignedUrlForPutObjectResponseBody) SetRequestId(v string) *GetPreSignedUrlForPutObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPreSignedUrlForPutObjectResponseBody) SetSuccess(v bool) *GetPreSignedUrlForPutObjectResponseBody {
	s.Success = &v
	return s
}

func (s *GetPreSignedUrlForPutObjectResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
