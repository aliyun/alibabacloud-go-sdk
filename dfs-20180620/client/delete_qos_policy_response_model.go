// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQosPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteQosPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteQosPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteQosPolicyResponseBody) *DeleteQosPolicyResponse
	GetBody() *DeleteQosPolicyResponseBody
}

type DeleteQosPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQosPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQosPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteQosPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteQosPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteQosPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteQosPolicyResponse) GetBody() *DeleteQosPolicyResponseBody {
	return s.Body
}

func (s *DeleteQosPolicyResponse) SetHeaders(v map[string]*string) *DeleteQosPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteQosPolicyResponse) SetStatusCode(v int32) *DeleteQosPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQosPolicyResponse) SetBody(v *DeleteQosPolicyResponseBody) *DeleteQosPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteQosPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
