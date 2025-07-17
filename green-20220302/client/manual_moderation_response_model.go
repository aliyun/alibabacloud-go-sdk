// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManualModerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ManualModerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ManualModerationResponse
	GetStatusCode() *int32
	SetBody(v *ManualModerationResponseBody) *ManualModerationResponse
	GetBody() *ManualModerationResponseBody
}

type ManualModerationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ManualModerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManualModerationResponse) String() string {
	return dara.Prettify(s)
}

func (s ManualModerationResponse) GoString() string {
	return s.String()
}

func (s *ManualModerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ManualModerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ManualModerationResponse) GetBody() *ManualModerationResponseBody {
	return s.Body
}

func (s *ManualModerationResponse) SetHeaders(v map[string]*string) *ManualModerationResponse {
	s.Headers = v
	return s
}

func (s *ManualModerationResponse) SetStatusCode(v int32) *ManualModerationResponse {
	s.StatusCode = &v
	return s
}

func (s *ManualModerationResponse) SetBody(v *ManualModerationResponseBody) *ManualModerationResponse {
	s.Body = v
	return s
}

func (s *ManualModerationResponse) Validate() error {
	return dara.Validate(s)
}
