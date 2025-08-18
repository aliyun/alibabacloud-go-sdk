// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSiteDeliveryTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSiteDeliveryTaskResponseBody
	GetRequestId() *string
}

type UpdateSiteDeliveryTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 34DCBC8A-****-****-****-6DAA11D7DDBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSiteDeliveryTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSiteDeliveryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSiteDeliveryTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSiteDeliveryTaskResponseBody) SetRequestId(v string) *UpdateSiteDeliveryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSiteDeliveryTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
