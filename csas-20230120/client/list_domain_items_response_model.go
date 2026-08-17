// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDomainItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDomainItemsResponse
	GetStatusCode() *int32
	SetBody(v *ListDomainItemsResponseBody) *ListDomainItemsResponse
	GetBody() *ListDomainItemsResponseBody
}

type ListDomainItemsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDomainItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDomainItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDomainItemsResponse) GoString() string {
	return s.String()
}

func (s *ListDomainItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDomainItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDomainItemsResponse) GetBody() *ListDomainItemsResponseBody {
	return s.Body
}

func (s *ListDomainItemsResponse) SetHeaders(v map[string]*string) *ListDomainItemsResponse {
	s.Headers = v
	return s
}

func (s *ListDomainItemsResponse) SetStatusCode(v int32) *ListDomainItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDomainItemsResponse) SetBody(v *ListDomainItemsResponseBody) *ListDomainItemsResponse {
	s.Body = v
	return s
}

func (s *ListDomainItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
