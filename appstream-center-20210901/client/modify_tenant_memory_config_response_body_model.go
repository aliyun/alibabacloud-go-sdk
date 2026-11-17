// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTenantMemoryConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyTenantMemoryConfigResponseBody
	GetRequestId() *string
}

type ModifyTenantMemoryConfigResponseBody struct {
	// The unique identifier of the request, which can be used for troubleshooting.
	//
	// example:
	//
	// 01A1CEEB-CDFF-18F9-B2BF-9997600D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTenantMemoryConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTenantMemoryConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTenantMemoryConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTenantMemoryConfigResponseBody) SetRequestId(v string) *ModifyTenantMemoryConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTenantMemoryConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
