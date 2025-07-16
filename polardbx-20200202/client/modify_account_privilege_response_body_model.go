// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountPrivilegeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ModifyAccountPrivilegeResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyAccountPrivilegeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyAccountPrivilegeResponseBody
	GetSuccess() *bool
}

type ModifyAccountPrivilegeResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73559800-3c8c-11ec-bd40-99cfcff3fe1e
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyAccountPrivilegeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountPrivilegeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyAccountPrivilegeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAccountPrivilegeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyAccountPrivilegeResponseBody) SetMessage(v string) *ModifyAccountPrivilegeResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyAccountPrivilegeResponseBody) SetRequestId(v string) *ModifyAccountPrivilegeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccountPrivilegeResponseBody) SetSuccess(v bool) *ModifyAccountPrivilegeResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyAccountPrivilegeResponseBody) Validate() error {
	return dara.Validate(s)
}
