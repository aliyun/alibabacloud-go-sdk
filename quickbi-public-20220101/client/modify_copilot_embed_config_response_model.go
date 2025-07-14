// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCopilotEmbedConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCopilotEmbedConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCopilotEmbedConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCopilotEmbedConfigResponseBody) *ModifyCopilotEmbedConfigResponse
	GetBody() *ModifyCopilotEmbedConfigResponseBody
}

type ModifyCopilotEmbedConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCopilotEmbedConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCopilotEmbedConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCopilotEmbedConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyCopilotEmbedConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCopilotEmbedConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCopilotEmbedConfigResponse) GetBody() *ModifyCopilotEmbedConfigResponseBody {
	return s.Body
}

func (s *ModifyCopilotEmbedConfigResponse) SetHeaders(v map[string]*string) *ModifyCopilotEmbedConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyCopilotEmbedConfigResponse) SetStatusCode(v int32) *ModifyCopilotEmbedConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCopilotEmbedConfigResponse) SetBody(v *ModifyCopilotEmbedConfigResponseBody) *ModifyCopilotEmbedConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyCopilotEmbedConfigResponse) Validate() error {
	return dara.Validate(s)
}
