// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMemoryListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMemoryListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMemoryListResponse
	GetStatusCode() *int32
	SetBody(v *QueryMemoryListResponseBody) *QueryMemoryListResponse
	GetBody() *QueryMemoryListResponseBody
}

type QueryMemoryListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMemoryListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMemoryListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMemoryListResponse) GoString() string {
	return s.String()
}

func (s *QueryMemoryListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMemoryListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMemoryListResponse) GetBody() *QueryMemoryListResponseBody {
	return s.Body
}

func (s *QueryMemoryListResponse) SetHeaders(v map[string]*string) *QueryMemoryListResponse {
	s.Headers = v
	return s
}

func (s *QueryMemoryListResponse) SetStatusCode(v int32) *QueryMemoryListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMemoryListResponse) SetBody(v *QueryMemoryListResponseBody) *QueryMemoryListResponse {
	s.Body = v
	return s
}

func (s *QueryMemoryListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
