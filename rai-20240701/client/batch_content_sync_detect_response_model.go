// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchContentSyncDetectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchContentSyncDetectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchContentSyncDetectResponse
	GetStatusCode() *int32
	SetBody(v *BatchContentSyncDetectResponseBody) *BatchContentSyncDetectResponse
	GetBody() *BatchContentSyncDetectResponseBody
}

type BatchContentSyncDetectResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchContentSyncDetectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchContentSyncDetectResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchContentSyncDetectResponse) GoString() string {
	return s.String()
}

func (s *BatchContentSyncDetectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchContentSyncDetectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchContentSyncDetectResponse) GetBody() *BatchContentSyncDetectResponseBody {
	return s.Body
}

func (s *BatchContentSyncDetectResponse) SetHeaders(v map[string]*string) *BatchContentSyncDetectResponse {
	s.Headers = v
	return s
}

func (s *BatchContentSyncDetectResponse) SetStatusCode(v int32) *BatchContentSyncDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchContentSyncDetectResponse) SetBody(v *BatchContentSyncDetectResponseBody) *BatchContentSyncDetectResponse {
	s.Body = v
	return s
}

func (s *BatchContentSyncDetectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
