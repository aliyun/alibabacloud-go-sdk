// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHostAccountResponseBody
	GetRequestId() *string
}

type ModifyHostAccountResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHostAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHostAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHostAccountResponseBody) SetRequestId(v string) *ModifyHostAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHostAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
