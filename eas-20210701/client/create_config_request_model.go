// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetValue(v string) *CreateConfigRequest
	GetValue() *string
}

type CreateConfigRequest struct {
	// 配置值
	//
	// example:
	//
	// {"model": "gpt-4"}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigRequest) GetValue() *string {
	return s.Value
}

func (s *CreateConfigRequest) SetValue(v string) *CreateConfigRequest {
	s.Value = &v
	return s
}

func (s *CreateConfigRequest) Validate() error {
	return dara.Validate(s)
}
