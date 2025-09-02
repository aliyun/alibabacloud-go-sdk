// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgUserGroupAddOrUpdateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DsgUserGroupAddOrUpdateResponseBody
	GetData() *bool
	SetErrorCode(v string) *DsgUserGroupAddOrUpdateResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgUserGroupAddOrUpdateResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgUserGroupAddOrUpdateResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DsgUserGroupAddOrUpdateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgUserGroupAddOrUpdateResponseBody
	GetSuccess() *bool
}

type DsgUserGroupAddOrUpdateResponseBody struct {
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
	// The request ID. You can use the ID to locate logs and troubleshoot issues.
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

func (s DsgUserGroupAddOrUpdateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgUserGroupAddOrUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *DsgUserGroupAddOrUpdateResponseBody) GetData() *bool {
	return s.Data
}

func (s *DsgUserGroupAddOrUpdateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgUserGroupAddOrUpdateResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgUserGroupAddOrUpdateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgUserGroupAddOrUpdateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgUserGroupAddOrUpdateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgUserGroupAddOrUpdateResponseBody) SetData(v bool) *DsgUserGroupAddOrUpdateResponseBody {
	s.Data = &v
	return s
}

func (s *DsgUserGroupAddOrUpdateResponseBody) SetErrorCode(v string) *DsgUserGroupAddOrUpdateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgUserGroupAddOrUpdateResponseBody) SetErrorMessage(v string) *DsgUserGroupAddOrUpdateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgUserGroupAddOrUpdateResponseBody) SetHttpStatusCode(v int32) *DsgUserGroupAddOrUpdateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgUserGroupAddOrUpdateResponseBody) SetRequestId(v string) *DsgUserGroupAddOrUpdateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgUserGroupAddOrUpdateResponseBody) SetSuccess(v bool) *DsgUserGroupAddOrUpdateResponseBody {
	s.Success = &v
	return s
}

func (s *DsgUserGroupAddOrUpdateResponseBody) Validate() error {
	return dara.Validate(s)
}
