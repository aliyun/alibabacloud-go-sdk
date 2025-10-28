// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManualModerationResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ManualModerationResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ManualModerationResultResponse
	GetStatusCode() *int32
	SetBody(v *ManualModerationResultResponseBody) *ManualModerationResultResponse
	GetBody() *ManualModerationResultResponseBody
}

type ManualModerationResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ManualModerationResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManualModerationResultResponse) String() string {
	return dara.Prettify(s)
}

func (s ManualModerationResultResponse) GoString() string {
	return s.String()
}

func (s *ManualModerationResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ManualModerationResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ManualModerationResultResponse) GetBody() *ManualModerationResultResponseBody {
	return s.Body
}

func (s *ManualModerationResultResponse) SetHeaders(v map[string]*string) *ManualModerationResultResponse {
	s.Headers = v
	return s
}

func (s *ManualModerationResultResponse) SetStatusCode(v int32) *ManualModerationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ManualModerationResultResponse) SetBody(v *ManualModerationResultResponseBody) *ManualModerationResultResponse {
	s.Body = v
	return s
}

func (s *ManualModerationResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
