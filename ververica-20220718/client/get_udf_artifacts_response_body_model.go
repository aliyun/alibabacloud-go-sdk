// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUdfArtifactsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*UdfArtifact) *GetUdfArtifactsResponseBody
	GetData() []*UdfArtifact
	SetErrorCode(v string) *GetUdfArtifactsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetUdfArtifactsResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetUdfArtifactsResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetUdfArtifactsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUdfArtifactsResponseBody
	GetSuccess() *bool
}

type GetUdfArtifactsResponseBody struct {
	// If the value of success was true, the details of the JAR or Python file that corresponds to the UDF were returned. If the value of success was false, a null value was returned.
	Data []*UdfArtifact `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s GetUdfArtifactsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUdfArtifactsResponseBody) GoString() string {
	return s.String()
}

func (s *GetUdfArtifactsResponseBody) GetData() []*UdfArtifact {
	return s.Data
}

func (s *GetUdfArtifactsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetUdfArtifactsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetUdfArtifactsResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetUdfArtifactsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUdfArtifactsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUdfArtifactsResponseBody) SetData(v []*UdfArtifact) *GetUdfArtifactsResponseBody {
	s.Data = v
	return s
}

func (s *GetUdfArtifactsResponseBody) SetErrorCode(v string) *GetUdfArtifactsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUdfArtifactsResponseBody) SetErrorMessage(v string) *GetUdfArtifactsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetUdfArtifactsResponseBody) SetHttpCode(v int32) *GetUdfArtifactsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetUdfArtifactsResponseBody) SetRequestId(v string) *GetUdfArtifactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUdfArtifactsResponseBody) SetSuccess(v bool) *GetUdfArtifactsResponseBody {
	s.Success = &v
	return s
}

func (s *GetUdfArtifactsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
