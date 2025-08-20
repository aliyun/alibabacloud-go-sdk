// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountPrivilegesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAccountPrivilegesResponseBody
	GetRequestId() *string
}

type ModifyAccountPrivilegesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9DD88DE7-824F-1082-AA57-575AFC6517A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountPrivilegesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountPrivilegesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAccountPrivilegesResponseBody) SetRequestId(v string) *ModifyAccountPrivilegesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccountPrivilegesResponseBody) Validate() error {
	return dara.Validate(s)
}
