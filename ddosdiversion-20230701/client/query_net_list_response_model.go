// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryNetListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryNetListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryNetListResponse
	GetStatusCode() *int32
	SetBody(v *QueryNetListResponseBody) *QueryNetListResponse
	GetBody() *QueryNetListResponseBody
}

type QueryNetListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryNetListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryNetListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryNetListResponse) GoString() string {
	return s.String()
}

func (s *QueryNetListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryNetListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryNetListResponse) GetBody() *QueryNetListResponseBody {
	return s.Body
}

func (s *QueryNetListResponse) SetHeaders(v map[string]*string) *QueryNetListResponse {
	s.Headers = v
	return s
}

func (s *QueryNetListResponse) SetStatusCode(v int32) *QueryNetListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryNetListResponse) SetBody(v *QueryNetListResponseBody) *QueryNetListResponse {
	s.Body = v
	return s
}

func (s *QueryNetListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
