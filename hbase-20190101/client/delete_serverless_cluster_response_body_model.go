// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServerlessClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteServerlessClusterResponseBody
	GetRequestId() *string
}

type DeleteServerlessClusterResponseBody struct {
	// example:
	//
	// 46950E74-59C4-4E3E-9B38-A33B*********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServerlessClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServerlessClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServerlessClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServerlessClusterResponseBody) SetRequestId(v string) *DeleteServerlessClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServerlessClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
