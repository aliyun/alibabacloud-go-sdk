// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParamStr(v string) *CreateOrderRequest
	GetParamStr() *string
}

type CreateOrderRequest struct {
	// This parameter is required.
	ParamStr *string `json:"paramStr,omitempty" xml:"paramStr,omitempty"`
}

func (s CreateOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderRequest) GetParamStr() *string {
	return s.ParamStr
}

func (s *CreateOrderRequest) SetParamStr(v string) *CreateOrderRequest {
	s.ParamStr = &v
	return s
}

func (s *CreateOrderRequest) Validate() error {
	return dara.Validate(s)
}
