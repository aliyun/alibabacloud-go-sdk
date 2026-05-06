// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelProviderTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteModelProviderTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteModelProviderTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteModelProviderTemplateResponseBody) *DeleteModelProviderTemplateResponse
	GetBody() *DeleteModelProviderTemplateResponseBody
}

type DeleteModelProviderTemplateResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteModelProviderTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteModelProviderTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelProviderTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteModelProviderTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteModelProviderTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteModelProviderTemplateResponse) GetBody() *DeleteModelProviderTemplateResponseBody {
	return s.Body
}

func (s *DeleteModelProviderTemplateResponse) SetHeaders(v map[string]*string) *DeleteModelProviderTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteModelProviderTemplateResponse) SetStatusCode(v int32) *DeleteModelProviderTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteModelProviderTemplateResponse) SetBody(v *DeleteModelProviderTemplateResponseBody) *DeleteModelProviderTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteModelProviderTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
