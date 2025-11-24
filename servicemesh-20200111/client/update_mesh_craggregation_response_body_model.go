// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMeshCRAggregationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMeshCRAggregationResponseBody
	GetRequestId() *string
}

type UpdateMeshCRAggregationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BD65C0AD-D3C6-48D3-8D93-38D2015C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMeshCRAggregationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeshCRAggregationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMeshCRAggregationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMeshCRAggregationResponseBody) SetRequestId(v string) *UpdateMeshCRAggregationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMeshCRAggregationResponseBody) Validate() error {
	return dara.Validate(s)
}
