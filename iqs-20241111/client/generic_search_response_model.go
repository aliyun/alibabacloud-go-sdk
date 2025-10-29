// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenericSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenericSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenericSearchResponse
	GetStatusCode() *int32
	SetBody(v *GenericSearchResult) *GenericSearchResponse
	GetBody() *GenericSearchResult
}

type GenericSearchResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenericSearchResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenericSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s GenericSearchResponse) GoString() string {
	return s.String()
}

func (s *GenericSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenericSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenericSearchResponse) GetBody() *GenericSearchResult {
	return s.Body
}

func (s *GenericSearchResponse) SetHeaders(v map[string]*string) *GenericSearchResponse {
	s.Headers = v
	return s
}

func (s *GenericSearchResponse) SetStatusCode(v int32) *GenericSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *GenericSearchResponse) SetBody(v *GenericSearchResult) *GenericSearchResponse {
	s.Body = v
	return s
}

func (s *GenericSearchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
