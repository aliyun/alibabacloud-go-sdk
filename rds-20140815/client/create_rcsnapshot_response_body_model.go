// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRCSnapshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRCSnapshotResponseBody
	GetRequestId() *string
	SetSnapshotId(v string) *CreateRCSnapshotResponseBody
	GetSnapshotId() *string
}

type CreateRCSnapshotResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CE93CC0D-B65D-5723-AAB1-08CB8BBABAB9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The snapshot ID.
	//
	// example:
	//
	// rcds-7mbefjzkqccvdev****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CreateRCSnapshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRCSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRCSnapshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRCSnapshotResponseBody) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateRCSnapshotResponseBody) SetRequestId(v string) *CreateRCSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRCSnapshotResponseBody) SetSnapshotId(v string) *CreateRCSnapshotResponseBody {
	s.SnapshotId = &v
	return s
}

func (s *CreateRCSnapshotResponseBody) Validate() error {
	return dara.Validate(s)
}
