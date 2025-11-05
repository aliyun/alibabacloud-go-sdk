// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishStagingConfigToProductionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishStagingConfigToProductionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishStagingConfigToProductionResponse
	GetStatusCode() *int32
	SetBody(v *PublishStagingConfigToProductionResponseBody) *PublishStagingConfigToProductionResponse
	GetBody() *PublishStagingConfigToProductionResponseBody
}

type PublishStagingConfigToProductionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishStagingConfigToProductionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishStagingConfigToProductionResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishStagingConfigToProductionResponse) GoString() string {
	return s.String()
}

func (s *PublishStagingConfigToProductionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishStagingConfigToProductionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishStagingConfigToProductionResponse) GetBody() *PublishStagingConfigToProductionResponseBody {
	return s.Body
}

func (s *PublishStagingConfigToProductionResponse) SetHeaders(v map[string]*string) *PublishStagingConfigToProductionResponse {
	s.Headers = v
	return s
}

func (s *PublishStagingConfigToProductionResponse) SetStatusCode(v int32) *PublishStagingConfigToProductionResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishStagingConfigToProductionResponse) SetBody(v *PublishStagingConfigToProductionResponseBody) *PublishStagingConfigToProductionResponse {
	s.Body = v
	return s
}

func (s *PublishStagingConfigToProductionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
