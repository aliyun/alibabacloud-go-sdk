// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFailoverDiskReplicaGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *FailoverDiskReplicaGroupResponseBody
	GetRequestId() *string
}

type FailoverDiskReplicaGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C123F94F-4E38-19AE-942A-A8D6F44F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FailoverDiskReplicaGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FailoverDiskReplicaGroupResponseBody) GoString() string {
	return s.String()
}

func (s *FailoverDiskReplicaGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FailoverDiskReplicaGroupResponseBody) SetRequestId(v string) *FailoverDiskReplicaGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *FailoverDiskReplicaGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
