// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetrieveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetrieveResponse
	GetStatusCode() *int32
	SetBody(v *RetrieveResponseBody) *RetrieveResponse
	GetBody() *RetrieveResponseBody
}

type RetrieveResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetrieveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetrieveResponse) String() string {
	return dara.Prettify(s)
}

func (s RetrieveResponse) GoString() string {
	return s.String()
}

func (s *RetrieveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetrieveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetrieveResponse) GetBody() *RetrieveResponseBody {
	return s.Body
}

func (s *RetrieveResponse) SetHeaders(v map[string]*string) *RetrieveResponse {
	s.Headers = v
	return s
}

func (s *RetrieveResponse) SetStatusCode(v int32) *RetrieveResponse {
	s.StatusCode = &v
	return s
}

func (s *RetrieveResponse) SetBody(v *RetrieveResponseBody) *RetrieveResponse {
	s.Body = v
	return s
}

func (s *RetrieveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
