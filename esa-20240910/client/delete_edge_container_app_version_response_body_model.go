// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEdgeContainerAppVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEdgeContainerAppVersionResponseBody
	GetRequestId() *string
}

type DeleteEdgeContainerAppVersionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// b021e538-9dde-46ed-a1f2-9469da8f3e77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEdgeContainerAppVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEdgeContainerAppVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEdgeContainerAppVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEdgeContainerAppVersionResponseBody) SetRequestId(v string) *DeleteEdgeContainerAppVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEdgeContainerAppVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
