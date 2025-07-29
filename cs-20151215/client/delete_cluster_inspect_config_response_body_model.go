// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterInspectConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteClusterInspectConfigResponseBody
	GetRequestId() *string
}

type DeleteClusterInspectConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 873DC52C-28AA-5A5C-938C-684D3D4B****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteClusterInspectConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterInspectConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterInspectConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteClusterInspectConfigResponseBody) SetRequestId(v string) *DeleteClusterInspectConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClusterInspectConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
