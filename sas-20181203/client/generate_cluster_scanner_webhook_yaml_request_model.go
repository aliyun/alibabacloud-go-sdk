// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateClusterScannerWebhookYamlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GenerateClusterScannerWebhookYamlRequest
	GetClusterId() *string
	SetWebhookOpen(v int32) *GenerateClusterScannerWebhookYamlRequest
	GetWebhookOpen() *int32
}

type GenerateClusterScannerWebhookYamlRequest struct {
	// The ID of the container cluster.
	//
	// > You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) API to obtain this parameter from the ClusterId field.
	//
	// This parameter is required.
	//
	// example:
	//
	// c3aaf6c8085f84791882eef200cd2****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Indicates whether the incremental scan switch is enabled. Values:
	//
	// - **0**: Not enabled
	//
	// - **1**: Enabled
	//
	// example:
	//
	// 1
	WebhookOpen *int32 `json:"WebhookOpen,omitempty" xml:"WebhookOpen,omitempty"`
}

func (s GenerateClusterScannerWebhookYamlRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateClusterScannerWebhookYamlRequest) GoString() string {
	return s.String()
}

func (s *GenerateClusterScannerWebhookYamlRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GenerateClusterScannerWebhookYamlRequest) GetWebhookOpen() *int32 {
	return s.WebhookOpen
}

func (s *GenerateClusterScannerWebhookYamlRequest) SetClusterId(v string) *GenerateClusterScannerWebhookYamlRequest {
	s.ClusterId = &v
	return s
}

func (s *GenerateClusterScannerWebhookYamlRequest) SetWebhookOpen(v int32) *GenerateClusterScannerWebhookYamlRequest {
	s.WebhookOpen = &v
	return s
}

func (s *GenerateClusterScannerWebhookYamlRequest) Validate() error {
	return dara.Validate(s)
}
