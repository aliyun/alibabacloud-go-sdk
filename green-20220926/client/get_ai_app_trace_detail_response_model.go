// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiAppTraceDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAiAppTraceDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAiAppTraceDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetAiAppTraceDetailResponseBody) *GetAiAppTraceDetailResponse
	GetBody() *GetAiAppTraceDetailResponseBody
}

type GetAiAppTraceDetailResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAiAppTraceDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAiAppTraceDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppTraceDetailResponse) GoString() string {
	return s.String()
}

func (s *GetAiAppTraceDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAiAppTraceDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAiAppTraceDetailResponse) GetBody() *GetAiAppTraceDetailResponseBody {
	return s.Body
}

func (s *GetAiAppTraceDetailResponse) SetHeaders(v map[string]*string) *GetAiAppTraceDetailResponse {
	s.Headers = v
	return s
}

func (s *GetAiAppTraceDetailResponse) SetStatusCode(v int32) *GetAiAppTraceDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAiAppTraceDetailResponse) SetBody(v *GetAiAppTraceDetailResponseBody) *GetAiAppTraceDetailResponse {
	s.Body = v
	return s
}

func (s *GetAiAppTraceDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
