// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomHotTopicBroadcastJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomHotTopicBroadcastJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomHotTopicBroadcastJobResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomHotTopicBroadcastJobResponseBody) *GetCustomHotTopicBroadcastJobResponse
	GetBody() *GetCustomHotTopicBroadcastJobResponseBody
}

type GetCustomHotTopicBroadcastJobResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomHotTopicBroadcastJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomHotTopicBroadcastJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomHotTopicBroadcastJobResponse) GoString() string {
	return s.String()
}

func (s *GetCustomHotTopicBroadcastJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomHotTopicBroadcastJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomHotTopicBroadcastJobResponse) GetBody() *GetCustomHotTopicBroadcastJobResponseBody {
	return s.Body
}

func (s *GetCustomHotTopicBroadcastJobResponse) SetHeaders(v map[string]*string) *GetCustomHotTopicBroadcastJobResponse {
	s.Headers = v
	return s
}

func (s *GetCustomHotTopicBroadcastJobResponse) SetStatusCode(v int32) *GetCustomHotTopicBroadcastJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomHotTopicBroadcastJobResponse) SetBody(v *GetCustomHotTopicBroadcastJobResponseBody) *GetCustomHotTopicBroadcastJobResponse {
	s.Body = v
	return s
}

func (s *GetCustomHotTopicBroadcastJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
