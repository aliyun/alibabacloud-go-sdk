// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchForbidVsStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchForbidVsStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchForbidVsStreamResponse
	GetStatusCode() *int32
	SetBody(v *BatchForbidVsStreamResponseBody) *BatchForbidVsStreamResponse
	GetBody() *BatchForbidVsStreamResponseBody
}

type BatchForbidVsStreamResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchForbidVsStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchForbidVsStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchForbidVsStreamResponse) GoString() string {
	return s.String()
}

func (s *BatchForbidVsStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchForbidVsStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchForbidVsStreamResponse) GetBody() *BatchForbidVsStreamResponseBody {
	return s.Body
}

func (s *BatchForbidVsStreamResponse) SetHeaders(v map[string]*string) *BatchForbidVsStreamResponse {
	s.Headers = v
	return s
}

func (s *BatchForbidVsStreamResponse) SetStatusCode(v int32) *BatchForbidVsStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchForbidVsStreamResponse) SetBody(v *BatchForbidVsStreamResponseBody) *BatchForbidVsStreamResponse {
	s.Body = v
	return s
}

func (s *BatchForbidVsStreamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
