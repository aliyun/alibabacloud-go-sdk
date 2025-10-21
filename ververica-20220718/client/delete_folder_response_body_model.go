// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFolderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteFolderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteFolderResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *DeleteFolderResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *DeleteFolderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteFolderResponseBody
	GetSuccess() *bool
}

type DeleteFolderResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteFolderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFolderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFolderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteFolderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteFolderResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *DeleteFolderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFolderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteFolderResponseBody) SetErrorCode(v string) *DeleteFolderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteFolderResponseBody) SetErrorMessage(v string) *DeleteFolderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFolderResponseBody) SetHttpCode(v int32) *DeleteFolderResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteFolderResponseBody) SetRequestId(v string) *DeleteFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFolderResponseBody) SetSuccess(v bool) *DeleteFolderResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFolderResponseBody) Validate() error {
	return dara.Validate(s)
}
