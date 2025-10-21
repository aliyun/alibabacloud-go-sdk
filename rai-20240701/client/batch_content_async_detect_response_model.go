// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchContentAsyncDetectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchContentAsyncDetectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchContentAsyncDetectResponse
	GetStatusCode() *int32
	SetBody(v *BatchContentAsyncDetectResponseBody) *BatchContentAsyncDetectResponse
	GetBody() *BatchContentAsyncDetectResponseBody
}

type BatchContentAsyncDetectResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchContentAsyncDetectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchContentAsyncDetectResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchContentAsyncDetectResponse) GoString() string {
	return s.String()
}

func (s *BatchContentAsyncDetectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchContentAsyncDetectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchContentAsyncDetectResponse) GetBody() *BatchContentAsyncDetectResponseBody {
	return s.Body
}

func (s *BatchContentAsyncDetectResponse) SetHeaders(v map[string]*string) *BatchContentAsyncDetectResponse {
	s.Headers = v
	return s
}

func (s *BatchContentAsyncDetectResponse) SetStatusCode(v int32) *BatchContentAsyncDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchContentAsyncDetectResponse) SetBody(v *BatchContentAsyncDetectResponseBody) *BatchContentAsyncDetectResponse {
	s.Body = v
	return s
}

func (s *BatchContentAsyncDetectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
