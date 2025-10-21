// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotUpdateJobResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *HotUpdateJobResult) *GetHotUpdateJobResultResponseBody
	GetData() *HotUpdateJobResult
	SetErrorCode(v string) *GetHotUpdateJobResultResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetHotUpdateJobResultResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetHotUpdateJobResultResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetHotUpdateJobResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHotUpdateJobResultResponseBody
	GetSuccess() *bool
}

type GetHotUpdateJobResultResponseBody struct {
	Data *HotUpdateJobResult `json:"data,omitempty" xml:"data,omitempty"`
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
	// CBC799F0-ABCF-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetHotUpdateJobResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotUpdateJobResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotUpdateJobResultResponseBody) GetData() *HotUpdateJobResult {
	return s.Data
}

func (s *GetHotUpdateJobResultResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetHotUpdateJobResultResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetHotUpdateJobResultResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetHotUpdateJobResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotUpdateJobResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHotUpdateJobResultResponseBody) SetData(v *HotUpdateJobResult) *GetHotUpdateJobResultResponseBody {
	s.Data = v
	return s
}

func (s *GetHotUpdateJobResultResponseBody) SetErrorCode(v string) *GetHotUpdateJobResultResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetHotUpdateJobResultResponseBody) SetErrorMessage(v string) *GetHotUpdateJobResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetHotUpdateJobResultResponseBody) SetHttpCode(v int32) *GetHotUpdateJobResultResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetHotUpdateJobResultResponseBody) SetRequestId(v string) *GetHotUpdateJobResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotUpdateJobResultResponseBody) SetSuccess(v bool) *GetHotUpdateJobResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotUpdateJobResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
