// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataLakePartitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateDataLakePartitionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDataLakePartitionResponseBody
	GetErrorMessage() *string
	SetPartition(v *DLPartition) *CreateDataLakePartitionResponseBody
	GetPartition() *DLPartition
	SetRequestId(v string) *CreateDataLakePartitionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataLakePartitionResponseBody
	GetSuccess() *bool
}

type CreateDataLakePartitionResponseBody struct {
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
	// The information about the partition.
	Partition *DLPartition `json:"Partition,omitempty" xml:"Partition,omitempty"`
	// The request ID. You can use the request ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 427688B8-ADFB-4C4E-9D45-EF5C1FD6E23D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDataLakePartitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataLakePartitionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataLakePartitionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDataLakePartitionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDataLakePartitionResponseBody) GetPartition() *DLPartition {
	return s.Partition
}

func (s *CreateDataLakePartitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataLakePartitionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataLakePartitionResponseBody) SetErrorCode(v string) *CreateDataLakePartitionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDataLakePartitionResponseBody) SetErrorMessage(v string) *CreateDataLakePartitionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDataLakePartitionResponseBody) SetPartition(v *DLPartition) *CreateDataLakePartitionResponseBody {
	s.Partition = v
	return s
}

func (s *CreateDataLakePartitionResponseBody) SetRequestId(v string) *CreateDataLakePartitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataLakePartitionResponseBody) SetSuccess(v bool) *CreateDataLakePartitionResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataLakePartitionResponseBody) Validate() error {
	if s.Partition != nil {
		if err := s.Partition.Validate(); err != nil {
			return err
		}
	}
	return nil
}
