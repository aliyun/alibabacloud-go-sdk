// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudClusterAllUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetCloudClusterAllUrlResponseBody
	GetCode() *int32
	SetData(v []*GetCloudClusterAllUrlResponseBodyData) *GetCloudClusterAllUrlResponseBody
	GetData() []*GetCloudClusterAllUrlResponseBodyData
	SetMessage(v string) *GetCloudClusterAllUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCloudClusterAllUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCloudClusterAllUrlResponseBody
	GetSuccess() *bool
}

type GetCloudClusterAllUrlResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// An array object.
	Data []*GetCloudClusterAllUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 99A663CB-8D7B-4B0D-A006-03C8EE38E7BB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCloudClusterAllUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCloudClusterAllUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetCloudClusterAllUrlResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetCloudClusterAllUrlResponseBody) GetData() []*GetCloudClusterAllUrlResponseBodyData {
	return s.Data
}

func (s *GetCloudClusterAllUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCloudClusterAllUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCloudClusterAllUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCloudClusterAllUrlResponseBody) SetCode(v int32) *GetCloudClusterAllUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetCloudClusterAllUrlResponseBody) SetData(v []*GetCloudClusterAllUrlResponseBodyData) *GetCloudClusterAllUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetCloudClusterAllUrlResponseBody) SetMessage(v string) *GetCloudClusterAllUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetCloudClusterAllUrlResponseBody) SetRequestId(v string) *GetCloudClusterAllUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCloudClusterAllUrlResponseBody) SetSuccess(v bool) *GetCloudClusterAllUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetCloudClusterAllUrlResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCloudClusterAllUrlResponseBodyData struct {
	// The identifier of the cloud service.
	//
	// example:
	//
	// amp
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The URLs for remote read and write. The value is a JSON string.
	RemoteUrl *GetCloudClusterAllUrlResponseBodyDataRemoteUrl `json:"RemoteUrl,omitempty" xml:"RemoteUrl,omitempty" type:"Struct"`
}

func (s GetCloudClusterAllUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCloudClusterAllUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCloudClusterAllUrlResponseBodyData) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetCloudClusterAllUrlResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *GetCloudClusterAllUrlResponseBodyData) GetRemoteUrl() *GetCloudClusterAllUrlResponseBodyDataRemoteUrl {
	return s.RemoteUrl
}

func (s *GetCloudClusterAllUrlResponseBodyData) SetProductCode(v string) *GetCloudClusterAllUrlResponseBodyData {
	s.ProductCode = &v
	return s
}

func (s *GetCloudClusterAllUrlResponseBodyData) SetRegion(v string) *GetCloudClusterAllUrlResponseBodyData {
	s.Region = &v
	return s
}

func (s *GetCloudClusterAllUrlResponseBodyData) SetRemoteUrl(v *GetCloudClusterAllUrlResponseBodyDataRemoteUrl) *GetCloudClusterAllUrlResponseBodyData {
	s.RemoteUrl = v
	return s
}

