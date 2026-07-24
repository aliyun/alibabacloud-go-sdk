// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSqlFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SqlFile) *DeleteSqlFileResponseBody
	GetData() *SqlFile
	SetErrorCode(v string) *DeleteSqlFileResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteSqlFileResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *DeleteSqlFileResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *DeleteSqlFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSqlFileResponseBody
	GetSuccess() *bool
}

type DeleteSqlFileResponseBody struct {
	// The basic information about the deleted file returned after a successful deletion. Only metadata is included. The content field is not returned. This field is valid when success is set to true.
	//
	// example:
	//
	// { "jobs": [ { "jid": "4df35f8e54554b23bf7dcd38a151****", "name": "69d001d5-419a-4bfc-9c2e-849cacd3****", "state": "RUNNING", "start-time": 1659154942068, "end-time": -1, "duration": 188161756, "last-modification": 1659154968305, "tasks": { "total": 2, "created": 0, "scheduled": 0, "deploying": 0, "running": 2, "finished": 0, "canceling": 0, "canceled": 0, "failed": 0, "reconciling": 0, "initializing": 0 } } ] }
	Data *SqlFile `json:"data,omitempty" xml:"data,omitempty"`
	// The error code returned when success is false. This value is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message returned when success is false. This value is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The HTTP status code, which is always 200. Use the success field to determine whether the request was successful.
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

func (s DeleteSqlFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSqlFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSqlFileResponseBody) GetData() *SqlFile {
	return s.Data
}

func (s *DeleteSqlFileResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteSqlFileResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteSqlFileResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *DeleteSqlFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSqlFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSqlFileResponseBody) SetData(v *SqlFile) *DeleteSqlFileResponseBody {
	s.Data = v
	return s
}

func (s *DeleteSqlFileResponseBody) SetErrorCode(v string) *DeleteSqlFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteSqlFileResponseBody) SetErrorMessage(v string) *DeleteSqlFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteSqlFileResponseBody) SetHttpCode(v int32) *DeleteSqlFileResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteSqlFileResponseBody) SetRequestId(v string) *DeleteSqlFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSqlFileResponseBody) SetSuccess(v bool) *DeleteSqlFileResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSqlFileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
