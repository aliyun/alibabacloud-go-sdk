// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFormRemarksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFormRemarksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFormRemarksResponse
	GetStatusCode() *int32
	SetBody(v *ListFormRemarksResponseBody) *ListFormRemarksResponse
	GetBody() *ListFormRemarksResponseBody
}

type ListFormRemarksResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFormRemarksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFormRemarksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFormRemarksResponse) GoString() string {
	return s.String()
}

func (s *ListFormRemarksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFormRemarksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFormRemarksResponse) GetBody() *ListFormRemarksResponseBody {
	return s.Body
}

func (s *ListFormRemarksResponse) SetHeaders(v map[string]*string) *ListFormRemarksResponse {
	s.Headers = v
	return s
}

func (s *ListFormRemarksResponse) SetStatusCode(v int32) *ListFormRemarksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFormRemarksResponse) SetBody(v *ListFormRemarksResponseBody) *ListFormRemarksResponse {
	s.Body = v
	return s
}

func (s *ListFormRemarksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