func (s *GetCloudClusterAllUrlResponseBodyData) Validate() error {
	if s.RemoteUrl != nil {
		if err := s.RemoteUrl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCloudClusterAllUrlResponseBodyDataRemoteUrl struct {
	// Indicates whether authentication is enabled.
	//
	// example:
	//
	// true
	AuthToken *bool `json:"AuthToken,omitempty" xml:"AuthToken,omitempty"`
	// The internal URL for Grafana.
	//
	// example:
	//
	// "http://cn-hangzhou-intranet.arms.aliyuncs.com:9090/api/v1/prometheus/XXXXXXXXXXXXXXXXX"
	GrafanaUrl *string `json:"GrafanaUrl,omitempty" xml:"GrafanaUrl,omitempty"`
	// The public URL for Grafana.
	//
	// example:
	//
	// "http://cn-hangzhou.arms.aliyuncs.com:9090/api/v1/prometheus/XXXXXXXXXXXXXXXXX"
	InternetGrafanaUrl *string `json:"InternetGrafanaUrl,omitempty" xml:"InternetGrafanaUrl,omitempty"`
	// The public URL for Pushgateway.
	//
	// example:
	//
	// "http://cn-hangzhou.arms.aliyuncs.com:9090/api/v1/prometheus/XXXXXXXXXXXXXXXXX"
	InternetPushGatewayUrl *string `json:"InternetPushGatewayUrl,omitempty" xml:"InternetPushGatewayUrl,omitempty"`
	// The public URL for remote read.
	//
	// example:
	//
	// "http://cn-hangzhou.arms.aliyuncs.com:9090/api/v1/prometheus/XXXXXXXXXXXXXXXXX"
	InternetRemoteReadUrl *string `json:"InternetRemoteReadUrl,omitempty" xml:"InternetRemoteReadUrl,omitempty"`
	// The public URL for remote write.
	//
	// example:
	//
	// "http://cn-hangzhou.arms.aliyuncs.com:9090/api/v1/prometheus/XXXXXXXXXXXXXXXXX"
	InternetRemoteWriteUrl *string `json:"InternetRemoteWriteUrl,omitempty" xml:"InternetRemoteWriteUrl,omitempty"`
	// The internal URL for Pushgateway.
	//
	// example:
	//
	// "http://cn-hangzhou-intranet.arms.aliyuncs.com:9090/api/v1/prometheus/XXXXXXXXXXXXXXXXX"
	PushGatewayUrl *string `json:"PushGatewayUrl,omitempty" xml:"PushGatewayUrl,omitempty"`
	// The internal URL for remote read.
	//
	// example:
	//
	// "http://cn-hangzhou-intranet.arms.aliyuncs.com:9090/api/v1/prometheus/XXXXXXXXXXXXXXXXX"
	RemoteReadUrl *string `json:"RemoteReadUrl,omitempty" xml:"RemoteReadUrl,omitempty"`
	// The internal URL for remote write.
	//
	// example:
	//
	// "http://cn-hangzhou-intranet.arms.aliyuncs.com:9090/api/v1/prometheus/XXXXXXXXXXXXXXXXX"
	RemoteWriteUrl *string `json:"RemoteWriteUrl,omitempty" xml:"RemoteWriteUrl,omitempty"`
	// The token value used for authentication.
	//
	// example:
	//
	// "eyJhbGciOiJIUzI1NiJ9.DKEIFJSL.KYK6uOtNVxTVHXJbH5MNqlsAuUtKzNlUvmAIiKc-QXw"
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetCloudClusterAllUrlResponseBodyDataRemoteUrl) String() string {
	return dara.Prettify(s)
}

func (s GetCloudClusterAllUrlResponseBodyDataRemoteUrl) GoString() string {
	return s.String()
}

func (s *GetCloudClusterAllUrlResponseBodyDataRemoteUrl) GetAuthToken() *bool {
	return s.AuthToken
}

func (s *GetCloudClusterAllUrlResponseBodyDataRemoteUrl) GetGrafanaUrl() *string {
	return s.GrafanaUrl
}

func (s *GetCloudClusterAllUrlResponseBodyDataRemoteUrl) GetInternetGrafanaUrl() *string {
	return s.InternetGrafanaUrl
}

func (s *GetCloudClusterAllUrlResponseBodyDataRemoteUrl) GetInternetPushGatewayUrl() *string {
	return s.InternetPushGatewayUrl
}

func (s *GetCloudClusterAllUrlResponseBodyDataRemoteUrl) GetInternetRemoteReadUrl() *string {
	return s.InternetRemoteReadUrl
}

func (s *GetCloudClusterAllUrlResponseBodyDataRemoteUrl) GetInternetRemoteWriteUrl() *string {
	return s.InternetRemoteWriteUrl
}

func (s *GetCloudClusterAllUrlResponseBodyDataRemoteUrl) GetPushGatewayUrl() *string {
	return s.PushGatewayUrl
}

func (s *GetCloudClusterAllUrlResponseBodyDataRemoteUrl) GetRemoteReadUrl() *string {
	return s.RemoteReadUrl
}

func (s *GetCloudClusterAllUrlResponseBodyDataRemoteUrl) GetRemoteWriteUrl() *string {
	return s.RemoteWriteUrl
}

func (s *GetCloudClusterAllUrlResponseBodyDataRemoteUrl) GetToken() *string {
	return s.Token
}

func (s *GetCloudClusterAllUrlResponseBodyDataRemoteUrl) SetAuthToken(v bool) *GetCloudClusterAllUrlResponseBodyDataRemoteUrl {
	s.AuthToken = &v
	return s
}

func (s *GetCloudClusterAllUrlResponseBodyDataRemoteUrl) SetGrafanaUrl(v string) *GetCloudClusterAllUrlResponseBodyDataRemoteUrl {
	s.GrafanaUrl = &v
	return s
}

func (s *GetCloudClusterAllUrlResponseBodyDataRemoteUrl) SetInternetGrafanaUrl(v string) *GetCloudClusterAllUrlResponseBodyDataRemoteUrl {
	s.InternetGrafanaUrl = &v
	return s
}

func (s *GetCloudClusterAllUrlResponseBodyDataRemoteUrl) SetInternetPushGatewayUrl(v string) *GetCloudClusterAllUrlResponseBodyDataRemoteUrl {
	s.InternetPushGatewayUrl = &v
	return s
}

func (s *GetCloudClusterAllUrlResponseBodyDataRemoteUrl) SetInternetRemoteReadUrl(v string) *GetCloudClusterAllUrlResponseBodyDataRemoteUrl {
	s.InternetRemoteReadUrl = &v
	return s
}

func (s *GetCloudClusterAllUrlResponseBodyDataRemoteUrl) SetInternetRemoteWriteUrl(v string) *GetCloudClusterAllUrlResponseBodyDataRemoteUrl {
	s.InternetRemoteWriteUrl = &v
	return s
}

func (s *GetCloudClusterAllUrlResponseBodyDataRemoteUrl) SetPushGatewayUrl(v string) *GetCloudClusterAllUrlResponseBodyDataRemoteUrl {
	s.PushGatewayUrl = &v
	return s
}

func (s *GetCloudClusterAllUrlResponseBodyDataRemoteUrl) SetRemoteReadUrl(v string) *GetCloudClusterAllUrlResponseBodyDataRemoteUrl {
	s.RemoteReadUrl = &v
	return s
}

func (s *GetCloudClusterAllUrlResponseBodyDataRemoteUrl) SetRemoteWriteUrl(v string) *GetCloudClusterAllUrlResponseBodyDataRemoteUrl {
	s.RemoteWriteUrl = &v
	return s
}

func (s *GetCloudClusterAllUrlResponseBodyDataRemoteUrl) SetToken(v string) *GetCloudClusterAllUrlResponseBodyDataRemoteUrl {
	s.Token = &v
	return s
}

func (s *GetCloudClusterAllUrlResponseBodyDataRemoteUrl) Validate() error {
	return dara.Validate(s)
}
