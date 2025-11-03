// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrefixListsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrefixListsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrefixListsResponse
	GetStatusCode() *int32
	SetBody(v *ListPrefixListsResponseBody) *ListPrefixListsResponse
	GetBody() *ListPrefixListsResponseBody
}

type ListPrefixListsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrefixListsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrefixListsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrefixListsResponse) GoString() string {
	return s.String()
}

func (s *ListPrefixListsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrefixListsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrefixListsResponse) GetBody() *ListPrefixListsResponseBody {
	return s.Body
}

func (s *ListPrefixListsResponse) SetHeaders(v map[string]*string) *ListPrefixListsResponse {
	s.Headers = v
	return s
}

func (s *ListPrefixListsResponse) SetStatusCode(v int32) *ListPrefixListsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrefixListsResponse) SetBody(v *ListPrefixListsResponseBody) *ListPrefixListsResponse {
	s.Body = v
	return s
}

func (s *ListPrefixListsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
