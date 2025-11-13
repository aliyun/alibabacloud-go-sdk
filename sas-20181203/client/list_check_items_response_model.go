// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCheckItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCheckItemsResponse
	GetStatusCode() *int32
	SetBody(v *ListCheckItemsResponseBody) *ListCheckItemsResponse
	GetBody() *ListCheckItemsResponseBody
}

type ListCheckItemsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCheckItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCheckItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCheckItemsResponse) GoString() string {
	return s.String()
}

func (s *ListCheckItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCheckItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCheckItemsResponse) GetBody() *ListCheckItemsResponseBody {
	return s.Body
}

func (s *ListCheckItemsResponse) SetHeaders(v map[string]*string) *ListCheckItemsResponse {
	s.Headers = v
	return s
}

func (s *ListCheckItemsResponse) SetStatusCode(v int32) *ListCheckItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCheckItemsResponse) SetBody(v *ListCheckItemsResponseBody) *ListCheckItemsResponse {
	s.Body = v
	return s
}

func (s *ListCheckItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
