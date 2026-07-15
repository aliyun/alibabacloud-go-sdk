// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageGenerationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetImageGenerationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetImageGenerationJobResponse
	GetStatusCode() *int32
	SetBody(v *GetImageGenerationJobResponseBody) *GetImageGenerationJobResponse
	GetBody() *GetImageGenerationJobResponseBody
}

type GetImageGenerationJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageGenerationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageGenerationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetImageGenerationJobResponse) GoString() string {
	return s.String()
}

func (s *GetImageGenerationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetImageGenerationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetImageGenerationJobResponse) GetBody() *GetImageGenerationJobResponseBody {
	return s.Body
}

func (s *GetImageGenerationJobResponse) SetHeaders(v map[string]*string) *GetImageGenerationJobResponse {
	s.Headers = v
	return s
}

func (s *GetImageGenerationJobResponse) SetStatusCode(v int32) *GetImageGenerationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageGenerationJobResponse) SetBody(v *GetImageGenerationJobResponseBody) *GetImageGenerationJobResponse {
	s.Body = v
	return s
}

func (s *GetImageGenerationJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
