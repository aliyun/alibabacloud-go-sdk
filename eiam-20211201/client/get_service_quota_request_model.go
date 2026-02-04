// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQuotaType(v string) *GetServiceQuotaRequest
	GetQuotaType() *string
}

type GetServiceQuotaRequest struct {
	// Quota 配额的唯一标识。
	//
	// This parameter is required.
	//
	// example:
	//
	// instanceTrialTimes。
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
}

func (s GetServiceQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetServiceQuotaRequest) GetQuotaType() *string {
	return s.QuotaType
}

func (s *GetServiceQuotaRequest) SetQuotaType(v string) *GetServiceQuotaRequest {
	s.QuotaType = &v
	return s
}

func (s *GetServiceQuotaRequest) Validate() error {
	return dara.Validate(s)
}
