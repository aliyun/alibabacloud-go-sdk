// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDiskReplicaPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveDiskReplicaPairResponseBody
	GetRequestId() *string
}

type RemoveDiskReplicaPairResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C123F94F-4E38-19AE-942A-A8D6F44F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveDiskReplicaPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDiskReplicaPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveDiskReplicaPairResponseBody) SetRequestId(v string) *RemoveDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveDiskReplicaPairResponseBody) Validate() error {
	return dara.Validate(s)
}
