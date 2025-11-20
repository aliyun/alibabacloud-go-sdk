// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTemplateResult) *DeleteTemplateResponse
	GetBody() *DeleteTemplateResult
}

type DeleteTemplateResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTemplateResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTemplateResponse) GetBody() *DeleteTemplateResult {
	return s.Body
}

func (s *DeleteTemplateResponse) SetHeaders(v map[string]*string) *DeleteTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteTemplateResponse) SetStatusCode(v int32) *DeleteTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTemplateResponse) SetBody(v *DeleteTemplateResult) *DeleteTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
