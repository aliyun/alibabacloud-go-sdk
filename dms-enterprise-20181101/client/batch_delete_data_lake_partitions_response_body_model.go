// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteDataLakePartitionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *BatchDeleteDataLakePartitionsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *BatchDeleteDataLakePartitionsResponseBody
	GetErrorMessage() *string
	SetPartitionErrors(v []*PartitionError) *BatchDeleteDataLakePartitionsResponseBody
	GetPartitionErrors() []*PartitionError
	SetRequestId(v string) *BatchDeleteDataLakePartitionsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *BatchDeleteDataLakePartitionsResponseBody
	GetSuccess() *string
}

type BatchDeleteDataLakePartitionsResponseBody struct {
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage    *string           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	PartitionErrors []*PartitionError `json:"PartitionErrors,omitempty" xml:"PartitionErrors,omitempty" type:"Repeated"`
	// example:
	//
	// 5B96E35F-A58E-5399-9041-09CF9A1E46EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchDeleteDataLakePartitionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteDataLakePartitionsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteDataLakePartitionsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BatchDeleteDataLakePartitionsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *BatchDeleteDataLakePartitionsResponseBody) GetPartitionErrors() []*PartitionError {
	return s.PartitionErrors
}

func (s *BatchDeleteDataLakePartitionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeleteDataLakePartitionsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *BatchDeleteDataLakePartitionsResponseBody) SetErrorCode(v string) *BatchDeleteDataLakePartitionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BatchDeleteDataLakePartitionsResponseBody) SetErrorMessage(v string) *BatchDeleteDataLakePartitionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BatchDeleteDataLakePartitionsResponseBody) SetPartitionErrors(v []*PartitionError) *BatchDeleteDataLakePartitionsResponseBody {
	s.PartitionErrors = v
	return s
}

func (s *BatchDeleteDataLakePartitionsResponseBody) SetRequestId(v string) *BatchDeleteDataLakePartitionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteDataLakePartitionsResponseBody) SetSuccess(v string) *BatchDeleteDataLakePartitionsResponseBody {
	s.Success = &v
	return s
}

func (s *BatchDeleteDataLakePartitionsResponseBody) Validate() error {
	return dara.Validate(s)
}
