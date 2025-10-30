// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedDataServiceApiDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuthorizedDataServiceApiDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuthorizedDataServiceApiDetailsResponse
	GetStatusCode() *int32
	SetBody(v *ListAuthorizedDataServiceApiDetailsResponseBody) *ListAuthorizedDataServiceApiDetailsResponse
	GetBody() *ListAuthorizedDataServiceApiDetailsResponseBody
}

type ListAuthorizedDataServiceApiDetailsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthorizedDataServiceApiDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthorizedDataServiceApiDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedDataServiceApiDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorizedDataServiceApiDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuthorizedDataServiceApiDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuthorizedDataServiceApiDetailsResponse) GetBody() *ListAuthorizedDataServiceApiDetailsResponseBody {
	return s.Body
}

func (s *ListAuthorizedDataServiceApiDetailsResponse) SetHeaders(v map[string]*string) *ListAuthorizedDataServiceApiDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponse) SetStatusCode(v int32) *ListAuthorizedDataServiceApiDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponse) SetBody(v *ListAuthorizedDataServiceApiDetailsResponseBody) *ListAuthorizedDataServiceApiDetailsResponse {
	s.Body = v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
