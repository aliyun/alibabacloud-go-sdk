// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchIndexJobRerunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchIndexJobRerunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchIndexJobRerunResponse
	GetStatusCode() *int32
	SetBody(v *SearchIndexJobRerunResponseBody) *SearchIndexJobRerunResponse
	GetBody() *SearchIndexJobRerunResponseBody
}

type SearchIndexJobRerunResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchIndexJobRerunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchIndexJobRerunResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchIndexJobRerunResponse) GoString() string {
	return s.String()
}

func (s *SearchIndexJobRerunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchIndexJobRerunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchIndexJobRerunResponse) GetBody() *SearchIndexJobRerunResponseBody {
	return s.Body
}

func (s *SearchIndexJobRerunResponse) SetHeaders(v map[string]*string) *SearchIndexJobRerunResponse {
	s.Headers = v
	return s
}

func (s *SearchIndexJobRerunResponse) SetStatusCode(v int32) *SearchIndexJobRerunResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchIndexJobRerunResponse) SetBody(v *SearchIndexJobRerunResponseBody) *SearchIndexJobRerunResponse {
	s.Body = v
	return s
}

func (s *SearchIndexJobRerunResponse) Validate() error {
	return dara.Validate(s)
}
