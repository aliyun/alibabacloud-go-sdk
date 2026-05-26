// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetValue(v string) *UpdateConfigRequest
	GetValue() *string
}

type UpdateConfigRequest struct {
	// 新的配置值
	//
	// example:
	//
	// 200
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigRequest) GetValue() *string {
	return s.Value
}

func (s *UpdateConfigRequest) SetValue(v string) *UpdateConfigRequest {
	s.Value = &v
	return s
}

func (s *UpdateConfigRequest) Validate() error {
	return dara.Validate(s)
}
