// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLineageInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *LineageInfo) *GetLineageInfoResponseBody
	GetData() *LineageInfo
	SetErrorCode(v string) *GetLineageInfoResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetLineageInfoResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetLineageInfoResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetLineageInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLineageInfoResponseBody
	GetSuccess() *bool
}

type GetLineageInfoResponseBody struct {
	// The lineage information.
	Data *LineageInfo `json:"data,omitempty" xml:"data,omitempty"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The status code returned. The value was fixed to 200. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetLineageInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLineageInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetLineageInfoResponseBody) GetData() *LineageInfo {
	return s.Data
}

func (s *GetLineageInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetLineageInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetLineageInfoResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetLineageInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLineageInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLineageInfoResponseBody) SetData(v *LineageInfo) *GetLineageInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetLineageInfoResponseBody) SetErrorCode(v string) *GetLineageInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetLineageInfoResponseBody) SetErrorMessage(v string) *GetLineageInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetLineageInfoResponseBody) SetHttpCode(v int32) *GetLineageInfoResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetLineageInfoResponseBody) SetRequestId(v string) *GetLineageInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLineageInfoResponseBody) SetSuccess(v bool) *GetLineageInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetLineageInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
