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
	// The request ID.
	//
	// example:
	//
	// 61B05CCF-EBAB-4BF3-BA67-D77256A721E2
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
