// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectStackGroupDriftResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectStackGroupDriftResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectStackGroupDriftResponse
	GetStatusCode() *int32
	SetBody(v *DetectStackGroupDriftResponseBody) *DetectStackGroupDriftResponse
	GetBody() *DetectStackGroupDriftResponseBody
}

type DetectStackGroupDriftResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectStackGroupDriftResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectStackGroupDriftResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectStackGroupDriftResponse) GoString() string {
	return s.String()
}

func (s *DetectStackGroupDriftResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectStackGroupDriftResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectStackGroupDriftResponse) GetBody() *DetectStackGroupDriftResponseBody {
	return s.Body
}

func (s *DetectStackGroupDriftResponse) SetHeaders(v map[string]*string) *DetectStackGroupDriftResponse {
	s.Headers = v
	return s
}

func (s *DetectStackGroupDriftResponse) SetStatusCode(v int32) *DetectStackGroupDriftResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectStackGroupDriftResponse) SetBody(v *DetectStackGroupDriftResponseBody) *DetectStackGroupDriftResponse {
	s.Body = v
	return s
}

func (s *DetectStackGroupDriftResponse) Validate() error {
	return dara.Validate(s)
}
