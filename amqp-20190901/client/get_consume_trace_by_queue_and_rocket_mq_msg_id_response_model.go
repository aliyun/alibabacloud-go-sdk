// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumeTraceByQueueAndRocketMqMsgIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConsumeTraceByQueueAndRocketMqMsgIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConsumeTraceByQueueAndRocketMqMsgIdResponse
	GetStatusCode() *int32
	SetBody(v *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody) *GetConsumeTraceByQueueAndRocketMqMsgIdResponse
	GetBody() *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody
}

type GetConsumeTraceByQueueAndRocketMqMsgIdResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConsumeTraceByQueueAndRocketMqMsgIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConsumeTraceByQueueAndRocketMqMsgIdResponse) GoString() string {
	return s.String()
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponse) GetBody() *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody {
	return s.Body
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponse) SetHeaders(v map[string]*string) *GetConsumeTraceByQueueAndRocketMqMsgIdResponse {
	s.Headers = v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponse) SetStatusCode(v int32) *GetConsumeTraceByQueueAndRocketMqMsgIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponse) SetBody(v *GetConsumeTraceByQueueAndRocketMqMsgIdResponseBody) *GetConsumeTraceByQueueAndRocketMqMsgIdResponse {
	s.Body = v
	return s
}

func (s *GetConsumeTraceByQueueAndRocketMqMsgIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
