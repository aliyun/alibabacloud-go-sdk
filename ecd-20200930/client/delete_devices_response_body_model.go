// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDevicesResponseBody
	GetRequestId() *string
}

type DeleteDevicesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2AF861B4-0ECC-130C-B100-21A01E02****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDevicesResponseBody) SetRequestId(v string) *DeleteDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDevicesResponseBody) Validate() error {
	return dara.Validate(s)
}
