// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserLogDeliveryQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *GetUserLogDeliveryQuotaRequest
	GetBusinessType() *string
}

type GetUserLogDeliveryQuotaRequest struct {
	// This parameter is required.
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
}

func (s GetUserLogDeliveryQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserLogDeliveryQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetUserLogDeliveryQuotaRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *GetUserLogDeliveryQuotaRequest) SetBusinessType(v string) *GetUserLogDeliveryQuotaRequest {
	s.BusinessType = &v
	return s
}

func (s *GetUserLogDeliveryQuotaRequest) Validate() error {
	return dara.Validate(s)
}
