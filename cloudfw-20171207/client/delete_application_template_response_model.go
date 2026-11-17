// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApplicationTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApplicationTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApplicationTemplateResponseBody) *DeleteApplicationTemplateResponse
	GetBody() *DeleteApplicationTemplateResponseBody
}

type DeleteApplicationTemplateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApplicationTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApplicationTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApplicationTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApplicationTemplateResponse) GetBody() *DeleteApplicationTemplateResponseBody {
	return s.Body
}

func (s *DeleteApplicationTemplateResponse) SetHeaders(v map[string]*string) *DeleteApplicationTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationTemplateResponse) SetStatusCode(v int32) *DeleteApplicationTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApplicationTemplateResponse) SetBody(v *DeleteApplicationTemplateResponseBody) *DeleteApplicationTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteApplicationTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
