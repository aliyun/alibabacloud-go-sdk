// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogDeliveryConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLogDeliveryConfigResponseBody
	GetRequestId() *string
}

type DeleteLogDeliveryConfigResponseBody struct {
	// example:
	//
	// D7861F61-5B61-46CE-A47C-*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLogDeliveryConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogDeliveryConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLogDeliveryConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLogDeliveryConfigResponseBody) SetRequestId(v string) *DeleteLogDeliveryConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLogDeliveryConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
