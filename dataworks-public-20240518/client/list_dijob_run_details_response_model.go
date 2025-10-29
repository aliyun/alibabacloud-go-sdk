// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDIJobRunDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDIJobRunDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDIJobRunDetailsResponse
	GetStatusCode() *int32
	SetBody(v *ListDIJobRunDetailsResponseBody) *ListDIJobRunDetailsResponse
	GetBody() *ListDIJobRunDetailsResponseBody
}

type ListDIJobRunDetailsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDIJobRunDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDIJobRunDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobRunDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListDIJobRunDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDIJobRunDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDIJobRunDetailsResponse) GetBody() *ListDIJobRunDetailsResponseBody {
	return s.Body
}

func (s *ListDIJobRunDetailsResponse) SetHeaders(v map[string]*string) *ListDIJobRunDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListDIJobRunDetailsResponse) SetStatusCode(v int32) *ListDIJobRunDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDIJobRunDetailsResponse) SetBody(v *ListDIJobRunDetailsResponseBody) *ListDIJobRunDetailsResponse {
	s.Body = v
	return s
}

func (s *ListDIJobRunDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
