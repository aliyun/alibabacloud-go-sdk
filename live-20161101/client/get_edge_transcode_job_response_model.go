// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeTranscodeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEdgeTranscodeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEdgeTranscodeJobResponse
	GetStatusCode() *int32
	SetBody(v *GetEdgeTranscodeJobResponseBody) *GetEdgeTranscodeJobResponse
	GetBody() *GetEdgeTranscodeJobResponseBody
}

type GetEdgeTranscodeJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEdgeTranscodeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEdgeTranscodeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeTranscodeJobResponse) GoString() string {
	return s.String()
}

func (s *GetEdgeTranscodeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEdgeTranscodeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEdgeTranscodeJobResponse) GetBody() *GetEdgeTranscodeJobResponseBody {
	return s.Body
}

func (s *GetEdgeTranscodeJobResponse) SetHeaders(v map[string]*string) *GetEdgeTranscodeJobResponse {
	s.Headers = v
	return s
}

func (s *GetEdgeTranscodeJobResponse) SetStatusCode(v int32) *GetEdgeTranscodeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEdgeTranscodeJobResponse) SetBody(v *GetEdgeTranscodeJobResponseBody) *GetEdgeTranscodeJobResponse {
	s.Body = v
	return s
}

func (s *GetEdgeTranscodeJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
