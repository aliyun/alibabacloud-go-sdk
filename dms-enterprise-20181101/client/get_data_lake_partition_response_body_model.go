// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataLakePartitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetDataLakePartitionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataLakePartitionResponseBody
	GetErrorMessage() *string
	SetPartition(v *DLPartition) *GetDataLakePartitionResponseBody
	GetPartition() *DLPartition
	SetRequestId(v string) *GetDataLakePartitionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataLakePartitionResponseBody
	GetSuccess() *bool
}

type GetDataLakePartitionResponseBody struct {
	// The error code that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The queried partition.
	Partition *DLPartition `json:"Partition,omitempty" xml:"Partition,omitempty"`
	// The request ID. You can use the request ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// FE8EE2F1-4880-46BC-A704-5CF63EAF9A04
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

func (s GetDataLakePartitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataLakePartitionResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataLakePartitionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataLakePartitionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataLakePartitionResponseBody) GetPartition() *DLPartition {
	return s.Partition
}

func (s *GetDataLakePartitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataLakePartitionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataLakePartitionResponseBody) SetErrorCode(v string) *GetDataLakePartitionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataLakePartitionResponseBody) SetErrorMessage(v string) *GetDataLakePartitionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataLakePartitionResponseBody) SetPartition(v *DLPartition) *GetDataLakePartitionResponseBody {
	s.Partition = v
	return s
}

func (s *GetDataLakePartitionResponseBody) SetRequestId(v string) *GetDataLakePartitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataLakePartitionResponseBody) SetSuccess(v bool) *GetDataLakePartitionResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataLakePartitionResponseBody) Validate() error {
	if s.Partition != nil {
		if err := s.Partition.Validate(); err != nil {
			return err
		}
	}
	return nil
}
