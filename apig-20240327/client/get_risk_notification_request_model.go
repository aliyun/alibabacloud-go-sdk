// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRiskNotificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRiskCode(v string) *GetRiskNotificationRequest
	GetRiskCode() *string
}

type GetRiskNotificationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// GW_VERSION_EXPIRED
	RiskCode *string `json:"riskCode,omitempty" xml:"riskCode,omitempty"`
}

func (s GetRiskNotificationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRiskNotificationRequest) GoString() string {
	return s.String()
}

func (s *GetRiskNotificationRequest) GetRiskCode() *string {
	return s.RiskCode
}

func (s *GetRiskNotificationRequest) SetRiskCode(v string) *GetRiskNotificationRequest {
	s.RiskCode = &v
	return s
}

func (s *GetRiskNotificationRequest) Validate() error {
	return dara.Validate(s)
}
