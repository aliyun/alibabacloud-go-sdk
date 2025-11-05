// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDiskReplicaPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopDiskReplicaPairResponseBody
	GetRequestId() *string
}

type StopDiskReplicaPairResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A37597A6-BB99-19B3-85EA-4C2B91F0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDiskReplicaPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *StopDiskReplicaPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDiskReplicaPairResponseBody) SetRequestId(v string) *StopDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDiskReplicaPairResponseBody) Validate() error {
	return dara.Validate(s)
}
