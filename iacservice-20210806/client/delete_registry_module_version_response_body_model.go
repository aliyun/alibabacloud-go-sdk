// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRegistryModuleVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRegistryModuleVersionResponseBody
	GetRequestId() *string
}

type DeleteRegistryModuleVersionResponseBody struct {
	// example:
	//
	// 491A1E2E-EA1E-5F90-958A-A53EB67780FC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteRegistryModuleVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRegistryModuleVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRegistryModuleVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRegistryModuleVersionResponseBody) SetRequestId(v string) *DeleteRegistryModuleVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRegistryModuleVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
