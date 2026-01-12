// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyForTryOnResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyForTryOnResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyForTryOnResponse
	GetStatusCode() *int32
	SetBody(v *ApplyForTryOnResponseBody) *ApplyForTryOnResponse
	GetBody() *ApplyForTryOnResponseBody
}

type ApplyForTryOnResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyForTryOnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyForTryOnResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyForTryOnResponse) GoString() string {
	return s.String()
}

func (s *ApplyForTryOnResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyForTryOnResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyForTryOnResponse) GetBody() *ApplyForTryOnResponseBody {
	return s.Body
}

func (s *ApplyForTryOnResponse) SetHeaders(v map[string]*string) *ApplyForTryOnResponse {
	s.Headers = v
	return s
}

func (s *ApplyForTryOnResponse) SetStatusCode(v int32) *ApplyForTryOnResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyForTryOnResponse) SetBody(v *ApplyForTryOnResponseBody) *ApplyForTryOnResponse {
	s.Body = v
	return s
}

func (s *ApplyForTryOnResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
