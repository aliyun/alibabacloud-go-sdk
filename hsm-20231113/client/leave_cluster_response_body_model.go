// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLeaveClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *LeaveClusterResponseBody
	GetRequestId() *string
}

type LeaveClusterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LeaveClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LeaveClusterResponseBody) GoString() string {
	return s.String()
}

func (s *LeaveClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LeaveClusterResponseBody) SetRequestId(v string) *LeaveClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *LeaveClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
