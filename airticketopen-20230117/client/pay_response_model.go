// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PayResponse
	GetStatusCode() *int32
	SetBody(v *PayResponseBody) *PayResponse
	GetBody() *PayResponseBody
}

type PayResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PayResponseBody   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PayResponse) String() string {
	return dara.Prettify(s)
}

func (s PayResponse) GoString() string {
	return s.String()
}

func (s *PayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PayResponse) GetBody() *PayResponseBody {
	return s.Body
}

func (s *PayResponse) SetHeaders(v map[string]*string) *PayResponse {
	s.Headers = v
	return s
}

func (s *PayResponse) SetStatusCode(v int32) *PayResponse {
	s.StatusCode = &v
	return s
}

func (s *PayResponse) SetBody(v *PayResponseBody) *PayResponse {
	s.Body = v
	return s
}

func (s *PayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
