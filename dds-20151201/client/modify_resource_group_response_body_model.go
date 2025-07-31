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
	// The ID of the request.
	//
	// example:
	//
	// 5DD0DEC3-24A1-5268-8E0B-5B264CA1****
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
