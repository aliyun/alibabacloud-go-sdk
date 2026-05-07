// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVersion(v string) *GetJobTemplateRequest
	GetVersion() *string
}

type GetJobTemplateRequest struct {
	// 不传返回默认版本；传具体数字返回该版本；传 all 返回全部版本
	//
	// example:
	//
	// all
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetJobTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetJobTemplateRequest) GetVersion() *string {
	return s.Version
}

func (s *GetJobTemplateRequest) SetVersion(v string) *GetJobTemplateRequest {
	s.Version = &v
	return s
}

func (s *GetJobTemplateRequest) Validate() error {
	return dara.Validate(s)
}
