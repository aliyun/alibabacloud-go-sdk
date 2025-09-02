// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgDesensPlanDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DsgDesensPlanDeleteResponseBody
	GetData() *bool
	SetErrorCode(v string) *DsgDesensPlanDeleteResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgDesensPlanDeleteResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgDesensPlanDeleteResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DsgDesensPlanDeleteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgDesensPlanDeleteResponseBody
	GetSuccess() *bool
}

type DsgDesensPlanDeleteResponseBody struct {
	// The operation result. Valid values:
	//
	// 	- true: The operation is successful.
	//
	// 	- false: The operation fails.
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

func (s DsgDesensPlanDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgDesensPlanDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *DsgDesensPlanDeleteResponseBody) GetData() *bool {
	return s.Data
}

func (s *DsgDesensPlanDeleteResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgDesensPlanDeleteResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgDesensPlanDeleteResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgDesensPlanDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgDesensPlanDeleteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgDesensPlanDeleteResponseBody) SetData(v bool) *DsgDesensPlanDeleteResponseBody {
	s.Data = &v
	return s
}

func (s *DsgDesensPlanDeleteResponseBody) SetErrorCode(v string) *DsgDesensPlanDeleteResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgDesensPlanDeleteResponseBody) SetErrorMessage(v string) *DsgDesensPlanDeleteResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgDesensPlanDeleteResponseBody) SetHttpStatusCode(v int32) *DsgDesensPlanDeleteResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgDesensPlanDeleteResponseBody) SetRequestId(v string) *DsgDesensPlanDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgDesensPlanDeleteResponseBody) SetSuccess(v bool) *DsgDesensPlanDeleteResponseBody {
	s.Success = &v
	return s
}

func (s *DsgDesensPlanDeleteResponseBody) Validate() error {
	return dara.Validate(s)
}
