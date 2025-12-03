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
	// example:
	//
	// AFAA617B-3268-5883-982B-DB8EC8CC1F1B
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
