// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairSkvDdbTableSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTairSkvDdbTableSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTairSkvDdbTableSchemaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTairSkvDdbTableSchemaResponseBody) *DescribeTairSkvDdbTableSchemaResponse
	GetBody() *DescribeTairSkvDdbTableSchemaResponseBody
}

type DescribeTairSkvDdbTableSchemaResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTairSkvDdbTableSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTairSkvDdbTableSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairSkvDdbTableSchemaResponse) GoString() string {
	return s.String()
}

func (s *DescribeTairSkvDdbTableSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTairSkvDdbTableSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTairSkvDdbTableSchemaResponse) GetBody() *DescribeTairSkvDdbTableSchemaResponseBody {
	return s.Body
}

func (s *DescribeTairSkvDdbTableSchemaResponse) SetHeaders(v map[string]*string) *DescribeTairSkvDdbTableSchemaResponse {
	s.Headers = v
	return s
}

func (s *DescribeTairSkvDdbTableSchemaResponse) SetStatusCode(v int32) *DescribeTairSkvDdbTableSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTairSkvDdbTableSchemaResponse) SetBody(v *DescribeTairSkvDdbTableSchemaResponseBody) *DescribeTairSkvDdbTableSchemaResponse {
	s.Body = v
	return s
}

func (s *DescribeTairSkvDdbTableSchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
