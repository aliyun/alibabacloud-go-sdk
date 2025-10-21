// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSqlPreviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SubmitPreviewResult) *SubmitSqlPreviewResponseBody
	GetData() *SubmitPreviewResult
	SetErrorCode(v string) *SubmitSqlPreviewResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SubmitSqlPreviewResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *SubmitSqlPreviewResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *SubmitSqlPreviewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitSqlPreviewResponseBody
	GetSuccess() *bool
}

type SubmitSqlPreviewResponseBody struct {
	// example:
	//
	// "[main] INFO  org.apache.flink.runtime.entrypoint.ClusterEntrypoint        [] - --------------------------------------------------------------------------------\\n2024-05-22 11:46:39,871 [main] INFO  org.apache.flink.runtime.entrypoint.ClusterEntrypoint"
	Data *SubmitPreviewResult `json:"data,omitempty" xml:"data,omitempty"`
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
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SubmitSqlPreviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSqlPreviewResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSqlPreviewResponseBody) GetData() *SubmitPreviewResult {
	return s.Data
}

func (s *SubmitSqlPreviewResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SubmitSqlPreviewResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SubmitSqlPreviewResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *SubmitSqlPreviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSqlPreviewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitSqlPreviewResponseBody) SetData(v *SubmitPreviewResult) *SubmitSqlPreviewResponseBody {
	s.Data = v
	return s
}

func (s *SubmitSqlPreviewResponseBody) SetErrorCode(v string) *SubmitSqlPreviewResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitSqlPreviewResponseBody) SetErrorMessage(v string) *SubmitSqlPreviewResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SubmitSqlPreviewResponseBody) SetHttpCode(v int32) *SubmitSqlPreviewResponseBody {
	s.HttpCode = &v
	return s
}

func (s *SubmitSqlPreviewResponseBody) SetRequestId(v string) *SubmitSqlPreviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSqlPreviewResponseBody) SetSuccess(v bool) *SubmitSqlPreviewResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitSqlPreviewResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
