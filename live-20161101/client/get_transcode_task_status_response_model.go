// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranscodeTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTranscodeTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTranscodeTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetTranscodeTaskStatusResponseBody) *GetTranscodeTaskStatusResponse
	GetBody() *GetTranscodeTaskStatusResponseBody
}

type GetTranscodeTaskStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTranscodeTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTranscodeTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetTranscodeTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTranscodeTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTranscodeTaskStatusResponse) GetBody() *GetTranscodeTaskStatusResponseBody {
	return s.Body
}

func (s *GetTranscodeTaskStatusResponse) SetHeaders(v map[string]*string) *GetTranscodeTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetTranscodeTaskStatusResponse) SetStatusCode(v int32) *GetTranscodeTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTranscodeTaskStatusResponse) SetBody(v *GetTranscodeTaskStatusResponseBody) *GetTranscodeTaskStatusResponse {
	s.Body = v
	return s
}

func (s *GetTranscodeTaskStatusResponse) Validate() error {
	return dara.Validate(s)
}
