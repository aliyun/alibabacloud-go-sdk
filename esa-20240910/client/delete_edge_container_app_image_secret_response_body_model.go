// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEdgeContainerAppImageSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEdgeContainerAppImageSecretResponseBody
	GetRequestId() *string
}

type DeleteEdgeContainerAppImageSecretResponseBody struct {
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEdgeContainerAppImageSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEdgeContainerAppImageSecretResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEdgeContainerAppImageSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEdgeContainerAppImageSecretResponseBody) SetRequestId(v string) *DeleteEdgeContainerAppImageSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEdgeContainerAppImageSecretResponseBody) Validate() error {
	return dara.Validate(s)
}
