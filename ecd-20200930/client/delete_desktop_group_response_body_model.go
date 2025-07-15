// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDesktopGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDesktopGroupResponseBody
	GetRequestId() *string
}

type DeleteDesktopGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDesktopGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDesktopGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDesktopGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDesktopGroupResponseBody) SetRequestId(v string) *DeleteDesktopGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDesktopGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
