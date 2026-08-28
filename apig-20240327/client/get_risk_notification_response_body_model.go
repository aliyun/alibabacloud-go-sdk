// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRiskNotificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRiskNotificationResponseBody
	GetCode() *string
	SetData(v *GetRiskNotificationResponseBodyData) *GetRiskNotificationResponseBody
	GetData() *GetRiskNotificationResponseBodyData
	SetMessage(v string) *GetRiskNotificationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRiskNotificationResponseBody
	GetRequestId() *string
}

type GetRiskNotificationResponseBody struct {
	// example:
	//
	// 200
	Code *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetRiskNotificationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// AE1BA0DF-D730-501D-B962-B8B1C23B4667
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRiskNotificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRiskNotificationResponseBody) GoString() string {
	return s.String()
}

func (s *GetRiskNotificationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRiskNotificationResponseBody) GetData() *GetRiskNotificationResponseBodyData {
	return s.Data
}

func (s *GetRiskNotificationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRiskNotificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRiskNotificationResponseBody) SetCode(v string) *GetRiskNotificationResponseBody {
	s.Code = &v
	return s
}

func (s *GetRiskNotificationResponseBody) SetData(v *GetRiskNotificationResponseBodyData) *GetRiskNotificationResponseBody {
	s.Data = v
	return s
}

func (s *GetRiskNotificationResponseBody) SetMessage(v string) *GetRiskNotificationResponseBody {
	s.Message = &v
	return s
}

func (s *GetRiskNotificationResponseBody) SetRequestId(v string) *GetRiskNotificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRiskNotificationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRiskNotificationResponseBodyData struct {
	// example:
	//
	// gw-xxxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// false
	IsMute *bool `json:"isMute,omitempty" xml:"isMute,omitempty"`
	// example:
	//
	// GW_VERSION_EXPIRED
	RiskCode *string `json:"riskCode,omitempty" xml:"riskCode,omitempty"`
}

func (s GetRiskNotificationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRiskNotificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRiskNotificationResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetRiskNotificationResponseBodyData) GetIsMute() *bool {
	return s.IsMute
}

func (s *GetRiskNotificationResponseBodyData) GetRiskCode() *string {
	return s.RiskCode
}

func (s *GetRiskNotificationResponseBodyData) SetGatewayId(v string) *GetRiskNotificationResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetRiskNotificationResponseBodyData) SetIsMute(v bool) *GetRiskNotificationResponseBodyData {
	s.IsMute = &v
	return s
}

func (s *GetRiskNotificationResponseBodyData) SetRiskCode(v string) *GetRiskNotificationResponseBodyData {
	s.RiskCode = &v
	return s
}

func (s *GetRiskNotificationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
