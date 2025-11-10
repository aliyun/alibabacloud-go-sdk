// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetConsumeOffsetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetConsumeOffsetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetConsumeOffsetResponse
	GetStatusCode() *int32
	SetBody(v *ResetConsumeOffsetResponseBody) *ResetConsumeOffsetResponse
	GetBody() *ResetConsumeOffsetResponseBody
}

type ResetConsumeOffsetResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetConsumeOffsetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetConsumeOffsetResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetConsumeOffsetResponse) GoString() string {
	return s.String()
}

func (s *ResetConsumeOffsetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetConsumeOffsetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetConsumeOffsetResponse) GetBody() *ResetConsumeOffsetResponseBody {
	return s.Body
}

func (s *ResetConsumeOffsetResponse) SetHeaders(v map[string]*string) *ResetConsumeOffsetResponse {
	s.Headers = v
	return s
}

func (s *ResetConsumeOffsetResponse) SetStatusCode(v int32) *ResetConsumeOffsetResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetConsumeOffsetResponse) SetBody(v *ResetConsumeOffsetResponseBody) *ResetConsumeOffsetResponse {
	s.Body = v
	return s
}

func (s *ResetConsumeOffsetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
