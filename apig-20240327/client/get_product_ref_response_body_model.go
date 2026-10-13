// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductRefResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetProductRefResponseBody
	GetCode() *string
	SetData(v *GetProductRefResponseBodyData) *GetProductRefResponseBody
	GetData() *GetProductRefResponseBodyData
	SetMessage(v string) *GetProductRefResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetProductRefResponseBody
	GetRequestId() *string
}

type GetProductRefResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *GetProductRefResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetProductRefResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProductRefResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductRefResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetProductRefResponseBody) GetData() *GetProductRefResponseBodyData {
	return s.Data
}

func (s *GetProductRefResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetProductRefResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProductRefResponseBody) SetCode(v string) *GetProductRefResponseBody {
	s.Code = &v
	return s
}

func (s *GetProductRefResponseBody) SetData(v *GetProductRefResponseBodyData) *GetProductRefResponseBody {
	s.Data = v
	return s
}

func (s *GetProductRefResponseBody) SetMessage(v string) *GetProductRefResponseBody {
	s.Message = &v
	return s
}

func (s *GetProductRefResponseBody) SetRequestId(v string) *GetProductRefResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProductRefResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProductRefResponseBodyData struct {
	// The API configurations currently associated with the API product. The structure is derived from the common struct HiMarketProductRefConfig. This struct contains the gatewayId and apigRefConfig members. The inner apigRefConfig struct contains the mcpServerId, mcpRouteId, mcpServerName, agentApiId, agentApiName, modelApiId, and modelApiName members.
	ApigRefConfig *HiMarketProductRefConfig `json:"apigRefConfig,omitempty" xml:"apigRefConfig,omitempty"`
	// The gateway ID. The ID of the gateway instance where the association configuration of the API product resides.
	//
	// example:
	//
	// gw-123456
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
}

func (s GetProductRefResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetProductRefResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProductRefResponseBodyData) GetApigRefConfig() *HiMarketProductRefConfig {
	return s.ApigRefConfig
}

func (s *GetProductRefResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetProductRefResponseBodyData) SetApigRefConfig(v *HiMarketProductRefConfig) *GetProductRefResponseBodyData {
	s.ApigRefConfig = v
	return s
}

func (s *GetProductRefResponseBodyData) SetGatewayId(v string) *GetProductRefResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetProductRefResponseBodyData) Validate() error {
	if s.ApigRefConfig != nil {
		if err := s.ApigRefConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
