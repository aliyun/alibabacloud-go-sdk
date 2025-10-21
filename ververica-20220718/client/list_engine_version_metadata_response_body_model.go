// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEngineVersionMetadataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *EngineVersionMetadataIndex) *ListEngineVersionMetadataResponseBody
	GetData() *EngineVersionMetadataIndex
	SetErrorCode(v string) *ListEngineVersionMetadataResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListEngineVersionMetadataResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *ListEngineVersionMetadataResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *ListEngineVersionMetadataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListEngineVersionMetadataResponseBody
	GetSuccess() *bool
}

type ListEngineVersionMetadataResponseBody struct {
	// 	- If the value of success was true, the engine versions that are supported by Realtime Compute for Apache Flink were returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data *EngineVersionMetadataIndex `json:"data,omitempty" xml:"data,omitempty"`
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
	// The value was fixed to 200.
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

func (s ListEngineVersionMetadataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEngineVersionMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *ListEngineVersionMetadataResponseBody) GetData() *EngineVersionMetadataIndex {
	return s.Data
}

func (s *ListEngineVersionMetadataResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListEngineVersionMetadataResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListEngineVersionMetadataResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListEngineVersionMetadataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEngineVersionMetadataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEngineVersionMetadataResponseBody) SetData(v *EngineVersionMetadataIndex) *ListEngineVersionMetadataResponseBody {
	s.Data = v
	return s
}

func (s *ListEngineVersionMetadataResponseBody) SetErrorCode(v string) *ListEngineVersionMetadataResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListEngineVersionMetadataResponseBody) SetErrorMessage(v string) *ListEngineVersionMetadataResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListEngineVersionMetadataResponseBody) SetHttpCode(v int32) *ListEngineVersionMetadataResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListEngineVersionMetadataResponseBody) SetRequestId(v string) *ListEngineVersionMetadataResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEngineVersionMetadataResponseBody) SetSuccess(v bool) *ListEngineVersionMetadataResponseBody {
	s.Success = &v
	return s
}

func (s *ListEngineVersionMetadataResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
