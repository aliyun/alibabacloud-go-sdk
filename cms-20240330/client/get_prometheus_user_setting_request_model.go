// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrometheusUserSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *GetPrometheusUserSettingRequest
	GetAliyunLang() *string
}

type GetPrometheusUserSettingRequest struct {
	// example:
	//
	// zh
	AliyunLang *string `json:"aliyunLang,omitempty" xml:"aliyunLang,omitempty"`
}

func (s GetPrometheusUserSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusUserSettingRequest) GoString() string {
	return s.String()
}

func (s *GetPrometheusUserSettingRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *GetPrometheusUserSettingRequest) SetAliyunLang(v string) *GetPrometheusUserSettingRequest {
	s.AliyunLang = &v
	return s
}

func (s *GetPrometheusUserSettingRequest) Validate() error {
	return dara.Validate(s)
}
