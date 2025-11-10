// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageContactsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMessageContactsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMessageContactsResponse
	GetStatusCode() *int32
	SetBody(v *ListMessageContactsResponseBody) *ListMessageContactsResponse
	GetBody() *ListMessageContactsResponseBody
}

type ListMessageContactsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMessageContactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMessageContactsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMessageContactsResponse) GoString() string {
	return s.String()
}

func (s *ListMessageContactsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMessageContactsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMessageContactsResponse) GetBody() *ListMessageContactsResponseBody {
	return s.Body
}

func (s *ListMessageContactsResponse) SetHeaders(v map[string]*string) *ListMessageContactsResponse {
	s.Headers = v
	return s
}

func (s *ListMessageContactsResponse) SetStatusCode(v int32) *ListMessageContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMessageContactsResponse) SetBody(v *ListMessageContactsResponseBody) *ListMessageContactsResponse {
	s.Body = v
	return s
}

func (s *ListMessageContactsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
