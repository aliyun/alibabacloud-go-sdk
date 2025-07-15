// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDesktopsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDesktopsResponseBody
	GetRequestId() *string
}

type DeleteDesktopsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDesktopsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDesktopsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDesktopsResponseBody) SetRequestId(v string) *DeleteDesktopsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDesktopsResponseBody) Validate() error {
	return dara.Validate(s)
}
