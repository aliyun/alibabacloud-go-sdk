// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSiteDeliveryTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSiteDeliveryTaskResponseBody
	GetRequestId() *string
}

type DeleteSiteDeliveryTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSiteDeliveryTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSiteDeliveryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSiteDeliveryTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSiteDeliveryTaskResponseBody) SetRequestId(v string) *DeleteSiteDeliveryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSiteDeliveryTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
