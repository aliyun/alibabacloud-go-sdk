// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgWhiteListDeleteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DsgWhiteListDeleteListResponseBody
	GetData() *bool
	SetErrorCode(v string) *DsgWhiteListDeleteListResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgWhiteListDeleteListResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgWhiteListDeleteListResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DsgWhiteListDeleteListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgWhiteListDeleteListResponseBody
	GetSuccess() *bool
}

type DsgWhiteListDeleteListResponseBody struct {
	// The operation result. Valid values:
	//
	// 	- true: The operation is successful.
	//
	// 	- false: The operation failed.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// 1029030003
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// param error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// 102400001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DsgWhiteListDeleteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgWhiteListDeleteListResponseBody) GoString() string {
	return s.String()
}

func (s *DsgWhiteListDeleteListResponseBody) GetData() *bool {
	return s.Data
}

func (s *DsgWhiteListDeleteListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgWhiteListDeleteListResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgWhiteListDeleteListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgWhiteListDeleteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgWhiteListDeleteListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgWhiteListDeleteListResponseBody) SetData(v bool) *DsgWhiteListDeleteListResponseBody {
	s.Data = &v
	return s
}

func (s *DsgWhiteListDeleteListResponseBody) SetErrorCode(v string) *DsgWhiteListDeleteListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgWhiteListDeleteListResponseBody) SetErrorMessage(v string) *DsgWhiteListDeleteListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgWhiteListDeleteListResponseBody) SetHttpStatusCode(v int32) *DsgWhiteListDeleteListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgWhiteListDeleteListResponseBody) SetRequestId(v string) *DsgWhiteListDeleteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgWhiteListDeleteListResponseBody) SetSuccess(v bool) *DsgWhiteListDeleteListResponseBody {
	s.Success = &v
	return s
}

func (s *DsgWhiteListDeleteListResponseBody) Validate() error {
	return dara.Validate(s)
}
