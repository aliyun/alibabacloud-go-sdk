// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoGenerationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoGenerationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoGenerationJobResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoGenerationJobResponseBody) *GetVideoGenerationJobResponse
	GetBody() *GetVideoGenerationJobResponseBody
}

type GetVideoGenerationJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoGenerationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoGenerationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoGenerationJobResponse) GoString() string {
	return s.String()
}

func (s *GetVideoGenerationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoGenerationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoGenerationJobResponse) GetBody() *GetVideoGenerationJobResponseBody {
	return s.Body
}

func (s *GetVideoGenerationJobResponse) SetHeaders(v map[string]*string) *GetVideoGenerationJobResponse {
	s.Headers = v
	return s
}

func (s *GetVideoGenerationJobResponse) SetStatusCode(v int32) *GetVideoGenerationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoGenerationJobResponse) SetBody(v *GetVideoGenerationJobResponseBody) *GetVideoGenerationJobResponse {
	s.Body = v
	return s
}

func (s *GetVideoGenerationJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
