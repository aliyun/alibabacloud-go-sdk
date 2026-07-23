// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventAnalysisJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEventAnalysisJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEventAnalysisJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEventAnalysisJobResponseBody) *DeleteEventAnalysisJobResponse
	GetBody() *DeleteEventAnalysisJobResponseBody
}

type DeleteEventAnalysisJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEventAnalysisJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEventAnalysisJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventAnalysisJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventAnalysisJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEventAnalysisJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEventAnalysisJobResponse) GetBody() *DeleteEventAnalysisJobResponseBody {
	return s.Body
}

func (s *DeleteEventAnalysisJobResponse) SetHeaders(v map[string]*string) *DeleteEventAnalysisJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventAnalysisJobResponse) SetStatusCode(v int32) *DeleteEventAnalysisJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventAnalysisJobResponse) SetBody(v *DeleteEventAnalysisJobResponseBody) *DeleteEventAnalysisJobResponse {
	s.Body = v
	return s
}

func (s *DeleteEventAnalysisJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
