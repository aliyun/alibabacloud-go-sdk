// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransportLayerApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTransportLayerApplicationResponseBody
	GetRequestId() *string
}

type UpdateTransportLayerApplicationResponseBody struct {
	// example:
	//
	// 9e5448c7-edaf-49aa-9887-0fcd0832306c
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTransportLayerApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransportLayerApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTransportLayerApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTransportLayerApplicationResponseBody) SetRequestId(v string) *UpdateTransportLayerApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTransportLayerApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
