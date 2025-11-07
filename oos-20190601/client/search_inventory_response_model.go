// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchInventoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchInventoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchInventoryResponse
	GetStatusCode() *int32
	SetBody(v *SearchInventoryResponseBody) *SearchInventoryResponse
	GetBody() *SearchInventoryResponseBody
}

type SearchInventoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchInventoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchInventoryResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchInventoryResponse) GoString() string {
	return s.String()
}

func (s *SearchInventoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchInventoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchInventoryResponse) GetBody() *SearchInventoryResponseBody {
	return s.Body
}

func (s *SearchInventoryResponse) SetHeaders(v map[string]*string) *SearchInventoryResponse {
	s.Headers = v
	return s
}

func (s *SearchInventoryResponse) SetStatusCode(v int32) *SearchInventoryResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchInventoryResponse) SetBody(v *SearchInventoryResponseBody) *SearchInventoryResponse {
	s.Body = v
	return s
}

func (s *SearchInventoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
