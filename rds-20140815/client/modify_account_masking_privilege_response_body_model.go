// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountMaskingPrivilegeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]*string) *ModifyAccountMaskingPrivilegeResponseBody
	GetData() map[string]*string
	SetMessage(v string) *ModifyAccountMaskingPrivilegeResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyAccountMaskingPrivilegeResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifyAccountMaskingPrivilegeResponseBody
	GetSuccess() *string
}

type ModifyAccountMaskingPrivilegeResponseBody struct {
	Data map[string]*string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2144F5CC-10C5-3B72-8C74-E52C********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyAccountMaskingPrivilegeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountMaskingPrivilegeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountMaskingPrivilegeResponseBody) GetData() map[string]*string {
	return s.Data
}

func (s *ModifyAccountMaskingPrivilegeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyAccountMaskingPrivilegeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAccountMaskingPrivilegeResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifyAccountMaskingPrivilegeResponseBody) SetData(v map[string]*string) *ModifyAccountMaskingPrivilegeResponseBody {
	s.Data = v
	return s
}

func (s *ModifyAccountMaskingPrivilegeResponseBody) SetMessage(v string) *ModifyAccountMaskingPrivilegeResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyAccountMaskingPrivilegeResponseBody) SetRequestId(v string) *ModifyAccountMaskingPrivilegeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccountMaskingPrivilegeResponseBody) SetSuccess(v string) *ModifyAccountMaskingPrivilegeResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyAccountMaskingPrivilegeResponseBody) Validate() error {
	return dara.Validate(s)
}
