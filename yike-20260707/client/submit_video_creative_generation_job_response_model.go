// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoCreativeGenerationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitVideoCreativeGenerationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitVideoCreativeGenerationJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitVideoCreativeGenerationJobResponseBody) *SubmitVideoCreativeGenerationJobResponse
	GetBody() *SubmitVideoCreativeGenerationJobResponseBody
}

type SubmitVideoCreativeGenerationJobResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitVideoCreativeGenerationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitVideoCreativeGenerationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoCreativeGenerationJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitVideoCreativeGenerationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitVideoCreativeGenerationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitVideoCreativeGenerationJobResponse) GetBody() *SubmitVideoCreativeGenerationJobResponseBody {
	return s.Body
}

func (s *SubmitVideoCreativeGenerationJobResponse) SetHeaders(v map[string]*string) *SubmitVideoCreativeGenerationJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitVideoCreativeGenerationJobResponse) SetStatusCode(v int32) *SubmitVideoCreativeGenerationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitVideoCreativeGenerationJobResponse) SetBody(v *SubmitVideoCreativeGenerationJobResponseBody) *SubmitVideoCreativeGenerationJobResponse {
	s.Body = v
	return s
}

func (s *SubmitVideoCreativeGenerationJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
