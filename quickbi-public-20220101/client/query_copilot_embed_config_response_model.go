// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCopilotEmbedConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCopilotEmbedConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCopilotEmbedConfigResponse
	GetStatusCode() *int32
	SetBody(v *QueryCopilotEmbedConfigResponseBody) *QueryCopilotEmbedConfigResponse
	GetBody() *QueryCopilotEmbedConfigResponseBody
}

type QueryCopilotEmbedConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCopilotEmbedConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCopilotEmbedConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCopilotEmbedConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryCopilotEmbedConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCopilotEmbedConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCopilotEmbedConfigResponse) GetBody() *QueryCopilotEmbedConfigResponseBody {
	return s.Body
}

func (s *QueryCopilotEmbedConfigResponse) SetHeaders(v map[string]*string) *QueryCopilotEmbedConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryCopilotEmbedConfigResponse) SetStatusCode(v int32) *QueryCopilotEmbedConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCopilotEmbedConfigResponse) SetBody(v *QueryCopilotEmbedConfigResponseBody) *QueryCopilotEmbedConfigResponse {
	s.Body = v
	return s
}

func (s *QueryCopilotEmbedConfigResponse) Validate() error {
	return dara.Validate(s)
}
