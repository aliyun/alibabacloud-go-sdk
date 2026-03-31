// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourceLogFieldConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyResourceLogFieldConfigResponseBody
	GetRequestId() *string
}

type ModifyResourceLogFieldConfigResponseBody struct {
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyResourceLogFieldConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourceLogFieldConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyResourceLogFieldConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyResourceLogFieldConfigResponseBody) SetRequestId(v string) *ModifyResourceLogFieldConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyResourceLogFieldConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
