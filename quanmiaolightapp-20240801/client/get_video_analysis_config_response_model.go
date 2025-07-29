// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoAnalysisConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoAnalysisConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoAnalysisConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoAnalysisConfigResponseBody) *GetVideoAnalysisConfigResponse
	GetBody() *GetVideoAnalysisConfigResponseBody
}

type GetVideoAnalysisConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoAnalysisConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoAnalysisConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisConfigResponse) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoAnalysisConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoAnalysisConfigResponse) GetBody() *GetVideoAnalysisConfigResponseBody {
	return s.Body
}

func (s *GetVideoAnalysisConfigResponse) SetHeaders(v map[string]*string) *GetVideoAnalysisConfigResponse {
	s.Headers = v
	return s
}

func (s *GetVideoAnalysisConfigResponse) SetStatusCode(v int32) *GetVideoAnalysisConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoAnalysisConfigResponse) SetBody(v *GetVideoAnalysisConfigResponseBody) *GetVideoAnalysisConfigResponse {
	s.Body = v
	return s
}

func (s *GetVideoAnalysisConfigResponse) Validate() error {
	return dara.Validate(s)
}
