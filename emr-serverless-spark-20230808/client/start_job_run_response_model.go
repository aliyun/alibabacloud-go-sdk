// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartJobRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartJobRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartJobRunResponse
	GetStatusCode() *int32
	SetBody(v *StartJobRunResponseBody) *StartJobRunResponse
	GetBody() *StartJobRunResponseBody
}

type StartJobRunResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartJobRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartJobRunResponse) String() string {
	return dara.Prettify(s)
}

func (s StartJobRunResponse) GoString() string {
	return s.String()
}

func (s *StartJobRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartJobRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartJobRunResponse) GetBody() *StartJobRunResponseBody {
	return s.Body
}

func (s *StartJobRunResponse) SetHeaders(v map[string]*string) *StartJobRunResponse {
	s.Headers = v
	return s
}

func (s *StartJobRunResponse) SetStatusCode(v int32) *StartJobRunResponse {
	s.StatusCode = &v
	return s
}

func (s *StartJobRunResponse) SetBody(v *StartJobRunResponseBody) *StartJobRunResponse {
	s.Body = v
	return s
}

func (s *StartJobRunResponse) Validate() error {
	return dara.Validate(s)
}
