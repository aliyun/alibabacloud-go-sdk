// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatFlowMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChatFlowMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChatFlowMetricResponse
	GetStatusCode() *int32
	SetBody(v *GetChatFlowMetricResponseBody) *GetChatFlowMetricResponse
	GetBody() *GetChatFlowMetricResponseBody
}

type GetChatFlowMetricResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatFlowMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatFlowMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChatFlowMetricResponse) GoString() string {
	return s.String()
}

func (s *GetChatFlowMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChatFlowMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChatFlowMetricResponse) GetBody() *GetChatFlowMetricResponseBody {
	return s.Body
}

func (s *GetChatFlowMetricResponse) SetHeaders(v map[string]*string) *GetChatFlowMetricResponse {
	s.Headers = v
	return s
}

func (s *GetChatFlowMetricResponse) SetStatusCode(v int32) *GetChatFlowMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatFlowMetricResponse) SetBody(v *GetChatFlowMetricResponseBody) *GetChatFlowMetricResponse {
	s.Body = v
	return s
}

func (s *GetChatFlowMetricResponse) Validate() error {
	return dara.Validate(s)
}
