// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteClusterNodesResponseBody
	GetRequestId() *string
}

type DeleteClusterNodesResponseBody struct {
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClusterNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteClusterNodesResponseBody) SetRequestId(v string) *DeleteClusterNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClusterNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
