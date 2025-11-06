// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRegistrantProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRegistrantProfileResponseBody
	GetRequestId() *string
}

type DeleteRegistrantProfileResponseBody struct {
	// example:
	//
	// C50E41A0-09F1-4491-8DB8-AF55BD2D0CC8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRegistrantProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRegistrantProfileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRegistrantProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRegistrantProfileResponseBody) SetRequestId(v string) *DeleteRegistrantProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRegistrantProfileResponseBody) Validate() error {
	return dara.Validate(s)
}
