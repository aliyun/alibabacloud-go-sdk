// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetCustomerConfigResponseBody
	GetAppId() *string
	SetAuditConfig(v string) *GetCustomerConfigResponseBody
	GetAuditConfig() *string
	SetDownloadSwitch(v string) *GetCustomerConfigResponseBody
	GetDownloadSwitch() *string
	SetMetricConfig(v string) *GetCustomerConfigResponseBody
	GetMetricConfig() *string
	SetRequestId(v string) *GetCustomerConfigResponseBody
	GetRequestId() *string
}

type GetCustomerConfigResponseBody struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AuditConfig    *string `json:"AuditConfig,omitempty" xml:"AuditConfig,omitempty"`
	DownloadSwitch *string `json:"DownloadSwitch,omitempty" xml:"DownloadSwitch,omitempty"`
	MetricConfig   *string `json:"MetricConfig,omitempty" xml:"MetricConfig,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCustomerConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomerConfigResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *GetCustomerConfigResponseBody) GetAuditConfig() *string {
	return s.AuditConfig
}

func (s *GetCustomerConfigResponseBody) GetDownloadSwitch() *string {
	return s.DownloadSwitch
}

func (s *GetCustomerConfigResponseBody) GetMetricConfig() *string {
	return s.MetricConfig
}

func (s *GetCustomerConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomerConfigResponseBody) SetAppId(v string) *GetCustomerConfigResponseBody {
	s.AppId = &v
	return s
}

func (s *GetCustomerConfigResponseBody) SetAuditConfig(v string) *GetCustomerConfigResponseBody {
	s.AuditConfig = &v
	return s
}

func (s *GetCustomerConfigResponseBody) SetDownloadSwitch(v string) *GetCustomerConfigResponseBody {
	s.DownloadSwitch = &v
	return s
}

func (s *GetCustomerConfigResponseBody) SetMetricConfig(v string) *GetCustomerConfigResponseBody {
	s.MetricConfig = &v
	return s
}

func (s *GetCustomerConfigResponseBody) SetRequestId(v string) *GetCustomerConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomerConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
