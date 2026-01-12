// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClothesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClothesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClothesResponse
	GetStatusCode() *int32
	SetBody(v *ListClothesResponseBody) *ListClothesResponse
	GetBody() *ListClothesResponseBody
}

type ListClothesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClothesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClothesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClothesResponse) GoString() string {
	return s.String()
}

func (s *ListClothesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClothesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClothesResponse) GetBody() *ListClothesResponseBody {
	return s.Body
}

func (s *ListClothesResponse) SetHeaders(v map[string]*string) *ListClothesResponse {
	s.Headers = v
	return s
}

func (s *ListClothesResponse) SetStatusCode(v int32) *ListClothesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClothesResponse) SetBody(v *ListClothesResponseBody) *ListClothesResponse {
	s.Body = v
	return s
}

func (s *ListClothesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
