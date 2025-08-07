// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAccountPasswordResponseBody
	GetRequestId() *string
}

type ModifyAccountPasswordResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAccountPasswordResponseBody) SetRequestId(v string) *ModifyAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccountPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
