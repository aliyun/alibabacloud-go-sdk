// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFormDataIdListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchFormDataIdListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchFormDataIdListResponse
	GetStatusCode() *int32
	SetBody(v *SearchFormDataIdListResponseBody) *SearchFormDataIdListResponse
	GetBody() *SearchFormDataIdListResponseBody
}

type SearchFormDataIdListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchFormDataIdListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchFormDataIdListResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataIdListResponse) GoString() string {
	return s.String()
}

func (s *SearchFormDataIdListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchFormDataIdListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchFormDataIdListResponse) GetBody() *SearchFormDataIdListResponseBody {
	return s.Body
}

func (s *SearchFormDataIdListResponse) SetHeaders(v map[string]*string) *SearchFormDataIdListResponse {
	s.Headers = v
	return s
}

func (s *SearchFormDataIdListResponse) SetStatusCode(v int32) *SearchFormDataIdListResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchFormDataIdListResponse) SetBody(v *SearchFormDataIdListResponseBody) *SearchFormDataIdListResponse {
	s.Body = v
	return s
}

func (s *SearchFormDataIdListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
