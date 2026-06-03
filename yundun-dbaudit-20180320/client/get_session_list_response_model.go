// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSessionListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSessionListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSessionListResponse
	GetStatusCode() *int32
	SetBody(v *GetSessionListResponseBody) *GetSessionListResponse
	GetBody() *GetSessionListResponseBody
}

type GetSessionListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSessionListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSessionListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSessionListResponse) GoString() string {
	return s.String()
}

func (s *GetSessionListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSessionListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSessionListResponse) GetBody() *GetSessionListResponseBody {
	return s.Body
}

func (s *GetSessionListResponse) SetHeaders(v map[string]*string) *GetSessionListResponse {
	s.Headers = v
	return s
}

func (s *GetSessionListResponse) SetStatusCode(v int32) *GetSessionListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSessionListResponse) SetBody(v *GetSessionListResponseBody) *GetSessionListResponse {
	s.Body = v
	return s
}

func (s *GetSessionListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
