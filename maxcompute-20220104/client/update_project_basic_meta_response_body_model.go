// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectBasicMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UpdateProjectBasicMetaResponseBody
	GetData() *string
	SetErrorCode(v string) *UpdateProjectBasicMetaResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *UpdateProjectBasicMetaResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *UpdateProjectBasicMetaResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *UpdateProjectBasicMetaResponseBody
	GetRequestId() *string
}

type UpdateProjectBasicMetaResponseBody struct {
	// The data returned.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The error code.
	//
	// example:
	//
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// This object does not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// - 1xx: informational response. The request is received and is being processed.
	//
	// - 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// - 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0b87b7e216652820458545253e8b0a
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateProjectBasicMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectBasicMetaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectBasicMetaResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateProjectBasicMetaResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateProjectBasicMetaResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *UpdateProjectBasicMetaResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *UpdateProjectBasicMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProjectBasicMetaResponseBody) SetData(v string) *UpdateProjectBasicMetaResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateProjectBasicMetaResponseBody) SetErrorCode(v string) *UpdateProjectBasicMetaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateProjectBasicMetaResponseBody) SetErrorMsg(v string) *UpdateProjectBasicMetaResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateProjectBasicMetaResponseBody) SetHttpCode(v int32) *UpdateProjectBasicMetaResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateProjectBasicMetaResponseBody) SetRequestId(v string) *UpdateProjectBasicMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectBasicMetaResponseBody) Validate() error {
	return dara.Validate(s)
}
