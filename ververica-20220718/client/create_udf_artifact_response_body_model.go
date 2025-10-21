// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUdfArtifactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateUdfArtifactResult) *CreateUdfArtifactResponseBody
	GetData() *CreateUdfArtifactResult
	SetErrorCode(v string) *CreateUdfArtifactResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateUdfArtifactResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *CreateUdfArtifactResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *CreateUdfArtifactResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateUdfArtifactResponseBody
	GetSuccess() *bool
}

type CreateUdfArtifactResponseBody struct {
	// The result of creating an artifact configuration for the UDF.
	Data *CreateUdfArtifactResult `json:"data,omitempty" xml:"data,omitempty"`
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
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateUdfArtifactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUdfArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUdfArtifactResponseBody) GetData() *CreateUdfArtifactResult {
	return s.Data
}

func (s *CreateUdfArtifactResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateUdfArtifactResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateUdfArtifactResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *CreateUdfArtifactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUdfArtifactResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateUdfArtifactResponseBody) SetData(v *CreateUdfArtifactResult) *CreateUdfArtifactResponseBody {
	s.Data = v
	return s
}

func (s *CreateUdfArtifactResponseBody) SetErrorCode(v string) *CreateUdfArtifactResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateUdfArtifactResponseBody) SetErrorMessage(v string) *CreateUdfArtifactResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateUdfArtifactResponseBody) SetHttpCode(v int32) *CreateUdfArtifactResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateUdfArtifactResponseBody) SetRequestId(v string) *CreateUdfArtifactResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUdfArtifactResponseBody) SetSuccess(v bool) *CreateUdfArtifactResponseBody {
	s.Success = &v
	return s
}

func (s *CreateUdfArtifactResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
