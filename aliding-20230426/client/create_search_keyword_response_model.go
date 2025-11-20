// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSearchKeywordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSearchKeywordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSearchKeywordResponse
	GetStatusCode() *int32
	SetBody(v *CreateSearchKeywordResponseBody) *CreateSearchKeywordResponse
	GetBody() *CreateSearchKeywordResponseBody
}

type CreateSearchKeywordResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSearchKeywordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSearchKeywordResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSearchKeywordResponse) GoString() string {
	return s.String()
}

func (s *CreateSearchKeywordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSearchKeywordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSearchKeywordResponse) GetBody() *CreateSearchKeywordResponseBody {
	return s.Body
}

func (s *CreateSearchKeywordResponse) SetHeaders(v map[string]*string) *CreateSearchKeywordResponse {
	s.Headers = v
	return s
}

func (s *CreateSearchKeywordResponse) SetStatusCode(v int32) *CreateSearchKeywordResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSearchKeywordResponse) SetBody(v *CreateSearchKeywordResponseBody) *CreateSearchKeywordResponse {
	s.Body = v
	return s
}

func (s *CreateSearchKeywordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
