// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishLiveStagingConfigToProductionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishLiveStagingConfigToProductionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishLiveStagingConfigToProductionResponse
	GetStatusCode() *int32
	SetBody(v *PublishLiveStagingConfigToProductionResponseBody) *PublishLiveStagingConfigToProductionResponse
	GetBody() *PublishLiveStagingConfigToProductionResponseBody
}

type PublishLiveStagingConfigToProductionResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishLiveStagingConfigToProductionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishLiveStagingConfigToProductionResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishLiveStagingConfigToProductionResponse) GoString() string {
	return s.String()
}

func (s *PublishLiveStagingConfigToProductionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishLiveStagingConfigToProductionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishLiveStagingConfigToProductionResponse) GetBody() *PublishLiveStagingConfigToProductionResponseBody {
	return s.Body
}

func (s *PublishLiveStagingConfigToProductionResponse) SetHeaders(v map[string]*string) *PublishLiveStagingConfigToProductionResponse {
	s.Headers = v
	return s
}

func (s *PublishLiveStagingConfigToProductionResponse) SetStatusCode(v int32) *PublishLiveStagingConfigToProductionResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishLiveStagingConfigToProductionResponse) SetBody(v *PublishLiveStagingConfigToProductionResponseBody) *PublishLiveStagingConfigToProductionResponse {
	s.Body = v
	return s
}

func (s *PublishLiveStagingConfigToProductionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
