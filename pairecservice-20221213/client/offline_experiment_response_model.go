// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineExperimentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OfflineExperimentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OfflineExperimentResponse
	GetStatusCode() *int32
	SetBody(v *OfflineExperimentResponseBody) *OfflineExperimentResponse
	GetBody() *OfflineExperimentResponseBody
}

type OfflineExperimentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OfflineExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OfflineExperimentResponse) String() string {
	return dara.Prettify(s)
}

func (s OfflineExperimentResponse) GoString() string {
	return s.String()
}

func (s *OfflineExperimentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OfflineExperimentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OfflineExperimentResponse) GetBody() *OfflineExperimentResponseBody {
	return s.Body
}

func (s *OfflineExperimentResponse) SetHeaders(v map[string]*string) *OfflineExperimentResponse {
	s.Headers = v
	return s
}

func (s *OfflineExperimentResponse) SetStatusCode(v int32) *OfflineExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *OfflineExperimentResponse) SetBody(v *OfflineExperimentResponseBody) *OfflineExperimentResponse {
	s.Body = v
	return s
}

func (s *OfflineExperimentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
