// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskReplicaPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDiskReplicaPairResponseBody
	GetRequestId() *string
}

type ModifyDiskReplicaPairResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C123F94F-4E38-19AE-942A-A8D6F44F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDiskReplicaPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDiskReplicaPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDiskReplicaPairResponseBody) SetRequestId(v string) *ModifyDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDiskReplicaPairResponseBody) Validate() error {
	return dara.Validate(s)
}
