// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHostGroupResponseBody
	GetRequestId() *string
}

type ModifyHostGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHostGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHostGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHostGroupResponseBody) SetRequestId(v string) *ModifyHostGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHostGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
