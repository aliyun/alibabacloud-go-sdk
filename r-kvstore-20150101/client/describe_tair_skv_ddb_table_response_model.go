// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairSkvDdbTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTairSkvDdbTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTairSkvDdbTableResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTairSkvDdbTableResponseBody) *DescribeTairSkvDdbTableResponse
	GetBody() *DescribeTairSkvDdbTableResponseBody
}

type DescribeTairSkvDdbTableResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTairSkvDdbTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTairSkvDdbTableResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairSkvDdbTableResponse) GoString() string {
	return s.String()
}

func (s *DescribeTairSkvDdbTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTairSkvDdbTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTairSkvDdbTableResponse) GetBody() *DescribeTairSkvDdbTableResponseBody {
	return s.Body
}

func (s *DescribeTairSkvDdbTableResponse) SetHeaders(v map[string]*string) *DescribeTairSkvDdbTableResponse {
	s.Headers = v
	return s
}

func (s *DescribeTairSkvDdbTableResponse) SetStatusCode(v int32) *DescribeTairSkvDdbTableResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTairSkvDdbTableResponse) SetBody(v *DescribeTairSkvDdbTableResponseBody) *DescribeTairSkvDdbTableResponse {
	s.Body = v
	return s
}

func (s *DescribeTairSkvDdbTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
