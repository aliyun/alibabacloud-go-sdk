// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumerStackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConsumerStackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConsumerStackResponse
	GetStatusCode() *int32
	SetBody(v *GetConsumerStackResponseBody) *GetConsumerStackResponse
	GetBody() *GetConsumerStackResponseBody
}

type GetConsumerStackResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConsumerStackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConsumerStackResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerStackResponse) GoString() string {
	return s.String()
}

func (s *GetConsumerStackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConsumerStackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConsumerStackResponse) GetBody() *GetConsumerStackResponseBody {
	return s.Body
}

func (s *GetConsumerStackResponse) SetHeaders(v map[string]*string) *GetConsumerStackResponse {
	s.Headers = v
	return s
}

func (s *GetConsumerStackResponse) SetStatusCode(v int32) *GetConsumerStackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConsumerStackResponse) SetBody(v *GetConsumerStackResponseBody) *GetConsumerStackResponse {
	s.Body = v
	return s
}

func (s *GetConsumerStackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
