// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSparkReplSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartSparkReplSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartSparkReplSessionResponse
	GetStatusCode() *int32
	SetBody(v *StartSparkReplSessionResponseBody) *StartSparkReplSessionResponse
	GetBody() *StartSparkReplSessionResponseBody
}

type StartSparkReplSessionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartSparkReplSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartSparkReplSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s StartSparkReplSessionResponse) GoString() string {
	return s.String()
}

func (s *StartSparkReplSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartSparkReplSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartSparkReplSessionResponse) GetBody() *StartSparkReplSessionResponseBody {
	return s.Body
}

func (s *StartSparkReplSessionResponse) SetHeaders(v map[string]*string) *StartSparkReplSessionResponse {
	s.Headers = v
	return s
}

func (s *StartSparkReplSessionResponse) SetStatusCode(v int32) *StartSparkReplSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *StartSparkReplSessionResponse) SetBody(v *StartSparkReplSessionResponseBody) *StartSparkReplSessionResponse {
	s.Body = v
	return s
}

func (s *StartSparkReplSessionResponse) Validate() error {
	return dara.Validate(s)
}
