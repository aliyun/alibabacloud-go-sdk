// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticDBTenantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAgenticDBTenantResponseBody
	GetRequestId() *string
}

type DeleteAgenticDBTenantResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 01A13E8C-3773-1A76-BA95-96EA81D260DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAgenticDBTenantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticDBTenantResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgenticDBTenantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAgenticDBTenantResponseBody) SetRequestId(v string) *DeleteAgenticDBTenantResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgenticDBTenantResponseBody) Validate() error {
	return dara.Validate(s)
}
