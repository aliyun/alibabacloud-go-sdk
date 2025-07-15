// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartEdgeTranscodeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartEdgeTranscodeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartEdgeTranscodeJobResponse
	GetStatusCode() *int32
	SetBody(v *StartEdgeTranscodeJobResponseBody) *StartEdgeTranscodeJobResponse
	GetBody() *StartEdgeTranscodeJobResponseBody
}

type StartEdgeTranscodeJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartEdgeTranscodeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartEdgeTranscodeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StartEdgeTranscodeJobResponse) GoString() string {
	return s.String()
}

func (s *StartEdgeTranscodeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartEdgeTranscodeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartEdgeTranscodeJobResponse) GetBody() *StartEdgeTranscodeJobResponseBody {
	return s.Body
}

func (s *StartEdgeTranscodeJobResponse) SetHeaders(v map[string]*string) *StartEdgeTranscodeJobResponse {
	s.Headers = v
	return s
}

func (s *StartEdgeTranscodeJobResponse) SetStatusCode(v int32) *StartEdgeTranscodeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StartEdgeTranscodeJobResponse) SetBody(v *StartEdgeTranscodeJobResponseBody) *StartEdgeTranscodeJobResponse {
	s.Body = v
	return s
}

func (s *StartEdgeTranscodeJobResponse) Validate() error {
	return dara.Validate(s)
}
