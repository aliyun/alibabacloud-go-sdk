// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishExperimentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishExperimentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishExperimentResponse
	GetStatusCode() *int32
	SetBody(v *PublishExperimentResponseBody) *PublishExperimentResponse
	GetBody() *PublishExperimentResponseBody
}

type PublishExperimentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishExperimentResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishExperimentResponse) GoString() string {
	return s.String()
}

func (s *PublishExperimentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishExperimentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishExperimentResponse) GetBody() *PublishExperimentResponseBody {
	return s.Body
}

func (s *PublishExperimentResponse) SetHeaders(v map[string]*string) *PublishExperimentResponse {
	s.Headers = v
	return s
}

func (s *PublishExperimentResponse) SetStatusCode(v int32) *PublishExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishExperimentResponse) SetBody(v *PublishExperimentResponseBody) *PublishExperimentResponse {
	s.Body = v
	return s
}

func (s *PublishExperimentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
