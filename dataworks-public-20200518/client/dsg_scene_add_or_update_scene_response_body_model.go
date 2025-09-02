// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgSceneAddOrUpdateSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DsgSceneAddOrUpdateSceneResponseBody
	GetData() *bool
	SetErrorCode(v string) *DsgSceneAddOrUpdateSceneResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgSceneAddOrUpdateSceneResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgSceneAddOrUpdateSceneResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DsgSceneAddOrUpdateSceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgSceneAddOrUpdateSceneResponseBody
	GetSuccess() *bool
}

type DsgSceneAddOrUpdateSceneResponseBody struct {
	// The operation result.
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

func (s DsgSceneAddOrUpdateSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgSceneAddOrUpdateSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DsgSceneAddOrUpdateSceneResponseBody) GetData() *bool {
	return s.Data
}

func (s *DsgSceneAddOrUpdateSceneResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgSceneAddOrUpdateSceneResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgSceneAddOrUpdateSceneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgSceneAddOrUpdateSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgSceneAddOrUpdateSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgSceneAddOrUpdateSceneResponseBody) SetData(v bool) *DsgSceneAddOrUpdateSceneResponseBody {
	s.Data = &v
	return s
}

func (s *DsgSceneAddOrUpdateSceneResponseBody) SetErrorCode(v string) *DsgSceneAddOrUpdateSceneResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgSceneAddOrUpdateSceneResponseBody) SetErrorMessage(v string) *DsgSceneAddOrUpdateSceneResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgSceneAddOrUpdateSceneResponseBody) SetHttpStatusCode(v int32) *DsgSceneAddOrUpdateSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgSceneAddOrUpdateSceneResponseBody) SetRequestId(v string) *DsgSceneAddOrUpdateSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgSceneAddOrUpdateSceneResponseBody) SetSuccess(v bool) *DsgSceneAddOrUpdateSceneResponseBody {
	s.Success = &v
	return s
}

func (s *DsgSceneAddOrUpdateSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
