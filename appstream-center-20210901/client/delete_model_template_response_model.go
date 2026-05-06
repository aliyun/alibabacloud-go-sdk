// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteModelTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteModelTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteModelTemplateResponseBody) *DeleteModelTemplateResponse
	GetBody() *DeleteModelTemplateResponseBody
}

type DeleteModelTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteModelTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteModelTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteModelTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteModelTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteModelTemplateResponse) GetBody() *DeleteModelTemplateResponseBody {
	return s.Body
}

func (s *DeleteModelTemplateResponse) SetHeaders(v map[string]*string) *DeleteModelTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteModelTemplateResponse) SetStatusCode(v int32) *DeleteModelTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteModelTemplateResponse) SetBody(v *DeleteModelTemplateResponseBody) *DeleteModelTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteModelTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
