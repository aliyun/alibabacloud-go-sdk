// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecursionZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRecursionZoneResponseBody
	GetRequestId() *string
}

type DeleteRecursionZoneResponseBody struct {
	// example:
	//
	// 389DFFA3-77A5-4A9E-BF3D-147C6F98A5BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRecursionZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecursionZoneResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRecursionZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRecursionZoneResponseBody) SetRequestId(v string) *DeleteRecursionZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRecursionZoneResponseBody) Validate() error {
	return dara.Validate(s)
}
