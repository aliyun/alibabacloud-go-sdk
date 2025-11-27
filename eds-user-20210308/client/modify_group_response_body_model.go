// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyGroupResponseBody
	GetRequestId() *string
}

type ModifyGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 868B8926-2E7A-5BE7-9897-E811E994****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyGroupResponseBody) SetRequestId(v string) *ModifyGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
