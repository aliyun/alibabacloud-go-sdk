// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDiskReplicaPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDiskReplicaPairResponseBody
	GetRequestId() *string
}

type DeleteDiskReplicaPairResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A37597A6-BB99-19B3-85EA-4C2B91F0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDiskReplicaPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDiskReplicaPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDiskReplicaPairResponseBody) SetRequestId(v string) *DeleteDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDiskReplicaPairResponseBody) Validate() error {
	return dara.Validate(s)
}
