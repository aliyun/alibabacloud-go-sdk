// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProceedingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProceedingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProceedingsResponse
	GetStatusCode() *int32
	SetBody(v *ListProceedingsResponseBody) *ListProceedingsResponse
	GetBody() *ListProceedingsResponseBody
}

type ListProceedingsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProceedingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProceedingsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProceedingsResponse) GoString() string {
	return s.String()
}

func (s *ListProceedingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProceedingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProceedingsResponse) GetBody() *ListProceedingsResponseBody {
	return s.Body
}

func (s *ListProceedingsResponse) SetHeaders(v map[string]*string) *ListProceedingsResponse {
	s.Headers = v
	return s
}

func (s *ListProceedingsResponse) SetStatusCode(v int32) *ListProceedingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProceedingsResponse) SetBody(v *ListProceedingsResponseBody) *ListProceedingsResponse {
	s.Body = v
	return s
}

func (s *ListProceedingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
