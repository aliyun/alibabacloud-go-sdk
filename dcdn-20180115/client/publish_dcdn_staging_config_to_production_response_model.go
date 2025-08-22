// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishDcdnStagingConfigToProductionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishDcdnStagingConfigToProductionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishDcdnStagingConfigToProductionResponse
	GetStatusCode() *int32
	SetBody(v *PublishDcdnStagingConfigToProductionResponseBody) *PublishDcdnStagingConfigToProductionResponse
	GetBody() *PublishDcdnStagingConfigToProductionResponseBody
}

type PublishDcdnStagingConfigToProductionResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishDcdnStagingConfigToProductionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishDcdnStagingConfigToProductionResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishDcdnStagingConfigToProductionResponse) GoString() string {
	return s.String()
}

func (s *PublishDcdnStagingConfigToProductionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishDcdnStagingConfigToProductionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishDcdnStagingConfigToProductionResponse) GetBody() *PublishDcdnStagingConfigToProductionResponseBody {
	return s.Body
}

func (s *PublishDcdnStagingConfigToProductionResponse) SetHeaders(v map[string]*string) *PublishDcdnStagingConfigToProductionResponse {
	s.Headers = v
	return s
}

func (s *PublishDcdnStagingConfigToProductionResponse) SetStatusCode(v int32) *PublishDcdnStagingConfigToProductionResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishDcdnStagingConfigToProductionResponse) SetBody(v *PublishDcdnStagingConfigToProductionResponseBody) *PublishDcdnStagingConfigToProductionResponse {
	s.Body = v
	return s
}

func (s *PublishDcdnStagingConfigToProductionResponse) Validate() error {
	return dara.Validate(s)
}
