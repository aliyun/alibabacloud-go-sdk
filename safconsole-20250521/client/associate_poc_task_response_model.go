// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociatePocTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociatePocTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociatePocTaskResponse
	GetStatusCode() *int32
	SetBody(v *AssociatePocTaskResponseBody) *AssociatePocTaskResponse
	GetBody() *AssociatePocTaskResponseBody
}

type AssociatePocTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociatePocTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociatePocTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociatePocTaskResponse) GoString() string {
	return s.String()
}

func (s *AssociatePocTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociatePocTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociatePocTaskResponse) GetBody() *AssociatePocTaskResponseBody {
	return s.Body
}

func (s *AssociatePocTaskResponse) SetHeaders(v map[string]*string) *AssociatePocTaskResponse {
	s.Headers = v
	return s
}

func (s *AssociatePocTaskResponse) SetStatusCode(v int32) *AssociatePocTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociatePocTaskResponse) SetBody(v *AssociatePocTaskResponseBody) *AssociatePocTaskResponse {
	s.Body = v
	return s
}

func (s *AssociatePocTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
