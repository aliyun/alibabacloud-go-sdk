// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReprotectDiskReplicaGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReprotectDiskReplicaGroupResponseBody
	GetRequestId() *string
}

type ReprotectDiskReplicaGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C123F94F-4E38-19AE-942A-A8D6F44F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReprotectDiskReplicaGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReprotectDiskReplicaGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ReprotectDiskReplicaGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReprotectDiskReplicaGroupResponseBody) SetRequestId(v string) *ReprotectDiskReplicaGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReprotectDiskReplicaGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
