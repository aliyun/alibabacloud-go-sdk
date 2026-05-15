// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigNumListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConfigNumListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConfigNumListResponse
	GetStatusCode() *int32
	SetBody(v *GetConfigNumListResponseBody) *GetConfigNumListResponse
	GetBody() *GetConfigNumListResponseBody
}

type GetConfigNumListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConfigNumListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConfigNumListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConfigNumListResponse) GoString() string {
	return s.String()
}

func (s *GetConfigNumListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConfigNumListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConfigNumListResponse) GetBody() *GetConfigNumListResponseBody {
	return s.Body
}

func (s *GetConfigNumListResponse) SetHeaders(v map[string]*string) *GetConfigNumListResponse {
	s.Headers = v
	return s
}

func (s *GetConfigNumListResponse) SetStatusCode(v int32) *GetConfigNumListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConfigNumListResponse) SetBody(v *GetConfigNumListResponseBody) *GetConfigNumListResponse {
	s.Body = v
	return s
}

func (s *GetConfigNumListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
