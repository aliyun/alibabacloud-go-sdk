// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVnoBatchRefundOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParamStr(v string) *VnoBatchRefundOrderRequest
	GetParamStr() *string
}

type VnoBatchRefundOrderRequest struct {
	// This parameter is required.
	ParamStr *string `json:"paramStr,omitempty" xml:"paramStr,omitempty"`
}

func (s VnoBatchRefundOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s VnoBatchRefundOrderRequest) GoString() string {
	return s.String()
}

func (s *VnoBatchRefundOrderRequest) GetParamStr() *string {
	return s.ParamStr
}

func (s *VnoBatchRefundOrderRequest) SetParamStr(v string) *VnoBatchRefundOrderRequest {
	s.ParamStr = &v
	return s
}

func (s *VnoBatchRefundOrderRequest) Validate() error {
	return dara.Validate(s)
}
