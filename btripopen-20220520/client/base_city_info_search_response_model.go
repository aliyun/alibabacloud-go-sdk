// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBaseCityInfoSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BaseCityInfoSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BaseCityInfoSearchResponse
	GetStatusCode() *int32
	SetBody(v *BaseCityInfoSearchResponseBody) *BaseCityInfoSearchResponse
	GetBody() *BaseCityInfoSearchResponseBody
}

type BaseCityInfoSearchResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BaseCityInfoSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BaseCityInfoSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s BaseCityInfoSearchResponse) GoString() string {
	return s.String()
}

func (s *BaseCityInfoSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BaseCityInfoSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BaseCityInfoSearchResponse) GetBody() *BaseCityInfoSearchResponseBody {
	return s.Body
}

func (s *BaseCityInfoSearchResponse) SetHeaders(v map[string]*string) *BaseCityInfoSearchResponse {
	s.Headers = v
	return s
}

func (s *BaseCityInfoSearchResponse) SetStatusCode(v int32) *BaseCityInfoSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *BaseCityInfoSearchResponse) SetBody(v *BaseCityInfoSearchResponseBody) *BaseCityInfoSearchResponse {
	s.Body = v
	return s
}

func (s *BaseCityInfoSearchResponse) Validate() error {
	return dara.Validate(s)
}
