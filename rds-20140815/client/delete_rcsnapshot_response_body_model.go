// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRCSnapshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRCSnapshotResponseBody
	GetRequestId() *string
}

type DeleteRCSnapshotResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8BE834C8-3C25-5AF8-BE3E-C8A690602A7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRCSnapshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRCSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRCSnapshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRCSnapshotResponseBody) SetRequestId(v string) *DeleteRCSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRCSnapshotResponseBody) Validate() error {
	return dara.Validate(s)
}
