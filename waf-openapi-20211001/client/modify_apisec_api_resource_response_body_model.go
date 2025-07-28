// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApisecApiResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyApisecApiResourceResponseBody
	GetRequestId() *string
}

type ModifyApisecApiResourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19****5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApisecApiResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApisecApiResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApisecApiResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApisecApiResourceResponseBody) SetRequestId(v string) *ModifyApisecApiResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApisecApiResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
