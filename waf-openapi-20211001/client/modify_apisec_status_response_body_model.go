// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApisecStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyApisecStatusResponseBody
	GetRequestId() *string
}

type ModifyApisecStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApisecStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApisecStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApisecStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApisecStatusResponseBody) SetRequestId(v string) *ModifyApisecStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApisecStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
