// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTenantConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyTenantConfigResponseBody
	GetRequestId() *string
}

type ModifyTenantConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTenantConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTenantConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTenantConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTenantConfigResponseBody) SetRequestId(v string) *ModifyTenantConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTenantConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
