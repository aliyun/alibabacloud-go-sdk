// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAiCallTaskPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAiCallTaskPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAiCallTaskPageResponse
	GetStatusCode() *int32
	SetBody(v *QueryAiCallTaskPageResponseBody) *QueryAiCallTaskPageResponse
	GetBody() *QueryAiCallTaskPageResponseBody
}

type QueryAiCallTaskPageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAiCallTaskPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAiCallTaskPageResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAiCallTaskPageResponse) GoString() string {
	return s.String()
}

func (s *QueryAiCallTaskPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAiCallTaskPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAiCallTaskPageResponse) GetBody() *QueryAiCallTaskPageResponseBody {
	return s.Body
}

func (s *QueryAiCallTaskPageResponse) SetHeaders(v map[string]*string) *QueryAiCallTaskPageResponse {
	s.Headers = v
	return s
}

func (s *QueryAiCallTaskPageResponse) SetStatusCode(v int32) *QueryAiCallTaskPageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAiCallTaskPageResponse) SetBody(v *QueryAiCallTaskPageResponseBody) *QueryAiCallTaskPageResponse {
	s.Body = v
	return s
}

func (s *QueryAiCallTaskPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
