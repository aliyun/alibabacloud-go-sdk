// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEdgeContainerAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEdgeContainerAppResponseBody
	GetRequestId() *string
	SetState(v string) *DeleteEdgeContainerAppResponseBody
	GetState() *string
}

type DeleteEdgeContainerAppResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Specifies whether the deletion is successful.
	//
	// example:
	//
	// ok
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DeleteEdgeContainerAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEdgeContainerAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEdgeContainerAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEdgeContainerAppResponseBody) GetState() *string {
	return s.State
}

func (s *DeleteEdgeContainerAppResponseBody) SetRequestId(v string) *DeleteEdgeContainerAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEdgeContainerAppResponseBody) SetState(v string) *DeleteEdgeContainerAppResponseBody {
	s.State = &v
	return s
}

func (s *DeleteEdgeContainerAppResponseBody) Validate() error {
	return dara.Validate(s)
}
