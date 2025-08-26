// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLakeTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteDataLakeTableResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteDataLakeTableResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteDataLakeTableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataLakeTableResponseBody
	GetSuccess() *bool
}

type DeleteDataLakeTableResponseBody struct {
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 8E88933E-E3D4-5BA8-8CBF-0A1CAE66****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataLakeTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLakeTableResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataLakeTableResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteDataLakeTableResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteDataLakeTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataLakeTableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataLakeTableResponseBody) SetErrorCode(v string) *DeleteDataLakeTableResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDataLakeTableResponseBody) SetErrorMessage(v string) *DeleteDataLakeTableResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDataLakeTableResponseBody) SetRequestId(v string) *DeleteDataLakeTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataLakeTableResponseBody) SetSuccess(v bool) *DeleteDataLakeTableResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataLakeTableResponseBody) Validate() error {
	return dara.Validate(s)
}
