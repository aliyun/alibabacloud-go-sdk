// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCustomHotTopicBroadcastJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitCustomHotTopicBroadcastJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitCustomHotTopicBroadcastJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitCustomHotTopicBroadcastJobResponseBody) *SubmitCustomHotTopicBroadcastJobResponse
	GetBody() *SubmitCustomHotTopicBroadcastJobResponseBody
}

type SubmitCustomHotTopicBroadcastJobResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitCustomHotTopicBroadcastJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitCustomHotTopicBroadcastJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomHotTopicBroadcastJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitCustomHotTopicBroadcastJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitCustomHotTopicBroadcastJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitCustomHotTopicBroadcastJobResponse) GetBody() *SubmitCustomHotTopicBroadcastJobResponseBody {
	return s.Body
}

func (s *SubmitCustomHotTopicBroadcastJobResponse) SetHeaders(v map[string]*string) *SubmitCustomHotTopicBroadcastJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobResponse) SetStatusCode(v int32) *SubmitCustomHotTopicBroadcastJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobResponse) SetBody(v *SubmitCustomHotTopicBroadcastJobResponseBody) *SubmitCustomHotTopicBroadcastJobResponse {
	s.Body = v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
