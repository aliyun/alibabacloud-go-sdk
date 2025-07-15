// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletionProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletionProtectionResponseBody
	GetRequestId() *string
}

type DeletionProtectionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// BAAEF103-96C4-4454-9210-066F2405F511
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletionProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletionProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeletionProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletionProtectionResponseBody) SetRequestId(v string) *DeletionProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletionProtectionResponseBody) Validate() error {
	return dara.Validate(s)
}
