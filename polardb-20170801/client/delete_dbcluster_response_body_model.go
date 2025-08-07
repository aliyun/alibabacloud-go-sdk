// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDBClusterResponseBody
	GetRequestId() *string
}

type DeleteDBClusterResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD86******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBClusterResponseBody) SetRequestId(v string) *DeleteDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
