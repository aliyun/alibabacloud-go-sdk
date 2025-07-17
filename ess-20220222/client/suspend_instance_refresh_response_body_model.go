// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendInstanceRefreshResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SuspendInstanceRefreshResponseBody
	GetRequestId() *string
}

type SuspendInstanceRefreshResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SuspendInstanceRefreshResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SuspendInstanceRefreshResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendInstanceRefreshResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SuspendInstanceRefreshResponseBody) SetRequestId(v string) *SuspendInstanceRefreshResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendInstanceRefreshResponseBody) Validate() error {
	return dara.Validate(s)
}
