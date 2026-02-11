// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostEventWhiteruleListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PostEventWhiteruleListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PostEventWhiteruleListResponse
	GetStatusCode() *int32
	SetBody(v *PostEventWhiteruleListResponseBody) *PostEventWhiteruleListResponse
	GetBody() *PostEventWhiteruleListResponseBody
}

type PostEventWhiteruleListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostEventWhiteruleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostEventWhiteruleListResponse) String() string {
	return dara.Prettify(s)
}

func (s PostEventWhiteruleListResponse) GoString() string {
	return s.String()
}

func (s *PostEventWhiteruleListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PostEventWhiteruleListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PostEventWhiteruleListResponse) GetBody() *PostEventWhiteruleListResponseBody {
	return s.Body
}

func (s *PostEventWhiteruleListResponse) SetHeaders(v map[string]*string) *PostEventWhiteruleListResponse {
	s.Headers = v
	return s
}

func (s *PostEventWhiteruleListResponse) SetStatusCode(v int32) *PostEventWhiteruleListResponse {
	s.StatusCode = &v
	return s
}

func (s *PostEventWhiteruleListResponse) SetBody(v *PostEventWhiteruleListResponseBody) *PostEventWhiteruleListResponse {
	s.Body = v
	return s
}

func (s *PostEventWhiteruleListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
