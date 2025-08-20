// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectStackDriftResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectStackDriftResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectStackDriftResponse
	GetStatusCode() *int32
	SetBody(v *DetectStackDriftResponseBody) *DetectStackDriftResponse
	GetBody() *DetectStackDriftResponseBody
}

type DetectStackDriftResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectStackDriftResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectStackDriftResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectStackDriftResponse) GoString() string {
	return s.String()
}

func (s *DetectStackDriftResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectStackDriftResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectStackDriftResponse) GetBody() *DetectStackDriftResponseBody {
	return s.Body
}

func (s *DetectStackDriftResponse) SetHeaders(v map[string]*string) *DetectStackDriftResponse {
	s.Headers = v
	return s
}

func (s *DetectStackDriftResponse) SetStatusCode(v int32) *DetectStackDriftResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectStackDriftResponse) SetBody(v *DetectStackDriftResponseBody) *DetectStackDriftResponse {
	s.Body = v
	return s
}

func (s *DetectStackDriftResponse) Validate() error {
	return dara.Validate(s)
}
