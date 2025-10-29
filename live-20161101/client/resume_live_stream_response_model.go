// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeLiveStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeLiveStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeLiveStreamResponse
	GetStatusCode() *int32
	SetBody(v *ResumeLiveStreamResponseBody) *ResumeLiveStreamResponse
	GetBody() *ResumeLiveStreamResponseBody
}

type ResumeLiveStreamResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeLiveStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeLiveStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeLiveStreamResponse) GoString() string {
	return s.String()
}

func (s *ResumeLiveStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeLiveStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeLiveStreamResponse) GetBody() *ResumeLiveStreamResponseBody {
	return s.Body
}

func (s *ResumeLiveStreamResponse) SetHeaders(v map[string]*string) *ResumeLiveStreamResponse {
	s.Headers = v
	return s
}

func (s *ResumeLiveStreamResponse) SetStatusCode(v int32) *ResumeLiveStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeLiveStreamResponse) SetBody(v *ResumeLiveStreamResponseBody) *ResumeLiveStreamResponse {
	s.Body = v
	return s
}

func (s *ResumeLiveStreamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
