// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMediaListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMediaListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMediaListResponse
	GetStatusCode() *int32
	SetBody(v *QueryMediaListResponseBody) *QueryMediaListResponse
	GetBody() *QueryMediaListResponseBody
}

type QueryMediaListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMediaListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMediaListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaListResponse) GoString() string {
	return s.String()
}

func (s *QueryMediaListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMediaListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMediaListResponse) GetBody() *QueryMediaListResponseBody {
	return s.Body
}

func (s *QueryMediaListResponse) SetHeaders(v map[string]*string) *QueryMediaListResponse {
	s.Headers = v
	return s
}

func (s *QueryMediaListResponse) SetStatusCode(v int32) *QueryMediaListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMediaListResponse) SetBody(v *QueryMediaListResponseBody) *QueryMediaListResponse {
	s.Body = v
	return s
}

func (s *QueryMediaListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
