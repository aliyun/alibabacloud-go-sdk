// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyResourceGroupResponseBody
	GetRequestId() *string
}

type ModifyResourceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 65BDA532-28AF-4122-AA39-B382721EEE64
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyResourceGroupResponseBody) SetRequestId(v string) *ModifyResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
