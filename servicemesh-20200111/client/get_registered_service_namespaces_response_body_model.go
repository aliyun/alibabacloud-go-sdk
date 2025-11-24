// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegisteredServiceNamespacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaces(v []*string) *GetRegisteredServiceNamespacesResponseBody
	GetNamespaces() []*string
	SetRequestId(v string) *GetRegisteredServiceNamespacesResponseBody
	GetRequestId() *string
}

type GetRegisteredServiceNamespacesResponseBody struct {
	// The names of the queried namespaces.
	Namespaces []*string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 31d3a0f0-07ed-4f6e-9004-1804498c****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRegisteredServiceNamespacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRegisteredServiceNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *GetRegisteredServiceNamespacesResponseBody) GetNamespaces() []*string {
	return s.Namespaces
}

func (s *GetRegisteredServiceNamespacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRegisteredServiceNamespacesResponseBody) SetNamespaces(v []*string) *GetRegisteredServiceNamespacesResponseBody {
	s.Namespaces = v
	return s
}

func (s *GetRegisteredServiceNamespacesResponseBody) SetRequestId(v string) *GetRegisteredServiceNamespacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRegisteredServiceNamespacesResponseBody) Validate() error {
	return dara.Validate(s)
}
