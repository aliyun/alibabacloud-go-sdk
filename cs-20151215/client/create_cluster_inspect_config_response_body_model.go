// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterInspectConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateClusterInspectConfigResponseBody
	GetRequestId() *string
}

type CreateClusterInspectConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 873DC52C-28AA-5A5C-938C-684D3D4B****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateClusterInspectConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterInspectConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterInspectConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateClusterInspectConfigResponseBody) SetRequestId(v string) *CreateClusterInspectConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterInspectConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
