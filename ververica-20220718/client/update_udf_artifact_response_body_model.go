// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUdfArtifactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateUdfArtifactResult) *UpdateUdfArtifactResponseBody
	GetData() *UpdateUdfArtifactResult
	SetErrorCode(v string) *UpdateUdfArtifactResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateUdfArtifactResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *UpdateUdfArtifactResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *UpdateUdfArtifactResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateUdfArtifactResponseBody
	GetSuccess() *bool
}

type UpdateUdfArtifactResponseBody struct {
	// The result of updating the JAR file of the UDF.
	Data *UpdateUdfArtifactResult `json:"data,omitempty" xml:"data,omitempty"`
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
	// CBC799F0-ABCD-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateUdfArtifactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUdfArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUdfArtifactResponseBody) GetData() *UpdateUdfArtifactResult {
	return s.Data
}

func (s *UpdateUdfArtifactResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateUdfArtifactResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateUdfArtifactResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *UpdateUdfArtifactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUdfArtifactResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateUdfArtifactResponseBody) SetData(v *UpdateUdfArtifactResult) *UpdateUdfArtifactResponseBody {
	s.Data = v
	return s
}

func (s *UpdateUdfArtifactResponseBody) SetErrorCode(v string) *UpdateUdfArtifactResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateUdfArtifactResponseBody) SetErrorMessage(v string) *UpdateUdfArtifactResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateUdfArtifactResponseBody) SetHttpCode(v int32) *UpdateUdfArtifactResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateUdfArtifactResponseBody) SetRequestId(v string) *UpdateUdfArtifactResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUdfArtifactResponseBody) SetSuccess(v bool) *UpdateUdfArtifactResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateUdfArtifactResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
