// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExperimentPlansResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExperimentPlansResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExperimentPlansResponse
	GetStatusCode() *int32
	SetBody(v *ListExperimentPlansResponseBody) *ListExperimentPlansResponse
	GetBody() *ListExperimentPlansResponseBody
}

type ListExperimentPlansResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExperimentPlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExperimentPlansResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExperimentPlansResponse) GoString() string {
	return s.String()
}

func (s *ListExperimentPlansResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExperimentPlansResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExperimentPlansResponse) GetBody() *ListExperimentPlansResponseBody {
	return s.Body
}

func (s *ListExperimentPlansResponse) SetHeaders(v map[string]*string) *ListExperimentPlansResponse {
	s.Headers = v
	return s
}

func (s *ListExperimentPlansResponse) SetStatusCode(v int32) *ListExperimentPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExperimentPlansResponse) SetBody(v *ListExperimentPlansResponseBody) *ListExperimentPlansResponse {
	s.Body = v
	return s
}

func (s *ListExperimentPlansResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
