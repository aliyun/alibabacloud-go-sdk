// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAiproEmbedConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAiproEmbedConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAiproEmbedConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAiproEmbedConfigResponseBody) *ModifyAiproEmbedConfigResponse
	GetBody() *ModifyAiproEmbedConfigResponseBody
}

type ModifyAiproEmbedConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAiproEmbedConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAiproEmbedConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAiproEmbedConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyAiproEmbedConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAiproEmbedConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAiproEmbedConfigResponse) GetBody() *ModifyAiproEmbedConfigResponseBody {
	return s.Body
}

func (s *ModifyAiproEmbedConfigResponse) SetHeaders(v map[string]*string) *ModifyAiproEmbedConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyAiproEmbedConfigResponse) SetStatusCode(v int32) *ModifyAiproEmbedConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAiproEmbedConfigResponse) SetBody(v *ModifyAiproEmbedConfigResponseBody) *ModifyAiproEmbedConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyAiproEmbedConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
