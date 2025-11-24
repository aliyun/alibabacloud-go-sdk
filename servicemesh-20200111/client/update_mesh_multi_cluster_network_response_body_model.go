// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMeshMultiClusterNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMeshMultiClusterNetworkResponseBody
	GetRequestId() *string
}

type UpdateMeshMultiClusterNetworkResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// BD65C0AD-D3C6-48D3-8D93-38D2015C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMeshMultiClusterNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeshMultiClusterNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMeshMultiClusterNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMeshMultiClusterNetworkResponseBody) SetRequestId(v string) *UpdateMeshMultiClusterNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMeshMultiClusterNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}
