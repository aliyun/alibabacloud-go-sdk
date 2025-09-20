// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeTriggerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeTriggerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeTriggerResponse
	GetStatusCode() *int32
	SetBody(v *ResumeTriggerResponseBody) *ResumeTriggerResponse
	GetBody() *ResumeTriggerResponseBody
}

type ResumeTriggerResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeTriggerResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeTriggerResponse) GoString() string {
	return s.String()
}

func (s *ResumeTriggerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeTriggerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeTriggerResponse) GetBody() *ResumeTriggerResponseBody {
	return s.Body
}

func (s *ResumeTriggerResponse) SetHeaders(v map[string]*string) *ResumeTriggerResponse {
	s.Headers = v
	return s
}

func (s *ResumeTriggerResponse) SetStatusCode(v int32) *ResumeTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeTriggerResponse) SetBody(v *ResumeTriggerResponseBody) *ResumeTriggerResponse {
	s.Body = v
	return s
}

func (s *ResumeTriggerResponse) Validate() error {
	return dara.Validate(s)
}
