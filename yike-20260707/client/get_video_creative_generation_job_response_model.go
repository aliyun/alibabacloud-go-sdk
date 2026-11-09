// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoCreativeGenerationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoCreativeGenerationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoCreativeGenerationJobResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoCreativeGenerationJobResponseBody) *GetVideoCreativeGenerationJobResponse
	GetBody() *GetVideoCreativeGenerationJobResponseBody
}

type GetVideoCreativeGenerationJobResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoCreativeGenerationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoCreativeGenerationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoCreativeGenerationJobResponse) GoString() string {
	return s.String()
}

func (s *GetVideoCreativeGenerationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoCreativeGenerationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoCreativeGenerationJobResponse) GetBody() *GetVideoCreativeGenerationJobResponseBody {
	return s.Body
}

func (s *GetVideoCreativeGenerationJobResponse) SetHeaders(v map[string]*string) *GetVideoCreativeGenerationJobResponse {
	s.Headers = v
	return s
}

func (s *GetVideoCreativeGenerationJobResponse) SetStatusCode(v int32) *GetVideoCreativeGenerationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoCreativeGenerationJobResponse) SetBody(v *GetVideoCreativeGenerationJobResponseBody) *GetVideoCreativeGenerationJobResponse {
	s.Body = v
	return s
}

func (s *GetVideoCreativeGenerationJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
