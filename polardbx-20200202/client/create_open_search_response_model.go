// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpenSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOpenSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOpenSearchResponse
	GetStatusCode() *int32
	SetBody(v *CreateOpenSearchResponseBody) *CreateOpenSearchResponse
	GetBody() *CreateOpenSearchResponseBody
}

type CreateOpenSearchResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOpenSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOpenSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenSearchResponse) GoString() string {
	return s.String()
}

func (s *CreateOpenSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOpenSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOpenSearchResponse) GetBody() *CreateOpenSearchResponseBody {
	return s.Body
}

func (s *CreateOpenSearchResponse) SetHeaders(v map[string]*string) *CreateOpenSearchResponse {
	s.Headers = v
	return s
}

func (s *CreateOpenSearchResponse) SetStatusCode(v int32) *CreateOpenSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOpenSearchResponse) SetBody(v *CreateOpenSearchResponseBody) *CreateOpenSearchResponse {
	s.Body = v
	return s
}

func (s *CreateOpenSearchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
