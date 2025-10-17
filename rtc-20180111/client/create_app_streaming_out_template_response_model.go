// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppStreamingOutTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppStreamingOutTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppStreamingOutTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppStreamingOutTemplateResponseBody) *CreateAppStreamingOutTemplateResponse
	GetBody() *CreateAppStreamingOutTemplateResponseBody
}

type CreateAppStreamingOutTemplateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppStreamingOutTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppStreamingOutTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppStreamingOutTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateAppStreamingOutTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppStreamingOutTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppStreamingOutTemplateResponse) GetBody() *CreateAppStreamingOutTemplateResponseBody {
	return s.Body
}

func (s *CreateAppStreamingOutTemplateResponse) SetHeaders(v map[string]*string) *CreateAppStreamingOutTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateAppStreamingOutTemplateResponse) SetStatusCode(v int32) *CreateAppStreamingOutTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppStreamingOutTemplateResponse) SetBody(v *CreateAppStreamingOutTemplateResponseBody) *CreateAppStreamingOutTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateAppStreamingOutTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
