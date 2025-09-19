// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAssetImportantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *ModifyAssetImportantResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyAssetImportantResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyAssetImportantResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifyAssetImportantResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyAssetImportantResponseBody
	GetSuccess() *bool
}

type ModifyAssetImportantResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// ServerError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// ServerError
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code returned. The value 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 23C85959-1540-514B-93CF-2992C53A1B4F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyAssetImportantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAssetImportantResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAssetImportantResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyAssetImportantResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyAssetImportantResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyAssetImportantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAssetImportantResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyAssetImportantResponseBody) SetErrCode(v string) *ModifyAssetImportantResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyAssetImportantResponseBody) SetErrMessage(v string) *ModifyAssetImportantResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyAssetImportantResponseBody) SetHttpStatusCode(v int32) *ModifyAssetImportantResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyAssetImportantResponseBody) SetRequestId(v string) *ModifyAssetImportantResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAssetImportantResponseBody) SetSuccess(v bool) *ModifyAssetImportantResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyAssetImportantResponseBody) Validate() error {
	return dara.Validate(s)
}
