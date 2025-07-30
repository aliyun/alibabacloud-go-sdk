// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBClusterBindingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDBClusterBindingResponseBody
	GetRequestId() *string
}

type DeleteDBClusterBindingResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F8900A96-67F7-5274-A41B-7722E1ECF8C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBClusterBindingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBClusterBindingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterBindingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBClusterBindingResponseBody) SetRequestId(v string) *DeleteDBClusterBindingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBClusterBindingResponseBody) Validate() error {
	return dara.Validate(s)
}
