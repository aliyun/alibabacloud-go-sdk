// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceVipResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ModifyDBInstanceVipResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyDBInstanceVipResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyDBInstanceVipResponseBody
	GetSuccess() *bool
}

type ModifyDBInstanceVipResponseBody struct {
	// example:
	//
	// *****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDBInstanceVipResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceVipResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceVipResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyDBInstanceVipResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceVipResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyDBInstanceVipResponseBody) SetMessage(v string) *ModifyDBInstanceVipResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyDBInstanceVipResponseBody) SetRequestId(v string) *ModifyDBInstanceVipResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceVipResponseBody) SetSuccess(v bool) *ModifyDBInstanceVipResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDBInstanceVipResponseBody) Validate() error {
	return dara.Validate(s)
}
