// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStopStreamsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchStopStreamsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchStopStreamsResponse
	GetStatusCode() *int32
	SetBody(v *BatchStopStreamsResponseBody) *BatchStopStreamsResponse
	GetBody() *BatchStopStreamsResponseBody
}

type BatchStopStreamsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchStopStreamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchStopStreamsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchStopStreamsResponse) GoString() string {
	return s.String()
}

func (s *BatchStopStreamsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchStopStreamsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchStopStreamsResponse) GetBody() *BatchStopStreamsResponseBody {
	return s.Body
}

func (s *BatchStopStreamsResponse) SetHeaders(v map[string]*string) *BatchStopStreamsResponse {
	s.Headers = v
	return s
}

func (s *BatchStopStreamsResponse) SetStatusCode(v int32) *BatchStopStreamsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchStopStreamsResponse) SetBody(v *BatchStopStreamsResponseBody) *BatchStopStreamsResponse {
	s.Body = v
	return s
}

func (s *BatchStopStreamsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
