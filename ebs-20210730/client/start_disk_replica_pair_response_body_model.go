// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDiskReplicaPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartDiskReplicaPairResponseBody
	GetRequestId() *string
}

type StartDiskReplicaPairResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A37597A6-BB99-19B3-85EA-4C2B91F0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDiskReplicaPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *StartDiskReplicaPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartDiskReplicaPairResponseBody) SetRequestId(v string) *StartDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDiskReplicaPairResponseBody) Validate() error {
	return dara.Validate(s)
}
