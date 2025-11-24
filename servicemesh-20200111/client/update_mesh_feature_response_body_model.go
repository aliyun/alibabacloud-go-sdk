// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMeshFeatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMeshFeatureResponseBody
	GetRequestId() *string
}

type UpdateMeshFeatureResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BD65C0AD-D3C6-48D3-8D93-38D2015C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMeshFeatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeshFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMeshFeatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMeshFeatureResponseBody) SetRequestId(v string) *UpdateMeshFeatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMeshFeatureResponseBody) Validate() error {
	return dara.Validate(s)
}
