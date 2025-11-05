// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDiskReplicaPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddDiskReplicaPairResponseBody
	GetRequestId() *string
}

type AddDiskReplicaPairResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C123F94F-4E38-19AE-942A-A8D6F44F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDiskReplicaPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *AddDiskReplicaPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDiskReplicaPairResponseBody) SetRequestId(v string) *AddDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDiskReplicaPairResponseBody) Validate() error {
	return dara.Validate(s)
}
