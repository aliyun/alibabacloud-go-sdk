// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyExperimentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyExperimentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyExperimentResponse
	GetStatusCode() *int32
	SetBody(v *CopyExperimentResponseBody) *CopyExperimentResponse
	GetBody() *CopyExperimentResponseBody
}

type CopyExperimentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyExperimentResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyExperimentResponse) GoString() string {
	return s.String()
}

func (s *CopyExperimentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyExperimentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyExperimentResponse) GetBody() *CopyExperimentResponseBody {
	return s.Body
}

func (s *CopyExperimentResponse) SetHeaders(v map[string]*string) *CopyExperimentResponse {
	s.Headers = v
	return s
}

func (s *CopyExperimentResponse) SetStatusCode(v int32) *CopyExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyExperimentResponse) SetBody(v *CopyExperimentResponseBody) *CopyExperimentResponse {
	s.Body = v
	return s
}

func (s *CopyExperimentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
