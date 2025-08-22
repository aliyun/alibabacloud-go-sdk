// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRegistryModuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRegistryModuleResponseBody
	GetRequestId() *string
}

type DeleteRegistryModuleResponseBody struct {
	// example:
	//
	// 545995A8-243D-5963-A940-B74FAF6009B5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteRegistryModuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRegistryModuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRegistryModuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRegistryModuleResponseBody) SetRequestId(v string) *DeleteRegistryModuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRegistryModuleResponseBody) Validate() error {
	return dara.Validate(s)
}
