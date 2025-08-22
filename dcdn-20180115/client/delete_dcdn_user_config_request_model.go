// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnUserConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionName(v string) *DeleteDcdnUserConfigRequest
	GetFunctionName() *string
}

type DeleteDcdnUserConfigRequest struct {
	// The name of the user feature that you want to delete. Currently, only DCDN Web Application Firewall (WAF) can be deleted. Default value: waf.
	//
	// example:
	//
	// waf
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
}

func (s DeleteDcdnUserConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnUserConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteDcdnUserConfigRequest) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DeleteDcdnUserConfigRequest) SetFunctionName(v string) *DeleteDcdnUserConfigRequest {
	s.FunctionName = &v
	return s
}

func (s *DeleteDcdnUserConfigRequest) Validate() error {
	return dara.Validate(s)
}
