// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPtsReportsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPtsReportsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPtsReportsResponse
	GetStatusCode() *int32
	SetBody(v *ListPtsReportsResponseBody) *ListPtsReportsResponse
	GetBody() *ListPtsReportsResponseBody
}

type ListPtsReportsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPtsReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPtsReportsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPtsReportsResponse) GoString() string {
	return s.String()
}

func (s *ListPtsReportsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPtsReportsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPtsReportsResponse) GetBody() *ListPtsReportsResponseBody {
	return s.Body
}

func (s *ListPtsReportsResponse) SetHeaders(v map[string]*string) *ListPtsReportsResponse {
	s.Headers = v
	return s
}

func (s *ListPtsReportsResponse) SetStatusCode(v int32) *ListPtsReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPtsReportsResponse) SetBody(v *ListPtsReportsResponseBody) *ListPtsReportsResponse {
	s.Body = v
	return s
}

func (s *ListPtsReportsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
