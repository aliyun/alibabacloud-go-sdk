// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GlobalSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GlobalSearchResponse
	GetStatusCode() *int32
	SetBody(v *GlobalSearchResult) *GlobalSearchResponse
	GetBody() *GlobalSearchResult
}

type GlobalSearchResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GlobalSearchResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GlobalSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s GlobalSearchResponse) GoString() string {
	return s.String()
}

func (s *GlobalSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GlobalSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GlobalSearchResponse) GetBody() *GlobalSearchResult {
	return s.Body
}

func (s *GlobalSearchResponse) SetHeaders(v map[string]*string) *GlobalSearchResponse {
	s.Headers = v
	return s
}

func (s *GlobalSearchResponse) SetStatusCode(v int32) *GlobalSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *GlobalSearchResponse) SetBody(v *GlobalSearchResult) *GlobalSearchResponse {
	s.Body = v
	return s
}

func (s *GlobalSearchResponse) Validate() error {
	return dara.Validate(s)
}
