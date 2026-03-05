// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskRuleLimitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTaskRuleLimitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTaskRuleLimitResponse
	GetStatusCode() *int32
	SetBody(v *QueryTaskRuleLimitResponseBody) *QueryTaskRuleLimitResponse
	GetBody() *QueryTaskRuleLimitResponseBody
}

type QueryTaskRuleLimitResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTaskRuleLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTaskRuleLimitResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskRuleLimitResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskRuleLimitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTaskRuleLimitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTaskRuleLimitResponse) GetBody() *QueryTaskRuleLimitResponseBody {
	return s.Body
}

func (s *QueryTaskRuleLimitResponse) SetHeaders(v map[string]*string) *QueryTaskRuleLimitResponse {
	s.Headers = v
	return s
}

func (s *QueryTaskRuleLimitResponse) SetStatusCode(v int32) *QueryTaskRuleLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTaskRuleLimitResponse) SetBody(v *QueryTaskRuleLimitResponseBody) *QueryTaskRuleLimitResponse {
	s.Body = v
	return s
}

func (s *QueryTaskRuleLimitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
