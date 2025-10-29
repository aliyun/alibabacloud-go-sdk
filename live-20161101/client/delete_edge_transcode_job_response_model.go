// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEdgeTranscodeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEdgeTranscodeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEdgeTranscodeJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEdgeTranscodeJobResponseBody) *DeleteEdgeTranscodeJobResponse
	GetBody() *DeleteEdgeTranscodeJobResponseBody
}

type DeleteEdgeTranscodeJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEdgeTranscodeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEdgeTranscodeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEdgeTranscodeJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteEdgeTranscodeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEdgeTranscodeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEdgeTranscodeJobResponse) GetBody() *DeleteEdgeTranscodeJobResponseBody {
	return s.Body
}

func (s *DeleteEdgeTranscodeJobResponse) SetHeaders(v map[string]*string) *DeleteEdgeTranscodeJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteEdgeTranscodeJobResponse) SetStatusCode(v int32) *DeleteEdgeTranscodeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEdgeTranscodeJobResponse) SetBody(v *DeleteEdgeTranscodeJobResponseBody) *DeleteEdgeTranscodeJobResponse {
	s.Body = v
	return s
}

func (s *DeleteEdgeTranscodeJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
