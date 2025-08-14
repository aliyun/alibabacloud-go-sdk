// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveTranscodeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLiveTranscodeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLiveTranscodeJobResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLiveTranscodeJobResponseBody) *UpdateLiveTranscodeJobResponse
	GetBody() *UpdateLiveTranscodeJobResponseBody
}

type UpdateLiveTranscodeJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveTranscodeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveTranscodeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveTranscodeJobResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveTranscodeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLiveTranscodeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLiveTranscodeJobResponse) GetBody() *UpdateLiveTranscodeJobResponseBody {
	return s.Body
}

func (s *UpdateLiveTranscodeJobResponse) SetHeaders(v map[string]*string) *UpdateLiveTranscodeJobResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveTranscodeJobResponse) SetStatusCode(v int32) *UpdateLiveTranscodeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveTranscodeJobResponse) SetBody(v *UpdateLiveTranscodeJobResponseBody) *UpdateLiveTranscodeJobResponse {
	s.Body = v
	return s
}

func (s *UpdateLiveTranscodeJobResponse) Validate() error {
	return dara.Validate(s)
}
