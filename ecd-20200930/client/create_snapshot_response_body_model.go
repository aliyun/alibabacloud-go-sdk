// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSnapshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSnapshotResponseBody
	GetRequestId() *string
	SetSnapshotId(v string) *CreateSnapshotResponseBody
	GetSnapshotId() *string
}

type CreateSnapshotResponseBody struct {
	// The ID of the region.
	//
	// example:
	//
	// 3EB7FCEE-D731-4948-85A3-4B2C341CA983
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the snapshot.
	//
	// example:
	//
	// s-2ze81owrnv9pity4****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CreateSnapshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSnapshotResponseBody) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateSnapshotResponseBody) SetRequestId(v string) *CreateSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSnapshotResponseBody) SetSnapshotId(v string) *CreateSnapshotResponseBody {
	s.SnapshotId = &v
	return s
}

func (s *CreateSnapshotResponseBody) Validate() error {
	return dara.Validate(s)
}
