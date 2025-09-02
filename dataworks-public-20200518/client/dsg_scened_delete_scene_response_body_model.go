// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgScenedDeleteSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DsgScenedDeleteSceneResponseBody
	GetData() *bool
	SetErrorCode(v string) *DsgScenedDeleteSceneResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgScenedDeleteSceneResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgScenedDeleteSceneResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DsgScenedDeleteSceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgScenedDeleteSceneResponseBody
	GetSuccess() *bool
}

type DsgScenedDeleteSceneResponseBody struct {
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
	// The request ID.
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

func (s DsgScenedDeleteSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgScenedDeleteSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DsgScenedDeleteSceneResponseBody) GetData() *bool {
	return s.Data
}

func (s *DsgScenedDeleteSceneResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgScenedDeleteSceneResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgScenedDeleteSceneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgScenedDeleteSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgScenedDeleteSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgScenedDeleteSceneResponseBody) SetData(v bool) *DsgScenedDeleteSceneResponseBody {
	s.Data = &v
	return s
}

func (s *DsgScenedDeleteSceneResponseBody) SetErrorCode(v string) *DsgScenedDeleteSceneResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgScenedDeleteSceneResponseBody) SetErrorMessage(v string) *DsgScenedDeleteSceneResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgScenedDeleteSceneResponseBody) SetHttpStatusCode(v int32) *DsgScenedDeleteSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgScenedDeleteSceneResponseBody) SetRequestId(v string) *DsgScenedDeleteSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgScenedDeleteSceneResponseBody) SetSuccess(v bool) *DsgScenedDeleteSceneResponseBody {
	s.Success = &v
	return s
}

func (s *DsgScenedDeleteSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
