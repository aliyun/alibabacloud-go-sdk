// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrometheusUserSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrometheusUserSetting(v map[string]*string) *GetPrometheusUserSettingResponseBody
	GetPrometheusUserSetting() map[string]*string
	SetRequestId(v string) *GetPrometheusUserSettingResponseBody
	GetRequestId() *string
}

type GetPrometheusUserSettingResponseBody struct {
	PrometheusUserSetting map[string]*string `json:"prometheusUserSetting,omitempty" xml:"prometheusUserSetting,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 058550FA-DDBE-519E-9C6D-93521B9A5E90
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetPrometheusUserSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusUserSettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetPrometheusUserSettingResponseBody) GetPrometheusUserSetting() map[string]*string {
	return s.PrometheusUserSetting
}

func (s *GetPrometheusUserSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPrometheusUserSettingResponseBody) SetPrometheusUserSetting(v map[string]*string) *GetPrometheusUserSettingResponseBody {
	s.PrometheusUserSetting = v
	return s
}

func (s *GetPrometheusUserSettingResponseBody) SetRequestId(v string) *GetPrometheusUserSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPrometheusUserSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
