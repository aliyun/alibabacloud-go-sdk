// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryJobListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryJobListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryJobListResponse
	GetStatusCode() *int32
	SetBody(v *QueryJobListResponseBody) *QueryJobListResponse
	GetBody() *QueryJobListResponseBody
}

type QueryJobListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryJobListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryJobListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryJobListResponse) GoString() string {
	return s.String()
}

func (s *QueryJobListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryJobListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryJobListResponse) GetBody() *QueryJobListResponseBody {
	return s.Body
}

func (s *QueryJobListResponse) SetHeaders(v map[string]*string) *QueryJobListResponse {
	s.Headers = v
	return s
}

func (s *QueryJobListResponse) SetStatusCode(v int32) *QueryJobListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryJobListResponse) SetBody(v *QueryJobListResponseBody) *QueryJobListResponse {
	s.Body = v
	return s
}

func (s *QueryJobListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
