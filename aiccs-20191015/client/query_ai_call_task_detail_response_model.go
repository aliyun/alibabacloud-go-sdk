// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAiCallTaskDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAiCallTaskDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAiCallTaskDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryAiCallTaskDetailResponseBody) *QueryAiCallTaskDetailResponse
	GetBody() *QueryAiCallTaskDetailResponseBody
}

type QueryAiCallTaskDetailResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAiCallTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAiCallTaskDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAiCallTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryAiCallTaskDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAiCallTaskDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAiCallTaskDetailResponse) GetBody() *QueryAiCallTaskDetailResponseBody {
	return s.Body
}

func (s *QueryAiCallTaskDetailResponse) SetHeaders(v map[string]*string) *QueryAiCallTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryAiCallTaskDetailResponse) SetStatusCode(v int32) *QueryAiCallTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAiCallTaskDetailResponse) SetBody(v *QueryAiCallTaskDetailResponseBody) *QueryAiCallTaskDetailResponse {
	s.Body = v
	return s
}

func (s *QueryAiCallTaskDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
