// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContactWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetContactWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetContactWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *GetContactWhiteListResponseBody) *GetContactWhiteListResponse
	GetBody() *GetContactWhiteListResponseBody
}

type GetContactWhiteListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContactWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContactWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetContactWhiteListResponse) GoString() string {
	return s.String()
}

func (s *GetContactWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetContactWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetContactWhiteListResponse) GetBody() *GetContactWhiteListResponseBody {
	return s.Body
}

func (s *GetContactWhiteListResponse) SetHeaders(v map[string]*string) *GetContactWhiteListResponse {
	s.Headers = v
	return s
}

func (s *GetContactWhiteListResponse) SetStatusCode(v int32) *GetContactWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContactWhiteListResponse) SetBody(v *GetContactWhiteListResponseBody) *GetContactWhiteListResponse {
	s.Body = v
	return s
}

func (s *GetContactWhiteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
