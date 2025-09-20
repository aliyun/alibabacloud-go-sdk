// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLocationDateClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLocationDateClusterResponseBody
	GetRequestId() *string
}

type DeleteLocationDateClusterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B121940C-9794-4EE3-8D6E-F8EC525F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLocationDateClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLocationDateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLocationDateClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLocationDateClusterResponseBody) SetRequestId(v string) *DeleteLocationDateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLocationDateClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
