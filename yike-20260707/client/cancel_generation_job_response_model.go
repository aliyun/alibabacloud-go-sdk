// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelGenerationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelGenerationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelGenerationJobResponse
	GetStatusCode() *int32
	SetBody(v *CancelGenerationJobResponseBody) *CancelGenerationJobResponse
	GetBody() *CancelGenerationJobResponseBody
}

type CancelGenerationJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelGenerationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelGenerationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelGenerationJobResponse) GoString() string {
	return s.String()
}

func (s *CancelGenerationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelGenerationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelGenerationJobResponse) GetBody() *CancelGenerationJobResponseBody {
	return s.Body
}

func (s *CancelGenerationJobResponse) SetHeaders(v map[string]*string) *CancelGenerationJobResponse {
	s.Headers = v
	return s
}

func (s *CancelGenerationJobResponse) SetStatusCode(v int32) *CancelGenerationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelGenerationJobResponse) SetBody(v *CancelGenerationJobResponseBody) *CancelGenerationJobResponse {
	s.Body = v
	return s
}

func (s *CancelGenerationJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
