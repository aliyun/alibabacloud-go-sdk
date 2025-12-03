// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryForCssOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParamStr(v string) *QueryForCssOrderRequest
	GetParamStr() *string
}

type QueryForCssOrderRequest struct {
	// This parameter is required.
	ParamStr *string `json:"paramStr,omitempty" xml:"paramStr,omitempty"`
}

func (s QueryForCssOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryForCssOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryForCssOrderRequest) GetParamStr() *string {
	return s.ParamStr
}

func (s *QueryForCssOrderRequest) SetParamStr(v string) *QueryForCssOrderRequest {
	s.ParamStr = &v
	return s
}

func (s *QueryForCssOrderRequest) Validate() error {
	return dara.Validate(s)
}
