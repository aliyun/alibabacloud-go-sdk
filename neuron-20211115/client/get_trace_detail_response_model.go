// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTraceDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTraceDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTraceDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetTraceDetailResponseBody) *GetTraceDetailResponse
	GetBody() *GetTraceDetailResponseBody
}

type GetTraceDetailResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTraceDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTraceDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTraceDetailResponse) GoString() string {
	return s.String()
}

func (s *GetTraceDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTraceDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTraceDetailResponse) GetBody() *GetTraceDetailResponseBody {
	return s.Body
}

func (s *GetTraceDetailResponse) SetHeaders(v map[string]*string) *GetTraceDetailResponse {
	s.Headers = v
	return s
}

func (s *GetTraceDetailResponse) SetStatusCode(v int32) *GetTraceDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTraceDetailResponse) SetBody(v *GetTraceDetailResponseBody) *GetTraceDetailResponse {
	s.Body = v
	return s
}

func (s *GetTraceDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
