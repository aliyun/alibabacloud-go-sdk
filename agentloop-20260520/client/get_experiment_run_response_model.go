// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExperimentRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExperimentRunResponse
	GetStatusCode() *int32
	SetBody(v *GetExperimentRunResponseBody) *GetExperimentRunResponse
	GetBody() *GetExperimentRunResponseBody
}

type GetExperimentRunResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExperimentRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExperimentRunResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentRunResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExperimentRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExperimentRunResponse) GetBody() *GetExperimentRunResponseBody {
	return s.Body
}

func (s *GetExperimentRunResponse) SetHeaders(v map[string]*string) *GetExperimentRunResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentRunResponse) SetStatusCode(v int32) *GetExperimentRunResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentRunResponse) SetBody(v *GetExperimentRunResponseBody) *GetExperimentRunResponse {
	s.Body = v
	return s
}

func (s *GetExperimentRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
