// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetJobTemplateDefaultVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVersion(v int32) *SetJobTemplateDefaultVersionRequest
	GetVersion() *int32
}

type SetJobTemplateDefaultVersionRequest struct {
	// 要设置为默认的模板版本号
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s SetJobTemplateDefaultVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s SetJobTemplateDefaultVersionRequest) GoString() string {
	return s.String()
}

func (s *SetJobTemplateDefaultVersionRequest) GetVersion() *int32 {
	return s.Version
}

func (s *SetJobTemplateDefaultVersionRequest) SetVersion(v int32) *SetJobTemplateDefaultVersionRequest {
	s.Version = &v
	return s
}

func (s *SetJobTemplateDefaultVersionRequest) Validate() error {
	return dara.Validate(s)
}
