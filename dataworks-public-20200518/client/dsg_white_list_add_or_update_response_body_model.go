// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgWhiteListAddOrUpdateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DsgWhiteListAddOrUpdateResponseBody
	GetData() *bool
	SetErrorCode(v string) *DsgWhiteListAddOrUpdateResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgWhiteListAddOrUpdateResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgWhiteListAddOrUpdateResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DsgWhiteListAddOrUpdateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgWhiteListAddOrUpdateResponseBody
	GetSuccess() *bool
}

type DsgWhiteListAddOrUpdateResponseBody struct {
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

func (s DsgWhiteListAddOrUpdateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgWhiteListAddOrUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *DsgWhiteListAddOrUpdateResponseBody) GetData() *bool {
	return s.Data
}

func (s *DsgWhiteListAddOrUpdateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgWhiteListAddOrUpdateResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgWhiteListAddOrUpdateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgWhiteListAddOrUpdateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgWhiteListAddOrUpdateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgWhiteListAddOrUpdateResponseBody) SetData(v bool) *DsgWhiteListAddOrUpdateResponseBody {
	s.Data = &v
	return s
}

func (s *DsgWhiteListAddOrUpdateResponseBody) SetErrorCode(v string) *DsgWhiteListAddOrUpdateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgWhiteListAddOrUpdateResponseBody) SetErrorMessage(v string) *DsgWhiteListAddOrUpdateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgWhiteListAddOrUpdateResponseBody) SetHttpStatusCode(v int32) *DsgWhiteListAddOrUpdateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgWhiteListAddOrUpdateResponseBody) SetRequestId(v string) *DsgWhiteListAddOrUpdateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgWhiteListAddOrUpdateResponseBody) SetSuccess(v bool) *DsgWhiteListAddOrUpdateResponseBody {
	s.Success = &v
	return s
}

func (s *DsgWhiteListAddOrUpdateResponseBody) Validate() error {
	return dara.Validate(s)
}
