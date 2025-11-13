// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateClusterScannerWebhookYamlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateClusterScannerWebhookYamlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateClusterScannerWebhookYamlResponse
	GetStatusCode() *int32
	SetBody(v *GenerateClusterScannerWebhookYamlResponseBody) *GenerateClusterScannerWebhookYamlResponse
	GetBody() *GenerateClusterScannerWebhookYamlResponseBody
}

type GenerateClusterScannerWebhookYamlResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateClusterScannerWebhookYamlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateClusterScannerWebhookYamlResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateClusterScannerWebhookYamlResponse) GoString() string {
	return s.String()
}

func (s *GenerateClusterScannerWebhookYamlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateClusterScannerWebhookYamlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateClusterScannerWebhookYamlResponse) GetBody() *GenerateClusterScannerWebhookYamlResponseBody {
	return s.Body
}

func (s *GenerateClusterScannerWebhookYamlResponse) SetHeaders(v map[string]*string) *GenerateClusterScannerWebhookYamlResponse {
	s.Headers = v
	return s
}

func (s *GenerateClusterScannerWebhookYamlResponse) SetStatusCode(v int32) *GenerateClusterScannerWebhookYamlResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateClusterScannerWebhookYamlResponse) SetBody(v *GenerateClusterScannerWebhookYamlResponseBody) *GenerateClusterScannerWebhookYamlResponse {
	s.Body = v
	return s
}

func (s *GenerateClusterScannerWebhookYamlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
