// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPartitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPartition(v *Partition) *GetPartitionResponseBody
	GetPartition() *Partition
	SetRequestId(v string) *GetPartitionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPartitionResponseBody
	GetSuccess() *bool
}

type GetPartitionResponseBody struct {
	// Partition details.
	Partition *Partition `json:"Partition,omitempty" xml:"Partition,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D1E2E5BC-xxxx-xxxx-xxxx-xxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPartitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPartitionResponseBody) GoString() string {
	return s.String()
}

func (s *GetPartitionResponseBody) GetPartition() *Partition {
	return s.Partition
}

func (s *GetPartitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPartitionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPartitionResponseBody) SetPartition(v *Partition) *GetPartitionResponseBody {
	s.Partition = v
	return s
}

func (s *GetPartitionResponseBody) SetRequestId(v string) *GetPartitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPartitionResponseBody) SetSuccess(v bool) *GetPartitionResponseBody {
	s.Success = &v
	return s
}

func (s *GetPartitionResponseBody) Validate() error {
	if s.Partition != nil {
		if err := s.Partition.Validate(); err != nil {
			return err
		}
	}
	return nil
}
