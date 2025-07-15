// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEdgeTranscodeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEdgeTranscodeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEdgeTranscodeJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateEdgeTranscodeJobResponseBody) *CreateEdgeTranscodeJobResponse
	GetBody() *CreateEdgeTranscodeJobResponseBody
}

type CreateEdgeTranscodeJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEdgeTranscodeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEdgeTranscodeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeTranscodeJobResponse) GoString() string {
	return s.String()
}

func (s *CreateEdgeTranscodeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEdgeTranscodeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEdgeTranscodeJobResponse) GetBody() *CreateEdgeTranscodeJobResponseBody {
	return s.Body
}

func (s *CreateEdgeTranscodeJobResponse) SetHeaders(v map[string]*string) *CreateEdgeTranscodeJobResponse {
	s.Headers = v
	return s
}

func (s *CreateEdgeTranscodeJobResponse) SetStatusCode(v int32) *CreateEdgeTranscodeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEdgeTranscodeJobResponse) SetBody(v *CreateEdgeTranscodeJobResponseBody) *CreateEdgeTranscodeJobResponse {
	s.Body = v
	return s
}

func (s *CreateEdgeTranscodeJobResponse) Validate() error {
	return dara.Validate(s)
}
