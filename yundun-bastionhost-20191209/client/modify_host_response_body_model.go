// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHostResponseBody
	GetRequestId() *string
}

type ModifyHostResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHostResponseBody) SetRequestId(v string) *ModifyHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHostResponseBody) Validate() error {
	return dara.Validate(s)
}
