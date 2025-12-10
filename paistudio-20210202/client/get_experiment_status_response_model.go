// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExperimentStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExperimentStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetExperimentStatusResponseBody) *GetExperimentStatusResponse
	GetBody() *GetExperimentStatusResponseBody
}

type GetExperimentStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExperimentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExperimentStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentStatusResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExperimentStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExperimentStatusResponse) GetBody() *GetExperimentStatusResponseBody {
	return s.Body
}

func (s *GetExperimentStatusResponse) SetHeaders(v map[string]*string) *GetExperimentStatusResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentStatusResponse) SetStatusCode(v int32) *GetExperimentStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentStatusResponse) SetBody(v *GetExperimentStatusResponseBody) *GetExperimentStatusResponse {
	s.Body = v
	return s
}

func (s *GetExperimentStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
