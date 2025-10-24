// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTraceM3u8JobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTraceM3u8JobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTraceM3u8JobResponse
	GetStatusCode() *int32
	SetBody(v *QueryTraceM3u8JobResponseBody) *QueryTraceM3u8JobResponse
	GetBody() *QueryTraceM3u8JobResponseBody
}

type QueryTraceM3u8JobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTraceM3u8JobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTraceM3u8JobResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTraceM3u8JobResponse) GoString() string {
	return s.String()
}

func (s *QueryTraceM3u8JobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTraceM3u8JobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTraceM3u8JobResponse) GetBody() *QueryTraceM3u8JobResponseBody {
	return s.Body
}

func (s *QueryTraceM3u8JobResponse) SetHeaders(v map[string]*string) *QueryTraceM3u8JobResponse {
	s.Headers = v
	return s
}

func (s *QueryTraceM3u8JobResponse) SetStatusCode(v int32) *QueryTraceM3u8JobResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTraceM3u8JobResponse) SetBody(v *QueryTraceM3u8JobResponseBody) *QueryTraceM3u8JobResponse {
	s.Body = v
	return s
}

func (s *QueryTraceM3u8JobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
