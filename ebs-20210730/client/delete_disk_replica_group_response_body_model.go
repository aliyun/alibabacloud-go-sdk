// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDiskReplicaGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDiskReplicaGroupResponseBody
	GetRequestId() *string
}

type DeleteDiskReplicaGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C123F94F-4E38-19AE-942A-A8D6F44F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDiskReplicaGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDiskReplicaGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDiskReplicaGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDiskReplicaGroupResponseBody) SetRequestId(v string) *DeleteDiskReplicaGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDiskReplicaGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
