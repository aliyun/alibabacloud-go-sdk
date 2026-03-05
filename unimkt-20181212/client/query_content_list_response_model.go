// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryContentListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryContentListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryContentListResponse
	GetStatusCode() *int32
	SetBody(v *QueryContentListResponseBody) *QueryContentListResponse
	GetBody() *QueryContentListResponseBody
}

type QueryContentListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryContentListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryContentListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryContentListResponse) GoString() string {
	return s.String()
}

func (s *QueryContentListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryContentListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryContentListResponse) GetBody() *QueryContentListResponseBody {
	return s.Body
}

func (s *QueryContentListResponse) SetHeaders(v map[string]*string) *QueryContentListResponse {
	s.Headers = v
	return s
}

func (s *QueryContentListResponse) SetStatusCode(v int32) *QueryContentListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryContentListResponse) SetBody(v *QueryContentListResponseBody) *QueryContentListResponse {
	s.Body = v
	return s
}

func (s *QueryContentListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
