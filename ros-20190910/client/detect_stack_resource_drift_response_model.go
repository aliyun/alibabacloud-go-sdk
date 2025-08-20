// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectStackResourceDriftResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectStackResourceDriftResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectStackResourceDriftResponse
	GetStatusCode() *int32
	SetBody(v *DetectStackResourceDriftResponseBody) *DetectStackResourceDriftResponse
	GetBody() *DetectStackResourceDriftResponseBody
}

type DetectStackResourceDriftResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectStackResourceDriftResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectStackResourceDriftResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectStackResourceDriftResponse) GoString() string {
	return s.String()
}

func (s *DetectStackResourceDriftResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectStackResourceDriftResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectStackResourceDriftResponse) GetBody() *DetectStackResourceDriftResponseBody {
	return s.Body
}

func (s *DetectStackResourceDriftResponse) SetHeaders(v map[string]*string) *DetectStackResourceDriftResponse {
	s.Headers = v
	return s
}

func (s *DetectStackResourceDriftResponse) SetStatusCode(v int32) *DetectStackResourceDriftResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectStackResourceDriftResponse) SetBody(v *DetectStackResourceDriftResponseBody) *DetectStackResourceDriftResponse {
	s.Body = v
	return s
}

func (s *DetectStackResourceDriftResponse) Validate() error {
	return dara.Validate(s)
}
