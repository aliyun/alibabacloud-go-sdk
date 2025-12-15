// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWebSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWebSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWebSearchResponse
	GetStatusCode() *int32
	SetBody(v *GetWebSearchResponseBody) *GetWebSearchResponse
	GetBody() *GetWebSearchResponseBody
}

type GetWebSearchResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWebSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWebSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWebSearchResponse) GoString() string {
	return s.String()
}

func (s *GetWebSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWebSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWebSearchResponse) GetBody() *GetWebSearchResponseBody {
	return s.Body
}

func (s *GetWebSearchResponse) SetHeaders(v map[string]*string) *GetWebSearchResponse {
	s.Headers = v
	return s
}

func (s *GetWebSearchResponse) SetStatusCode(v int32) *GetWebSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWebSearchResponse) SetBody(v *GetWebSearchResponseBody) *GetWebSearchResponse {
	s.Body = v
	return s
}

func (s *GetWebSearchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
