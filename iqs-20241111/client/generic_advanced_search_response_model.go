// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenericAdvancedSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenericAdvancedSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenericAdvancedSearchResponse
	GetStatusCode() *int32
	SetBody(v *GenericSearchResult) *GenericAdvancedSearchResponse
	GetBody() *GenericSearchResult
}

type GenericAdvancedSearchResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenericSearchResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenericAdvancedSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s GenericAdvancedSearchResponse) GoString() string {
	return s.String()
}

func (s *GenericAdvancedSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenericAdvancedSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenericAdvancedSearchResponse) GetBody() *GenericSearchResult {
	return s.Body
}

func (s *GenericAdvancedSearchResponse) SetHeaders(v map[string]*string) *GenericAdvancedSearchResponse {
	s.Headers = v
	return s
}

func (s *GenericAdvancedSearchResponse) SetStatusCode(v int32) *GenericAdvancedSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *GenericAdvancedSearchResponse) SetBody(v *GenericSearchResult) *GenericAdvancedSearchResponse {
	s.Body = v
	return s
}

func (s *GenericAdvancedSearchResponse) Validate() error {
	return dara.Validate(s)
}
