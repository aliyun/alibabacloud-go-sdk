// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTraceM3u8JobListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTraceM3u8JobListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTraceM3u8JobListResponse
	GetStatusCode() *int32
	SetBody(v *QueryTraceM3u8JobListResponseBody) *QueryTraceM3u8JobListResponse
	GetBody() *QueryTraceM3u8JobListResponseBody
}

type QueryTraceM3u8JobListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTraceM3u8JobListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTraceM3u8JobListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTraceM3u8JobListResponse) GoString() string {
	return s.String()
}

func (s *QueryTraceM3u8JobListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTraceM3u8JobListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTraceM3u8JobListResponse) GetBody() *QueryTraceM3u8JobListResponseBody {
	return s.Body
}

func (s *QueryTraceM3u8JobListResponse) SetHeaders(v map[string]*string) *QueryTraceM3u8JobListResponse {
	s.Headers = v
	return s
}

func (s *QueryTraceM3u8JobListResponse) SetStatusCode(v int32) *QueryTraceM3u8JobListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTraceM3u8JobListResponse) SetBody(v *QueryTraceM3u8JobListResponseBody) *QueryTraceM3u8JobListResponse {
	s.Body = v
	return s
}

func (s *QueryTraceM3u8JobListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
