// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationEcuResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationEcuResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationEcuResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationEcuResponseBody) *ListApplicationEcuResponse
	GetBody() *ListApplicationEcuResponseBody
}

type ListApplicationEcuResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationEcuResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationEcuResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationEcuResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationEcuResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationEcuResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationEcuResponse) GetBody() *ListApplicationEcuResponseBody {
	return s.Body
}

func (s *ListApplicationEcuResponse) SetHeaders(v map[string]*string) *ListApplicationEcuResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationEcuResponse) SetStatusCode(v int32) *ListApplicationEcuResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationEcuResponse) SetBody(v *ListApplicationEcuResponseBody) *ListApplicationEcuResponse {
	s.Body = v
	return s
}

func (s *ListApplicationEcuResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
