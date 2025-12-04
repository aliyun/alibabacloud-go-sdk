// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAdvancedQueryTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAdvancedQueryTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAdvancedQueryTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAdvancedQueryTemplateResponseBody) *DeleteAdvancedQueryTemplateResponse
	GetBody() *DeleteAdvancedQueryTemplateResponseBody
}

type DeleteAdvancedQueryTemplateResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAdvancedQueryTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAdvancedQueryTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAdvancedQueryTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteAdvancedQueryTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAdvancedQueryTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAdvancedQueryTemplateResponse) GetBody() *DeleteAdvancedQueryTemplateResponseBody {
	return s.Body
}

func (s *DeleteAdvancedQueryTemplateResponse) SetHeaders(v map[string]*string) *DeleteAdvancedQueryTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteAdvancedQueryTemplateResponse) SetStatusCode(v int32) *DeleteAdvancedQueryTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAdvancedQueryTemplateResponse) SetBody(v *DeleteAdvancedQueryTemplateResponseBody) *DeleteAdvancedQueryTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteAdvancedQueryTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
