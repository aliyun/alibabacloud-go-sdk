// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAISearchResourceUpdateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AISearchResourceUpdateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AISearchResourceUpdateResponse
	GetStatusCode() *int32
	SetBody(v *AISearchResourceUpdateResponseBody) *AISearchResourceUpdateResponse
	GetBody() *AISearchResourceUpdateResponseBody
}

type AISearchResourceUpdateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AISearchResourceUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AISearchResourceUpdateResponse) String() string {
	return dara.Prettify(s)
}

func (s AISearchResourceUpdateResponse) GoString() string {
	return s.String()
}

func (s *AISearchResourceUpdateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AISearchResourceUpdateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AISearchResourceUpdateResponse) GetBody() *AISearchResourceUpdateResponseBody {
	return s.Body
}

func (s *AISearchResourceUpdateResponse) SetHeaders(v map[string]*string) *AISearchResourceUpdateResponse {
	s.Headers = v
	return s
}

func (s *AISearchResourceUpdateResponse) SetStatusCode(v int32) *AISearchResourceUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *AISearchResourceUpdateResponse) SetBody(v *AISearchResourceUpdateResponseBody) *AISearchResourceUpdateResponse {
	s.Body = v
	return s
}

func (s *AISearchResourceUpdateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
