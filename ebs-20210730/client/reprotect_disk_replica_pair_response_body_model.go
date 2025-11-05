// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReprotectDiskReplicaPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReprotectDiskReplicaPairResponseBody
	GetRequestId() *string
}

type ReprotectDiskReplicaPairResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C123F94F-4E38-19AE-942A-A8D6F44F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReprotectDiskReplicaPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReprotectDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *ReprotectDiskReplicaPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReprotectDiskReplicaPairResponseBody) SetRequestId(v string) *ReprotectDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReprotectDiskReplicaPairResponseBody) Validate() error {
	return dara.Validate(s)
}
