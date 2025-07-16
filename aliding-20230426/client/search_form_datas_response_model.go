// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFormDatasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchFormDatasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchFormDatasResponse
	GetStatusCode() *int32
	SetBody(v *SearchFormDatasResponseBody) *SearchFormDatasResponse
	GetBody() *SearchFormDatasResponseBody
}

type SearchFormDatasResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchFormDatasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchFormDatasResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDatasResponse) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchFormDatasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchFormDatasResponse) GetBody() *SearchFormDatasResponseBody {
	return s.Body
}

func (s *SearchFormDatasResponse) SetHeaders(v map[string]*string) *SearchFormDatasResponse {
	s.Headers = v
	return s
}

func (s *SearchFormDatasResponse) SetStatusCode(v int32) *SearchFormDatasResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchFormDatasResponse) SetBody(v *SearchFormDatasResponseBody) *SearchFormDatasResponse {
	s.Body = v
	return s
}

func (s *SearchFormDatasResponse) Validate() error {
	return dara.Validate(s)
}
