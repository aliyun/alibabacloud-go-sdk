// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateClusterScannerWebhookYamlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GenerateClusterScannerWebhookYamlResponseBody
	GetClusterId() *string
	SetId(v int64) *GenerateClusterScannerWebhookYamlResponseBody
	GetId() *int64
	SetRequestId(v string) *GenerateClusterScannerWebhookYamlResponseBody
	GetRequestId() *string
	SetWebhookOpen(v int32) *GenerateClusterScannerWebhookYamlResponseBody
	GetWebhookOpen() *int32
}

type GenerateClusterScannerWebhookYamlResponseBody struct {
	// example:
	//
	// c471f0f61b9c04f8380556e922cf1****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 131231
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// BE120DAB-F4E7-4C53-ADC3-A97578ABF384
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	WebhookOpen *int32 `json:"WebhookOpen,omitempty" xml:"WebhookOpen,omitempty"`
}

func (s GenerateClusterScannerWebhookYamlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateClusterScannerWebhookYamlResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateClusterScannerWebhookYamlResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *GenerateClusterScannerWebhookYamlResponseBody) GetId() *int64 {
	return s.Id
}

func (s *GenerateClusterScannerWebhookYamlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateClusterScannerWebhookYamlResponseBody) GetWebhookOpen() *int32 {
	return s.WebhookOpen
}

func (s *GenerateClusterScannerWebhookYamlResponseBody) SetClusterId(v string) *GenerateClusterScannerWebhookYamlResponseBody {
	s.ClusterId = &v
	return s
}

func (s *GenerateClusterScannerWebhookYamlResponseBody) SetId(v int64) *GenerateClusterScannerWebhookYamlResponseBody {
	s.Id = &v
	return s
}

func (s *GenerateClusterScannerWebhookYamlResponseBody) SetRequestId(v string) *GenerateClusterScannerWebhookYamlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateClusterScannerWebhookYamlResponseBody) SetWebhookOpen(v int32) *GenerateClusterScannerWebhookYamlResponseBody {
	s.WebhookOpen = &v
	return s
}

func (s *GenerateClusterScannerWebhookYamlResponseBody) Validate() error {
	return dara.Validate(s)
}
