// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartReplicaGroupDrillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDrillId(v string) *StartReplicaGroupDrillResponseBody
	GetDrillId() *string
	SetRequestId(v string) *StartReplicaGroupDrillResponseBody
	GetRequestId() *string
}

type StartReplicaGroupDrillResponseBody struct {
	// The drill ID.
	//
	// example:
	//
	// pg-drill-xxxx
	DrillId *string `json:"DrillId,omitempty" xml:"DrillId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C123F94F-4E38-19AE-942A-A8D6F44F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartReplicaGroupDrillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartReplicaGroupDrillResponseBody) GoString() string {
	return s.String()
}

func (s *StartReplicaGroupDrillResponseBody) GetDrillId() *string {
	return s.DrillId
}

func (s *StartReplicaGroupDrillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartReplicaGroupDrillResponseBody) SetDrillId(v string) *StartReplicaGroupDrillResponseBody {
	s.DrillId = &v
	return s
}

func (s *StartReplicaGroupDrillResponseBody) SetRequestId(v string) *StartReplicaGroupDrillResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartReplicaGroupDrillResponseBody) Validate() error {
	return dara.Validate(s)
}
