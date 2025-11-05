// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskReplicaGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDiskReplicaGroupResponseBody
	GetRequestId() *string
}

type ModifyDiskReplicaGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C123F94F-4E38-19AE-942A-A8D6F44F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDiskReplicaGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskReplicaGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDiskReplicaGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDiskReplicaGroupResponseBody) SetRequestId(v string) *ModifyDiskReplicaGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDiskReplicaGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
