// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunPptOutlineGenerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunPptOutlineGenerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunPptOutlineGenerationResponse
	GetStatusCode() *int32
	SetBody(v *RunPptOutlineGenerationResponseBody) *RunPptOutlineGenerationResponse
	GetBody() *RunPptOutlineGenerationResponseBody
}

type RunPptOutlineGenerationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunPptOutlineGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunPptOutlineGenerationResponse) String() string {
	return dara.Prettify(s)
}

func (s RunPptOutlineGenerationResponse) GoString() string {
	return s.String()
}

func (s *RunPptOutlineGenerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunPptOutlineGenerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunPptOutlineGenerationResponse) GetBody() *RunPptOutlineGenerationResponseBody {
	return s.Body
}

func (s *RunPptOutlineGenerationResponse) SetHeaders(v map[string]*string) *RunPptOutlineGenerationResponse {
	s.Headers = v
	return s
}

func (s *RunPptOutlineGenerationResponse) SetStatusCode(v int32) *RunPptOutlineGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunPptOutlineGenerationResponse) SetBody(v *RunPptOutlineGenerationResponseBody) *RunPptOutlineGenerationResponse {
	s.Body = v
	return s
}

func (s *RunPptOutlineGenerationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
