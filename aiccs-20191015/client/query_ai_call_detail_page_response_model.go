// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAiCallDetailPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAiCallDetailPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAiCallDetailPageResponse
	GetStatusCode() *int32
	SetBody(v *QueryAiCallDetailPageResponseBody) *QueryAiCallDetailPageResponse
	GetBody() *QueryAiCallDetailPageResponseBody
}

type QueryAiCallDetailPageResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAiCallDetailPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAiCallDetailPageResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAiCallDetailPageResponse) GoString() string {
	return s.String()
}

func (s *QueryAiCallDetailPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAiCallDetailPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAiCallDetailPageResponse) GetBody() *QueryAiCallDetailPageResponseBody {
	return s.Body
}

func (s *QueryAiCallDetailPageResponse) SetHeaders(v map[string]*string) *QueryAiCallDetailPageResponse {
	s.Headers = v
	return s
}

func (s *QueryAiCallDetailPageResponse) SetStatusCode(v int32) *QueryAiCallDetailPageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAiCallDetailPageResponse) SetBody(v *QueryAiCallDetailPageResponseBody) *QueryAiCallDetailPageResponse {
	s.Body = v
	return s
}

func (s *QueryAiCallDetailPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
