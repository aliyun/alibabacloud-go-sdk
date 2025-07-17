// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateDataLakePartitionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *BatchCreateDataLakePartitionsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *BatchCreateDataLakePartitionsResponseBody
	GetErrorMessage() *string
	SetPartitions(v []*DLPartition) *BatchCreateDataLakePartitionsResponseBody
	GetPartitions() []*DLPartition
	SetRequestId(v string) *BatchCreateDataLakePartitionsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *BatchCreateDataLakePartitionsResponseBody
	GetSuccess() *string
}

type BatchCreateDataLakePartitionsResponseBody struct {
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Partitions   []*DLPartition `json:"Partitions,omitempty" xml:"Partitions,omitempty" type:"Repeated"`
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchCreateDataLakePartitionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateDataLakePartitionsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateDataLakePartitionsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BatchCreateDataLakePartitionsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *BatchCreateDataLakePartitionsResponseBody) GetPartitions() []*DLPartition {
	return s.Partitions
}

func (s *BatchCreateDataLakePartitionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchCreateDataLakePartitionsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *BatchCreateDataLakePartitionsResponseBody) SetErrorCode(v string) *BatchCreateDataLakePartitionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BatchCreateDataLakePartitionsResponseBody) SetErrorMessage(v string) *BatchCreateDataLakePartitionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BatchCreateDataLakePartitionsResponseBody) SetPartitions(v []*DLPartition) *BatchCreateDataLakePartitionsResponseBody {
	s.Partitions = v
	return s
}

func (s *BatchCreateDataLakePartitionsResponseBody) SetRequestId(v string) *BatchCreateDataLakePartitionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCreateDataLakePartitionsResponseBody) SetSuccess(v string) *BatchCreateDataLakePartitionsResponseBody {
	s.Success = &v
	return s
}

func (s *BatchCreateDataLakePartitionsResponseBody) Validate() error {
	return dara.Validate(s)
}
