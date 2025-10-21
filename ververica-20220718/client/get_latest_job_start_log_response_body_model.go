// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLatestJobStartLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetLatestJobStartLogResponseBody
	GetData() *string
	SetErrorCode(v string) *GetLatestJobStartLogResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetLatestJobStartLogResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetLatestJobStartLogResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetLatestJobStartLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLatestJobStartLogResponseBody
	GetSuccess() *bool
}

type GetLatestJobStartLogResponseBody struct {
	// If the value of success was false, the latest logs of the deployment were returned. If the value of success was true, a null value was returned.
	//
	// example:
	//
	// "[main] INFO  org.apache.flink.runtime.entrypoint.ClusterEntrypoint        [] - --------------------------------------------------------------------------------\\n2024-05-22 11:46:39,871 [main] INFO  org.apache.flink.runtime.entrypoint.ClusterEntrypoint"
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// If the value of success was false, an error code was returned. If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// If the value of success was false, an error message was returned. If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The status code returned. The value was fixed to 200.
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

func (s GetLatestJobStartLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLatestJobStartLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetLatestJobStartLogResponseBody) GetData() *string {
	return s.Data
}

func (s *GetLatestJobStartLogResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetLatestJobStartLogResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetLatestJobStartLogResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetLatestJobStartLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLatestJobStartLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLatestJobStartLogResponseBody) SetData(v string) *GetLatestJobStartLogResponseBody {
	s.Data = &v
	return s
}

func (s *GetLatestJobStartLogResponseBody) SetErrorCode(v string) *GetLatestJobStartLogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetLatestJobStartLogResponseBody) SetErrorMessage(v string) *GetLatestJobStartLogResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetLatestJobStartLogResponseBody) SetHttpCode(v int32) *GetLatestJobStartLogResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetLatestJobStartLogResponseBody) SetRequestId(v string) *GetLatestJobStartLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLatestJobStartLogResponseBody) SetSuccess(v bool) *GetLatestJobStartLogResponseBody {
	s.Success = &v
	return s
}

func (s *GetLatestJobStartLogResponseBody) Validate() error {
	return dara.Validate(s)
}
