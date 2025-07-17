// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataLakePartitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateDataLakePartitionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateDataLakePartitionResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateDataLakePartitionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDataLakePartitionResponseBody
	GetSuccess() *bool
}

type UpdateDataLakePartitionResponseBody struct {
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
	// 427688B8-ADFB-4C4E-9D45-EF5C1FD6E23D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataLakePartitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataLakePartitionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataLakePartitionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateDataLakePartitionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateDataLakePartitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataLakePartitionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataLakePartitionResponseBody) SetErrorCode(v string) *UpdateDataLakePartitionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDataLakePartitionResponseBody) SetErrorMessage(v string) *UpdateDataLakePartitionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDataLakePartitionResponseBody) SetRequestId(v string) *UpdateDataLakePartitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataLakePartitionResponseBody) SetSuccess(v bool) *UpdateDataLakePartitionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataLakePartitionResponseBody) Validate() error {
	return dara.Validate(s)
}
