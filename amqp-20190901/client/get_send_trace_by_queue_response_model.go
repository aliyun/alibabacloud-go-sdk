// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSendTraceByQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSendTraceByQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSendTraceByQueueResponse
	GetStatusCode() *int32
	SetBody(v *GetSendTraceByQueueResponseBody) *GetSendTraceByQueueResponse
	GetBody() *GetSendTraceByQueueResponseBody
}

type GetSendTraceByQueueResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSendTraceByQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSendTraceByQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSendTraceByQueueResponse) GoString() string {
	return s.String()
}

func (s *GetSendTraceByQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSendTraceByQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSendTraceByQueueResponse) GetBody() *GetSendTraceByQueueResponseBody {
	return s.Body
}

func (s *GetSendTraceByQueueResponse) SetHeaders(v map[string]*string) *GetSendTraceByQueueResponse {
	s.Headers = v
	return s
}

func (s *GetSendTraceByQueueResponse) SetStatusCode(v int32) *GetSendTraceByQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSendTraceByQueueResponse) SetBody(v *GetSendTraceByQueueResponseBody) *GetSendTraceByQueueResponse {
	s.Body = v
	return s
}

func (s *GetSendTraceByQueueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
