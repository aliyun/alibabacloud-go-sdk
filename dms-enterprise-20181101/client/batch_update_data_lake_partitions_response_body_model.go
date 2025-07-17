// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateDataLakePartitionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *BatchUpdateDataLakePartitionsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *BatchUpdateDataLakePartitionsResponseBody
	GetErrorMessage() *string
	SetPartitionErrors(v []*PartitionError) *BatchUpdateDataLakePartitionsResponseBody
	GetPartitionErrors() []*PartitionError
	SetRequestId(v string) *BatchUpdateDataLakePartitionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchUpdateDataLakePartitionsResponseBody
	GetSuccess() *bool
}

type BatchUpdateDataLakePartitionsResponseBody struct {
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
	// C5B8E84B-42B6-4374-AD5A-6264E1753378
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchUpdateDataLakePartitionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateDataLakePartitionsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateDataLakePartitionsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BatchUpdateDataLakePartitionsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *BatchUpdateDataLakePartitionsResponseBody) GetPartitionErrors() []*PartitionError {
	return s.PartitionErrors
}

func (s *BatchUpdateDataLakePartitionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchUpdateDataLakePartitionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchUpdateDataLakePartitionsResponseBody) SetErrorCode(v string) *BatchUpdateDataLakePartitionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BatchUpdateDataLakePartitionsResponseBody) SetErrorMessage(v string) *BatchUpdateDataLakePartitionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BatchUpdateDataLakePartitionsResponseBody) SetPartitionErrors(v []*PartitionError) *BatchUpdateDataLakePartitionsResponseBody {
	s.PartitionErrors = v
	return s
}

func (s *BatchUpdateDataLakePartitionsResponseBody) SetRequestId(v string) *BatchUpdateDataLakePartitionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUpdateDataLakePartitionsResponseBody) SetSuccess(v bool) *BatchUpdateDataLakePartitionsResponseBody {
	s.Success = &v
	return s
}

func (s *BatchUpdateDataLakePartitionsResponseBody) Validate() error {
	return dara.Validate(s)
}
