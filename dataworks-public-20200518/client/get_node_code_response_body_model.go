// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetNodeCodeResponseBody
	GetData() *string
	SetErrorCode(v string) *GetNodeCodeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetNodeCodeResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetNodeCodeResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetNodeCodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetNodeCodeResponseBody
	GetSuccess() *bool
}

type GetNodeCodeResponseBody struct {
	// The code of the node.
	//
	// example:
	//
	// select a;
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// E6F0DBDD-5AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNodeCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNodeCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeCodeResponseBody) GetData() *string {
	return s.Data
}

func (s *GetNodeCodeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetNodeCodeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetNodeCodeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetNodeCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNodeCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetNodeCodeResponseBody) SetData(v string) *GetNodeCodeResponseBody {
	s.Data = &v
	return s
}

func (s *GetNodeCodeResponseBody) SetErrorCode(v string) *GetNodeCodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetNodeCodeResponseBody) SetErrorMessage(v string) *GetNodeCodeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetNodeCodeResponseBody) SetHttpStatusCode(v int32) *GetNodeCodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetNodeCodeResponseBody) SetRequestId(v string) *GetNodeCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeCodeResponseBody) SetSuccess(v bool) *GetNodeCodeResponseBody {
	s.Success = &v
	return s
}

func (s *GetNodeCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
