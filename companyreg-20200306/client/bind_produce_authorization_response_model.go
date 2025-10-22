// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindProduceAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindProduceAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindProduceAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *BindProduceAuthorizationResponseBody) *BindProduceAuthorizationResponse
	GetBody() *BindProduceAuthorizationResponseBody
}

type BindProduceAuthorizationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindProduceAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindProduceAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s BindProduceAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *BindProduceAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindProduceAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindProduceAuthorizationResponse) GetBody() *BindProduceAuthorizationResponseBody {
	return s.Body
}

func (s *BindProduceAuthorizationResponse) SetHeaders(v map[string]*string) *BindProduceAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *BindProduceAuthorizationResponse) SetStatusCode(v int32) *BindProduceAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *BindProduceAuthorizationResponse) SetBody(v *BindProduceAuthorizationResponseBody) *BindProduceAuthorizationResponse {
	s.Body = v
	return s
}

func (s *BindProduceAuthorizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
