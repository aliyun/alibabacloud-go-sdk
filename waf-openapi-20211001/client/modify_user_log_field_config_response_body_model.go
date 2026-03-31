// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserLogFieldConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyUserLogFieldConfigResponseBody
	GetRequestId() *string
}

type ModifyUserLogFieldConfigResponseBody struct {
	// example:
	//
	// E2D63742-9BAA-*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserLogFieldConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserLogFieldConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserLogFieldConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyUserLogFieldConfigResponseBody) SetRequestId(v string) *ModifyUserLogFieldConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyUserLogFieldConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
