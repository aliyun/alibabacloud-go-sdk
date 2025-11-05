// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearReplicaGroupDrillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ClearReplicaGroupDrillResponseBody
	GetRequestId() *string
}

type ClearReplicaGroupDrillResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C123F94F-4E38-19AE-942A-A8D6F44F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClearReplicaGroupDrillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClearReplicaGroupDrillResponseBody) GoString() string {
	return s.String()
}

func (s *ClearReplicaGroupDrillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClearReplicaGroupDrillResponseBody) SetRequestId(v string) *ClearReplicaGroupDrillResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClearReplicaGroupDrillResponseBody) Validate() error {
	return dara.Validate(s)
}
