// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateConsumerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateConsumerResponse
	GetStatusCode() *int32
	SetBody(v *CreateConsumerResponseBody) *CreateConsumerResponse
	GetBody() *CreateConsumerResponseBody
}

type CreateConsumerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConsumerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConsumerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerResponse) GoString() string {
	return s.String()
}

func (s *CreateConsumerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateConsumerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateConsumerResponse) GetBody() *CreateConsumerResponseBody {
	return s.Body
}

func (s *CreateConsumerResponse) SetHeaders(v map[string]*string) *CreateConsumerResponse {
	s.Headers = v
	return s
}

func (s *CreateConsumerResponse) SetStatusCode(v int32) *CreateConsumerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConsumerResponse) SetBody(v *CreateConsumerResponseBody) *CreateConsumerResponse {
	s.Body = v
	return s
}

func (s *CreateConsumerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
