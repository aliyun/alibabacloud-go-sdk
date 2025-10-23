// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIsCompletedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IsCompletedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IsCompletedResponse
	GetStatusCode() *int32
	SetBody(v *IsCompletedResponseBody) *IsCompletedResponse
	GetBody() *IsCompletedResponseBody
}

type IsCompletedResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IsCompletedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IsCompletedResponse) String() string {
	return dara.Prettify(s)
}

func (s IsCompletedResponse) GoString() string {
	return s.String()
}

func (s *IsCompletedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IsCompletedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IsCompletedResponse) GetBody() *IsCompletedResponseBody {
	return s.Body
}

func (s *IsCompletedResponse) SetHeaders(v map[string]*string) *IsCompletedResponse {
	s.Headers = v
	return s
}

func (s *IsCompletedResponse) SetStatusCode(v int32) *IsCompletedResponse {
	s.StatusCode = &v
	return s
}

func (s *IsCompletedResponse) SetBody(v *IsCompletedResponseBody) *IsCompletedResponse {
	s.Body = v
	return s
}

func (s *IsCompletedResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
