// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateCacheResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTemplateCacheResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTemplateCacheResponse
	GetStatusCode() *int32
	SetBody(v *CreateTemplateCacheResponseBody) *CreateTemplateCacheResponse
	GetBody() *CreateTemplateCacheResponseBody
}

type CreateTemplateCacheResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTemplateCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTemplateCacheResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateCacheResponse) GoString() string {
	return s.String()
}

func (s *CreateTemplateCacheResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTemplateCacheResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTemplateCacheResponse) GetBody() *CreateTemplateCacheResponseBody {
	return s.Body
}

func (s *CreateTemplateCacheResponse) SetHeaders(v map[string]*string) *CreateTemplateCacheResponse {
	s.Headers = v
	return s
}

func (s *CreateTemplateCacheResponse) SetStatusCode(v int32) *CreateTemplateCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTemplateCacheResponse) SetBody(v *CreateTemplateCacheResponseBody) *CreateTemplateCacheResponse {
	s.Body = v
	return s
}

func (s *CreateTemplateCacheResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
