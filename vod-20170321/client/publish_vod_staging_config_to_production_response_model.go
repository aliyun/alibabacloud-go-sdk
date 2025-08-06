// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishVodStagingConfigToProductionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishVodStagingConfigToProductionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishVodStagingConfigToProductionResponse
	GetStatusCode() *int32
	SetBody(v *PublishVodStagingConfigToProductionResponseBody) *PublishVodStagingConfigToProductionResponse
	GetBody() *PublishVodStagingConfigToProductionResponseBody
}

type PublishVodStagingConfigToProductionResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishVodStagingConfigToProductionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishVodStagingConfigToProductionResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishVodStagingConfigToProductionResponse) GoString() string {
	return s.String()
}

func (s *PublishVodStagingConfigToProductionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishVodStagingConfigToProductionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishVodStagingConfigToProductionResponse) GetBody() *PublishVodStagingConfigToProductionResponseBody {
	return s.Body
}

func (s *PublishVodStagingConfigToProductionResponse) SetHeaders(v map[string]*string) *PublishVodStagingConfigToProductionResponse {
	s.Headers = v
	return s
}

func (s *PublishVodStagingConfigToProductionResponse) SetStatusCode(v int32) *PublishVodStagingConfigToProductionResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishVodStagingConfigToProductionResponse) SetBody(v *PublishVodStagingConfigToProductionResponseBody) *PublishVodStagingConfigToProductionResponse {
	s.Body = v
	return s
}

func (s *PublishVodStagingConfigToProductionResponse) Validate() error {
	return dara.Validate(s)
}
