// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAddonRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *GetAddonRequest
	GetAliyunLang() *string
	SetVersion(v string) *GetAddonRequest
	GetVersion() *string
}

type GetAddonRequest struct {
	// example:
	//
	// zh
	AliyunLang *string `json:"aliyunLang,omitempty" xml:"aliyunLang,omitempty"`
	// example:
	//
	// *
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetAddonRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAddonRequest) GoString() string {
	return s.String()
}

func (s *GetAddonRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *GetAddonRequest) GetVersion() *string {
	return s.Version
}

func (s *GetAddonRequest) SetAliyunLang(v string) *GetAddonRequest {
	s.AliyunLang = &v
	return s
}

func (s *GetAddonRequest) SetVersion(v string) *GetAddonRequest {
	s.Version = &v
	return s
}

func (s *GetAddonRequest) Validate() error {
	return dara.Validate(s)
}
