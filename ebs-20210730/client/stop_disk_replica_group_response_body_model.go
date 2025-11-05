// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDiskReplicaGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopDiskReplicaGroupResponseBody
	GetRequestId() *string
}

type StopDiskReplicaGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C123F94F-4E38-19AE-942A-A8D6F44F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDiskReplicaGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopDiskReplicaGroupResponseBody) GoString() string {
	return s.String()
}

func (s *StopDiskReplicaGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDiskReplicaGroupResponseBody) SetRequestId(v string) *StopDiskReplicaGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDiskReplicaGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
