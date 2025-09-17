// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransportLayerApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTransportLayerApplicationResponseBody
	GetRequestId() *string
}

type DeleteTransportLayerApplicationResponseBody struct {
	// example:
	//
	// 90510C29-1E40-5A11-93F1-B9F5EDF57EE1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTransportLayerApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransportLayerApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTransportLayerApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTransportLayerApplicationResponseBody) SetRequestId(v string) *DeleteTransportLayerApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTransportLayerApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
