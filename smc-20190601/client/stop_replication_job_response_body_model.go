// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopReplicationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopReplicationJobResponseBody
	GetRequestId() *string
}

type StopReplicationJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopReplicationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopReplicationJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopReplicationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopReplicationJobResponseBody) SetRequestId(v string) *StopReplicationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopReplicationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
