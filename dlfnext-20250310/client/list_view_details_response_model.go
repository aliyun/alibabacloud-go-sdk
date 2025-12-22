// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListViewDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListViewDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListViewDetailsResponse
	GetStatusCode() *int32
	SetBody(v *ListViewDetailsResponseBody) *ListViewDetailsResponse
	GetBody() *ListViewDetailsResponseBody
}

type ListViewDetailsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListViewDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListViewDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListViewDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListViewDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListViewDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListViewDetailsResponse) GetBody() *ListViewDetailsResponseBody {
	return s.Body
}

func (s *ListViewDetailsResponse) SetHeaders(v map[string]*string) *ListViewDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListViewDetailsResponse) SetStatusCode(v int32) *ListViewDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListViewDetailsResponse) SetBody(v *ListViewDetailsResponseBody) *ListViewDetailsResponse {
	s.Body = v
	return s
}

func (s *ListViewDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
