// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySecurityIpsResponseBody
	GetRequestId() *string
}

type ModifySecurityIpsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 871C698F-B43D-4D1D-ACD6-DF56B0F89978
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySecurityIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySecurityIpsResponseBody) SetRequestId(v string) *ModifySecurityIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySecurityIpsResponseBody) Validate() error {
	return dara.Validate(s)
}
