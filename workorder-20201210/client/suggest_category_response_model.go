// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuggestCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SuggestCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SuggestCategoryResponse
	GetStatusCode() *int32
	SetBody(v *SuggestCategoryResponseBody) *SuggestCategoryResponse
	GetBody() *SuggestCategoryResponseBody
}

type SuggestCategoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuggestCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuggestCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s SuggestCategoryResponse) GoString() string {
	return s.String()
}

func (s *SuggestCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SuggestCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SuggestCategoryResponse) GetBody() *SuggestCategoryResponseBody {
	return s.Body
}

func (s *SuggestCategoryResponse) SetHeaders(v map[string]*string) *SuggestCategoryResponse {
	s.Headers = v
	return s
}

func (s *SuggestCategoryResponse) SetStatusCode(v int32) *SuggestCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *SuggestCategoryResponse) SetBody(v *SuggestCategoryResponseBody) *SuggestCategoryResponse {
	s.Body = v
	return s
}

func (s *SuggestCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
