// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiskReplicaPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *CreateDiskReplicaPairResponseBody
	GetOrderId() *string
	SetReplicaPairId(v string) *CreateDiskReplicaPairResponseBody
	GetReplicaPairId() *string
	SetRequestId(v string) *CreateDiskReplicaPairResponseBody
	GetRequestId() *string
}

type CreateDiskReplicaPairResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 123456****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the replication pair.
	//
	// example:
	//
	// pair-cn-dsa****
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C123F94F-4E38-19AE-942A-A8D6F44F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDiskReplicaPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiskReplicaPairResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateDiskReplicaPairResponseBody) GetReplicaPairId() *string {
	return s.ReplicaPairId
}

func (s *CreateDiskReplicaPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDiskReplicaPairResponseBody) SetOrderId(v string) *CreateDiskReplicaPairResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDiskReplicaPairResponseBody) SetReplicaPairId(v string) *CreateDiskReplicaPairResponseBody {
	s.ReplicaPairId = &v
	return s
}

func (s *CreateDiskReplicaPairResponseBody) SetRequestId(v string) *CreateDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDiskReplicaPairResponseBody) Validate() error {
	return dara.Validate(s)
}
