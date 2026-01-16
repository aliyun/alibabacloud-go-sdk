// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterNodePoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteClusterNodePoolResponseBody
	GetRequestId() *string
}

type DeleteClusterNodePoolResponseBody struct {
	// example:
	//
	// F3B261DD-3858-4D3C-877D-303ADF374600
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClusterNodePoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterNodePoolResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodePoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteClusterNodePoolResponseBody) SetRequestId(v string) *DeleteClusterNodePoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClusterNodePoolResponseBody) Validate() error {
	return dara.Validate(s)
}
