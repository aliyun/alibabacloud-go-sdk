// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostShareKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHostShareKeyResponseBody
	GetRequestId() *string
}

type ModifyHostShareKeyResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHostShareKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostShareKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHostShareKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHostShareKeyResponseBody) SetRequestId(v string) *ModifyHostShareKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHostShareKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
