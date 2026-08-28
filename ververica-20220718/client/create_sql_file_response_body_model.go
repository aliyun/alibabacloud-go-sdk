// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSqlFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SqlFile) *CreateSqlFileResponseBody
	GetData() *SqlFile
	SetErrorCode(v string) *CreateSqlFileResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateSqlFileResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *CreateSqlFileResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *CreateSqlFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSqlFileResponseBody
	GetSuccess() *bool
}

type CreateSqlFileResponseBody struct {
	// The variable configuration settings.
	//
	// example:
	//
	// "[main] INFO  org.apache.flink.runtime.entrypoint.ClusterEntrypoint        [] - --------------------------------------------------------------------------------\\n2024-05-22 11:46:39,871 [main] INFO  org.apache.flink.runtime.entrypoint.ClusterEntrypoint"
	Data *SqlFile `json:"data,omitempty" xml:"data,omitempty"`
	// The business error code returned when success is false. This value is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The business error message returned when success is false. This value is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The business status code, which is always 200. Use success to determine whether the request was successful.
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
	// Indicates whether the business request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateSqlFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSqlFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSqlFileResponseBody) GetData() *SqlFile {
	return s.Data
}

func (s *CreateSqlFileResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateSqlFileResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateSqlFileResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *CreateSqlFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSqlFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSqlFileResponseBody) SetData(v *SqlFile) *CreateSqlFileResponseBody {
	s.Data = v
	return s
}

func (s *CreateSqlFileResponseBody) SetErrorCode(v string) *CreateSqlFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateSqlFileResponseBody) SetErrorMessage(v string) *CreateSqlFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateSqlFileResponseBody) SetHttpCode(v int32) *CreateSqlFileResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateSqlFileResponseBody) SetRequestId(v string) *CreateSqlFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSqlFileResponseBody) SetSuccess(v bool) *CreateSqlFileResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSqlFileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
