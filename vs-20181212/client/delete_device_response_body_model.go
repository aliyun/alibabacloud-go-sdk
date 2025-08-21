// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDeviceResponseBody
	GetRequestId() *string
}

type DeleteDeviceResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDeviceResponseBody) SetRequestId(v string) *DeleteDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
