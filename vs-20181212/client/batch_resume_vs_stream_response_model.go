// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchResumeVsStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchResumeVsStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchResumeVsStreamResponse
	GetStatusCode() *int32
	SetBody(v *BatchResumeVsStreamResponseBody) *BatchResumeVsStreamResponse
	GetBody() *BatchResumeVsStreamResponseBody
}

type BatchResumeVsStreamResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchResumeVsStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchResumeVsStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchResumeVsStreamResponse) GoString() string {
	return s.String()
}

func (s *BatchResumeVsStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchResumeVsStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchResumeVsStreamResponse) GetBody() *BatchResumeVsStreamResponseBody {
	return s.Body
}

func (s *BatchResumeVsStreamResponse) SetHeaders(v map[string]*string) *BatchResumeVsStreamResponse {
	s.Headers = v
	return s
}

func (s *BatchResumeVsStreamResponse) SetStatusCode(v int32) *BatchResumeVsStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchResumeVsStreamResponse) SetBody(v *BatchResumeVsStreamResponseBody) *BatchResumeVsStreamResponse {
	s.Body = v
	return s
}

func (s *BatchResumeVsStreamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
