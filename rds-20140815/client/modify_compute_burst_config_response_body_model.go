// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyComputeBurstConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyComputeBurstConfigResponseBody
	GetRequestId() *string
}

type ModifyComputeBurstConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C816A4BF-A6EC-4722-95F9-2055859CCFD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyComputeBurstConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyComputeBurstConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyComputeBurstConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyComputeBurstConfigResponseBody) SetRequestId(v string) *ModifyComputeBurstConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyComputeBurstConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
