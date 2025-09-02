// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetInstanceLogResponseBody
	GetData() *string
	SetErrorCode(v string) *GetInstanceLogResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetInstanceLogResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetInstanceLogResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetInstanceLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInstanceLogResponseBody
	GetSuccess() *bool
}

type GetInstanceLogResponseBody struct {
	// The content of the logs.
	//
	// example:
	//
	// log_content
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// E6F0DBDD-5AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceLogResponseBody) GetData() *string {
	return s.Data
}

func (s *GetInstanceLogResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetInstanceLogResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetInstanceLogResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInstanceLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInstanceLogResponseBody) SetData(v string) *GetInstanceLogResponseBody {
	s.Data = &v
	return s
}

func (s *GetInstanceLogResponseBody) SetErrorCode(v string) *GetInstanceLogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetInstanceLogResponseBody) SetErrorMessage(v string) *GetInstanceLogResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetInstanceLogResponseBody) SetHttpStatusCode(v int32) *GetInstanceLogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceLogResponseBody) SetRequestId(v string) *GetInstanceLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceLogResponseBody) SetSuccess(v bool) *GetInstanceLogResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceLogResponseBody) Validate() error {
	return dara.Validate(s)
}
