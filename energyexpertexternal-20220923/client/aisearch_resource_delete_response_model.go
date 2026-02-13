// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAISearchResourceDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AISearchResourceDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AISearchResourceDeleteResponse
	GetStatusCode() *int32
	SetBody(v *AISearchResourceDeleteResponseBody) *AISearchResourceDeleteResponse
	GetBody() *AISearchResourceDeleteResponseBody
}

type AISearchResourceDeleteResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AISearchResourceDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AISearchResourceDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s AISearchResourceDeleteResponse) GoString() string {
	return s.String()
}

func (s *AISearchResourceDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AISearchResourceDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AISearchResourceDeleteResponse) GetBody() *AISearchResourceDeleteResponseBody {
	return s.Body
}

func (s *AISearchResourceDeleteResponse) SetHeaders(v map[string]*string) *AISearchResourceDeleteResponse {
	s.Headers = v
	return s
}

func (s *AISearchResourceDeleteResponse) SetStatusCode(v int32) *AISearchResourceDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *AISearchResourceDeleteResponse) SetBody(v *AISearchResourceDeleteResponseBody) *AISearchResourceDeleteResponse {
	s.Body = v
	return s
}

func (s *AISearchResourceDeleteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
