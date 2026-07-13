// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureNatGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ConfigureNatGatewayResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *ConfigureNatGatewayResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *ConfigureNatGatewayResponseBody
	GetMessage() *string
	SetRequestId(v string) *ConfigureNatGatewayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ConfigureNatGatewayResponseBody
	GetSuccess() *bool
}

type ConfigureNatGatewayResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// req-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigureNatGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigureNatGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureNatGatewayResponseBody) GetCode() *string {
	return s.Code
}

func (s *ConfigureNatGatewayResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *ConfigureNatGatewayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ConfigureNatGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigureNatGatewayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ConfigureNatGatewayResponseBody) SetCode(v string) *ConfigureNatGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *ConfigureNatGatewayResponseBody) SetData(v map[string]interface{}) *ConfigureNatGatewayResponseBody {
	s.Data = v
	return s
}

func (s *ConfigureNatGatewayResponseBody) SetMessage(v string) *ConfigureNatGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *ConfigureNatGatewayResponseBody) SetRequestId(v string) *ConfigureNatGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureNatGatewayResponseBody) SetSuccess(v bool) *ConfigureNatGatewayResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureNatGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
