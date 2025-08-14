// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAudioProduceJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAudioProduceJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAudioProduceJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAudioProduceJobResponseBody) *SubmitAudioProduceJobResponse
	GetBody() *SubmitAudioProduceJobResponseBody
}

type SubmitAudioProduceJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAudioProduceJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAudioProduceJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAudioProduceJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAudioProduceJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAudioProduceJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAudioProduceJobResponse) GetBody() *SubmitAudioProduceJobResponseBody {
	return s.Body
}

func (s *SubmitAudioProduceJobResponse) SetHeaders(v map[string]*string) *SubmitAudioProduceJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAudioProduceJobResponse) SetStatusCode(v int32) *SubmitAudioProduceJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAudioProduceJobResponse) SetBody(v *SubmitAudioProduceJobResponseBody) *SubmitAudioProduceJobResponse {
	s.Body = v
	return s
}

func (s *SubmitAudioProduceJobResponse) Validate() error {
	return dara.Validate(s)
}
