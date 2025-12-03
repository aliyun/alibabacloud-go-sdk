// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenCallbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParamStr(v string) *OpenCallbackRequest
	GetParamStr() *string
}

type OpenCallbackRequest struct {
	// This parameter is required.
	ParamStr *string `json:"paramStr,omitempty" xml:"paramStr,omitempty"`
}

func (s OpenCallbackRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenCallbackRequest) GoString() string {
	return s.String()
}

func (s *OpenCallbackRequest) GetParamStr() *string {
	return s.ParamStr
}

func (s *OpenCallbackRequest) SetParamStr(v string) *OpenCallbackRequest {
	s.ParamStr = &v
	return s
}

func (s *OpenCallbackRequest) Validate() error {
	return dara.Validate(s)
}
