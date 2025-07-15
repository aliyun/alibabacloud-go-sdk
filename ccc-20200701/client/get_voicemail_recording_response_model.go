// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVoicemailRecordingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVoicemailRecordingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVoicemailRecordingResponse
	GetStatusCode() *int32
	SetBody(v *GetVoicemailRecordingResponseBody) *GetVoicemailRecordingResponse
	GetBody() *GetVoicemailRecordingResponseBody
}

type GetVoicemailRecordingResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVoicemailRecordingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVoicemailRecordingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVoicemailRecordingResponse) GoString() string {
	return s.String()
}

func (s *GetVoicemailRecordingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVoicemailRecordingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVoicemailRecordingResponse) GetBody() *GetVoicemailRecordingResponseBody {
	return s.Body
}

func (s *GetVoicemailRecordingResponse) SetHeaders(v map[string]*string) *GetVoicemailRecordingResponse {
	s.Headers = v
	return s
}

func (s *GetVoicemailRecordingResponse) SetStatusCode(v int32) *GetVoicemailRecordingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVoicemailRecordingResponse) SetBody(v *GetVoicemailRecordingResponseBody) *GetVoicemailRecordingResponse {
	s.Body = v
	return s
}

func (s *GetVoicemailRecordingResponse) Validate() error {
	return dara.Validate(s)
}
