// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLakePartitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteDataLakePartitionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteDataLakePartitionResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteDataLakePartitionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataLakePartitionResponseBody
	GetSuccess() *bool
}

type DeleteDataLakePartitionResponseBody struct {
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
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataLakePartitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLakePartitionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataLakePartitionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteDataLakePartitionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteDataLakePartitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataLakePartitionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataLakePartitionResponseBody) SetErrorCode(v string) *DeleteDataLakePartitionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDataLakePartitionResponseBody) SetErrorMessage(v string) *DeleteDataLakePartitionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDataLakePartitionResponseBody) SetRequestId(v string) *DeleteDataLakePartitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataLakePartitionResponseBody) SetSuccess(v bool) *DeleteDataLakePartitionResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataLakePartitionResponseBody) Validate() error {
	return dara.Validate(s)
}
