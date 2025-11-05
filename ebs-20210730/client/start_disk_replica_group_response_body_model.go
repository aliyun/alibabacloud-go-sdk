// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDiskReplicaGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartDiskReplicaGroupResponseBody
	GetRequestId() *string
}

type StartDiskReplicaGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C123F94F-4E38-19AE-942A-A8D6F44F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDiskReplicaGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartDiskReplicaGroupResponseBody) GoString() string {
	return s.String()
}

func (s *StartDiskReplicaGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartDiskReplicaGroupResponseBody) SetRequestId(v string) *StartDiskReplicaGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDiskReplicaGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
