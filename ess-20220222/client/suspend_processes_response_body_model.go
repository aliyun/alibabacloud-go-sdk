// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendProcessesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SuspendProcessesResponseBody
	GetRequestId() *string
}

type SuspendProcessesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 3E2033F0-03B4-419D-BCE2-C2339DB51****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SuspendProcessesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SuspendProcessesResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendProcessesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SuspendProcessesResponseBody) SetRequestId(v string) *SuspendProcessesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendProcessesResponseBody) Validate() error {
	return dara.Validate(s)
}
