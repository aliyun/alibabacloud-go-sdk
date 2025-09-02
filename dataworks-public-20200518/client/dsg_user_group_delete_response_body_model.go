// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgUserGroupDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DsgUserGroupDeleteResponseBody
	GetData() *bool
	SetErrorCode(v string) *DsgUserGroupDeleteResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgUserGroupDeleteResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgUserGroupDeleteResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DsgUserGroupDeleteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgUserGroupDeleteResponseBody
	GetSuccess() *bool
}

type DsgUserGroupDeleteResponseBody struct {
	// The operation result. Valid values:
	//
	// 	- true: The user group is deleted.
	//
	// 	- false: The user group fails to be deleted.
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

func (s DsgUserGroupDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgUserGroupDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *DsgUserGroupDeleteResponseBody) GetData() *bool {
	return s.Data
}

func (s *DsgUserGroupDeleteResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgUserGroupDeleteResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgUserGroupDeleteResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgUserGroupDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgUserGroupDeleteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgUserGroupDeleteResponseBody) SetData(v bool) *DsgUserGroupDeleteResponseBody {
	s.Data = &v
	return s
}

func (s *DsgUserGroupDeleteResponseBody) SetErrorCode(v string) *DsgUserGroupDeleteResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgUserGroupDeleteResponseBody) SetErrorMessage(v string) *DsgUserGroupDeleteResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgUserGroupDeleteResponseBody) SetHttpStatusCode(v int32) *DsgUserGroupDeleteResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgUserGroupDeleteResponseBody) SetRequestId(v string) *DsgUserGroupDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgUserGroupDeleteResponseBody) SetSuccess(v bool) *DsgUserGroupDeleteResponseBody {
	s.Success = &v
	return s
}

func (s *DsgUserGroupDeleteResponseBody) Validate() error {
	return dara.Validate(s)
}
