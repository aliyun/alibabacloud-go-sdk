// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNamespaceScopeSidecarConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateNamespaceScopeSidecarConfigResponseBody
	GetRequestId() *string
}

type UpdateNamespaceScopeSidecarConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 31d3a0f0-07ed-4f6e-9004-1804498c****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNamespaceScopeSidecarConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNamespaceScopeSidecarConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceScopeSidecarConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNamespaceScopeSidecarConfigResponseBody) SetRequestId(v string) *UpdateNamespaceScopeSidecarConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNamespaceScopeSidecarConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
