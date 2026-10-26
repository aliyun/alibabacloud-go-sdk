// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAiproEmbedConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAiproEmbedConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAiproEmbedConfigResponse
	GetStatusCode() *int32
	SetBody(v *QueryAiproEmbedConfigResponseBody) *QueryAiproEmbedConfigResponse
	GetBody() *QueryAiproEmbedConfigResponseBody
}

type QueryAiproEmbedConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAiproEmbedConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAiproEmbedConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAiproEmbedConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryAiproEmbedConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAiproEmbedConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAiproEmbedConfigResponse) GetBody() *QueryAiproEmbedConfigResponseBody {
	return s.Body
}

func (s *QueryAiproEmbedConfigResponse) SetHeaders(v map[string]*string) *QueryAiproEmbedConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryAiproEmbedConfigResponse) SetStatusCode(v int32) *QueryAiproEmbedConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAiproEmbedConfigResponse) SetBody(v *QueryAiproEmbedConfigResponseBody) *QueryAiproEmbedConfigResponse {
	s.Body = v
	return s
}

func (s *QueryAiproEmbedConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
