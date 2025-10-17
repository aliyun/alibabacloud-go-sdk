// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoAnalysisConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVideoAnalysisConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVideoAnalysisConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVideoAnalysisConfigResponseBody) *UpdateVideoAnalysisConfigResponse
	GetBody() *UpdateVideoAnalysisConfigResponseBody
}

type UpdateVideoAnalysisConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVideoAnalysisConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVideoAnalysisConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoAnalysisConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateVideoAnalysisConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVideoAnalysisConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVideoAnalysisConfigResponse) GetBody() *UpdateVideoAnalysisConfigResponseBody {
	return s.Body
}

func (s *UpdateVideoAnalysisConfigResponse) SetHeaders(v map[string]*string) *UpdateVideoAnalysisConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateVideoAnalysisConfigResponse) SetStatusCode(v int32) *UpdateVideoAnalysisConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVideoAnalysisConfigResponse) SetBody(v *UpdateVideoAnalysisConfigResponseBody) *UpdateVideoAnalysisConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateVideoAnalysisConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
