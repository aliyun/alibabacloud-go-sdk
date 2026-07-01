// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticDBTenantApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAgenticDBTenantApiKeyResponseBody
	GetRequestId() *string
}

type DeleteAgenticDBTenantApiKeyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E5F6A7B8-C9D0-1234-EFAB-345678901234
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAgenticDBTenantApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticDBTenantApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgenticDBTenantApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAgenticDBTenantApiKeyResponseBody) SetRequestId(v string) *DeleteAgenticDBTenantApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgenticDBTenantApiKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
