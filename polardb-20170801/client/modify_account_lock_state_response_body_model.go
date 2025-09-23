// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountLockStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAccountLockStateResponseBody
	GetRequestId() *string
}

type ModifyAccountLockStateResponseBody struct {
	// example:
	//
	// B762E8C5-0129-51DB-80C8-********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountLockStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountLockStateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountLockStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAccountLockStateResponseBody) SetRequestId(v string) *ModifyAccountLockStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccountLockStateResponseBody) Validate() error {
	return dara.Validate(s)
}
