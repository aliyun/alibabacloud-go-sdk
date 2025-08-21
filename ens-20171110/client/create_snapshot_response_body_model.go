// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSnapshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *CreateSnapshotResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateSnapshotResponseBody
	GetRequestId() *string
	SetSnapShotId(v []*string) *CreateSnapshotResponseBody
	GetSnapShotId() []*string
}

type CreateSnapshotResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// 21969183547****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IDs of the snapshots.
	//
	// example:
	//
	// s-bp17441ohwka0yuh****
	SnapShotId []*string `json:"SnapShotId,omitempty" xml:"SnapShotId,omitempty" type:"Repeated"`
}

func (s CreateSnapshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateSnapshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSnapshotResponseBody) GetSnapShotId() []*string {
	return s.SnapShotId
}

func (s *CreateSnapshotResponseBody) SetOrderId(v string) *CreateSnapshotResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateSnapshotResponseBody) SetRequestId(v string) *CreateSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSnapshotResponseBody) SetSnapShotId(v []*string) *CreateSnapshotResponseBody {
	s.SnapShotId = v
	return s
}

func (s *CreateSnapshotResponseBody) Validate() error {
	return dara.Validate(s)
}
