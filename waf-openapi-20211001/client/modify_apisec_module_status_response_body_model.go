// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApisecModuleStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyApisecModuleStatusResponseBody
	GetRequestId() *string
}

type ModifyApisecModuleStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19****5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApisecModuleStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApisecModuleStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApisecModuleStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApisecModuleStatusResponseBody) SetRequestId(v string) *ModifyApisecModuleStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApisecModuleStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
