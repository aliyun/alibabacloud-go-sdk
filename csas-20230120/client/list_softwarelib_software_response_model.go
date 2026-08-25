// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSoftwarelibSoftwareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSoftwarelibSoftwareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSoftwarelibSoftwareResponse
	GetStatusCode() *int32
	SetBody(v *ListSoftwarelibSoftwareResponseBody) *ListSoftwarelibSoftwareResponse
	GetBody() *ListSoftwarelibSoftwareResponseBody
}

type ListSoftwarelibSoftwareResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSoftwarelibSoftwareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSoftwarelibSoftwareResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSoftwarelibSoftwareResponse) GoString() string {
	return s.String()
}

func (s *ListSoftwarelibSoftwareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSoftwarelibSoftwareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSoftwarelibSoftwareResponse) GetBody() *ListSoftwarelibSoftwareResponseBody {
	return s.Body
}

func (s *ListSoftwarelibSoftwareResponse) SetHeaders(v map[string]*string) *ListSoftwarelibSoftwareResponse {
	s.Headers = v
	return s
}

func (s *ListSoftwarelibSoftwareResponse) SetStatusCode(v int32) *ListSoftwarelibSoftwareResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponse) SetBody(v *ListSoftwarelibSoftwareResponseBody) *ListSoftwarelibSoftwareResponse {
	s.Body = v
	return s
}

func (s *ListSoftwarelibSoftwareResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
