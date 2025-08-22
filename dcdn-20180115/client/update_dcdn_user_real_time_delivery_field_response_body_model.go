// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDcdnUserRealTimeDeliveryFieldResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDcdnUserRealTimeDeliveryFieldResponseBody
	GetRequestId() *string
}

type UpdateDcdnUserRealTimeDeliveryFieldResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 3EACD23C-F49F-4BF7-B9AD-C2CD3BA888C4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDcdnUserRealTimeDeliveryFieldResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDcdnUserRealTimeDeliveryFieldResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDcdnUserRealTimeDeliveryFieldResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDcdnUserRealTimeDeliveryFieldResponseBody) SetRequestId(v string) *UpdateDcdnUserRealTimeDeliveryFieldResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDcdnUserRealTimeDeliveryFieldResponseBody) Validate() error {
	return dara.Validate(s)
}
