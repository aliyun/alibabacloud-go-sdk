// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterScannerYamlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCaCertBase64(v string) *GetClusterScannerYamlResponseBody
	GetCaCertBase64() *string
	SetClusterEnvInfo(v string) *GetClusterScannerYamlResponseBody
	GetClusterEnvInfo() *string
	SetClusterId(v string) *GetClusterScannerYamlResponseBody
	GetClusterId() *string
	SetImage(v string) *GetClusterScannerYamlResponseBody
	GetImage() *string
	SetRequestId(v string) *GetClusterScannerYamlResponseBody
	GetRequestId() *string
	SetTlsCertBase64(v string) *GetClusterScannerYamlResponseBody
	GetTlsCertBase64() *string
	SetTlsKeyBase64(v string) *GetClusterScannerYamlResponseBody
	GetTlsKeyBase64() *string
	SetWebhookOpen(v int32) *GetClusterScannerYamlResponseBody
	GetWebhookOpen() *int32
}

type GetClusterScannerYamlResponseBody struct {
	// The CA certificate in Base64 encoding.
	//
	// example:
	//
	// xxx
	CaCertBase64 *string `json:"CaCertBase64,omitempty" xml:"CaCertBase64,omitempty"`
	// The cluster environment context.
	//
	// example:
	//
	// xxx
	ClusterEnvInfo *string `json:"ClusterEnvInfo,omitempty" xml:"ClusterEnvInfo,omitempty"`
	// The ID of the container cluster.
	//
	// example:
	//
	// c7c190a82d9a048be9038d352840f****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The container image information.
	//
	// example:
	//
	// xxxx-registry.cn-shanghai.cr.aliyuncs.com/default/scanner:v1
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The request ID. Alibaba Cloud generates a unique ID for each request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 24A20733-10A0-4AF6-BE6B-E3322413BB68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The webhook certificate in Base64 encoding.
	//
	// example:
	//
	// xxx
	TlsCertBase64 *string `json:"TlsCertBase64,omitempty" xml:"TlsCertBase64,omitempty"`
	// The webhook private key in Base64 encoding.
	//
	// example:
	//
	// xxx
	TlsKeyBase64 *string `json:"TlsKeyBase64,omitempty" xml:"TlsKeyBase64,omitempty"`
	// Indicates whether incremental scanning is enabled. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
	//
	// example:
	//
	// 1
	WebhookOpen *int32 `json:"WebhookOpen,omitempty" xml:"WebhookOpen,omitempty"`
}

func (s GetClusterScannerYamlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClusterScannerYamlResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterScannerYamlResponseBody) GetCaCertBase64() *string {
	return s.CaCertBase64
}

func (s *GetClusterScannerYamlResponseBody) GetClusterEnvInfo() *string {
	return s.ClusterEnvInfo
}

func (s *GetClusterScannerYamlResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetClusterScannerYamlResponseBody) GetImage() *string {
	return s.Image
}

func (s *GetClusterScannerYamlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClusterScannerYamlResponseBody) GetTlsCertBase64() *string {
	return s.TlsCertBase64
}

func (s *GetClusterScannerYamlResponseBody) GetTlsKeyBase64() *string {
	return s.TlsKeyBase64
}

func (s *GetClusterScannerYamlResponseBody) GetWebhookOpen() *int32 {
	return s.WebhookOpen
}

func (s *GetClusterScannerYamlResponseBody) SetCaCertBase64(v string) *GetClusterScannerYamlResponseBody {
	s.CaCertBase64 = &v
	return s
}

func (s *GetClusterScannerYamlResponseBody) SetClusterEnvInfo(v string) *GetClusterScannerYamlResponseBody {
	s.ClusterEnvInfo = &v
	return s
}

func (s *GetClusterScannerYamlResponseBody) SetClusterId(v string) *GetClusterScannerYamlResponseBody {
	s.ClusterId = &v
	return s
}

func (s *GetClusterScannerYamlResponseBody) SetImage(v string) *GetClusterScannerYamlResponseBody {
	s.Image = &v
	return s
}

func (s *GetClusterScannerYamlResponseBody) SetRequestId(v string) *GetClusterScannerYamlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterScannerYamlResponseBody) SetTlsCertBase64(v string) *GetClusterScannerYamlResponseBody {
	s.TlsCertBase64 = &v
	return s
}

func (s *GetClusterScannerYamlResponseBody) SetTlsKeyBase64(v string) *GetClusterScannerYamlResponseBody {
	s.TlsKeyBase64 = &v
	return s
}

func (s *GetClusterScannerYamlResponseBody) SetWebhookOpen(v int32) *GetClusterScannerYamlResponseBody {
	s.WebhookOpen = &v
	return s
}

func (s *GetClusterScannerYamlResponseBody) Validate() error {
	return dara.Validate(s)
}
