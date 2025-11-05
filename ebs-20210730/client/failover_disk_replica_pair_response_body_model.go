// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFailoverDiskReplicaPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *FailoverDiskReplicaPairResponseBody
	GetRequestId() *string
}

type FailoverDiskReplicaPairResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C123F94F-4E38-19AE-942A-A8D6F44F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FailoverDiskReplicaPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FailoverDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *FailoverDiskReplicaPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FailoverDiskReplicaPairResponseBody) SetRequestId(v string) *FailoverDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *FailoverDiskReplicaPairResponseBody) Validate() error {
	return dara.Validate(s)
}
