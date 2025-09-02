// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetaCollectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteMetaCollectionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteMetaCollectionResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DeleteMetaCollectionResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteMetaCollectionResponseBody
	GetRequestId() *string
	SetStatus(v bool) *DeleteMetaCollectionResponseBody
	GetStatus() *bool
	SetSuccess(v bool) *DeleteMetaCollectionResponseBody
	GetSuccess() *bool
}

type DeleteMetaCollectionResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 999999
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified product does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the operation. Valid values:
	//
	// true: succeeded
	//
	// false: failed
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// true
	//
	// false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMetaCollectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetaCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMetaCollectionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteMetaCollectionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteMetaCollectionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteMetaCollectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMetaCollectionResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *DeleteMetaCollectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMetaCollectionResponseBody) SetErrorCode(v string) *DeleteMetaCollectionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteMetaCollectionResponseBody) SetErrorMessage(v string) *DeleteMetaCollectionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteMetaCollectionResponseBody) SetHttpStatusCode(v int32) *DeleteMetaCollectionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteMetaCollectionResponseBody) SetRequestId(v string) *DeleteMetaCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMetaCollectionResponseBody) SetStatus(v bool) *DeleteMetaCollectionResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteMetaCollectionResponseBody) SetSuccess(v bool) *DeleteMetaCollectionResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMetaCollectionResponseBody) Validate() error {
	return dara.Validate(s)
}
