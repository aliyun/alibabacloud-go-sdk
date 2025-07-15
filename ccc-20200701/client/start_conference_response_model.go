// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartConferenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartConferenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartConferenceResponse
	GetStatusCode() *int32
	SetBody(v *StartConferenceResponseBody) *StartConferenceResponse
	GetBody() *StartConferenceResponseBody
}

type StartConferenceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartConferenceResponse) String() string {
	return dara.Prettify(s)
}

func (s StartConferenceResponse) GoString() string {
	return s.String()
}

func (s *StartConferenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartConferenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartConferenceResponse) GetBody() *StartConferenceResponseBody {
	return s.Body
}

func (s *StartConferenceResponse) SetHeaders(v map[string]*string) *StartConferenceResponse {
	s.Headers = v
	return s
}

func (s *StartConferenceResponse) SetStatusCode(v int32) *StartConferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartConferenceResponse) SetBody(v *StartConferenceResponseBody) *StartConferenceResponse {
	s.Body = v
	return s
}

func (s *StartConferenceResponse) Validate() error {
	return dara.Validate(s)
}
