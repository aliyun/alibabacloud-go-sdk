// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBundlesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteBundlesResponseBody
	GetRequestId() *string
}

type DeleteBundlesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBundlesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBundlesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBundlesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBundlesResponseBody) SetRequestId(v string) *DeleteBundlesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBundlesResponseBody) Validate() error {
	return dara.Validate(s)
}
