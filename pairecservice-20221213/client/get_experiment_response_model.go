// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExperimentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExperimentResponse
	GetStatusCode() *int32
	SetBody(v *GetExperimentResponseBody) *GetExperimentResponse
	GetBody() *GetExperimentResponseBody
}

type GetExperimentResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExperimentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExperimentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExperimentResponse) GetBody() *GetExperimentResponseBody {
	return s.Body
}

func (s *GetExperimentResponse) SetHeaders(v map[string]*string) *GetExperimentResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentResponse) SetStatusCode(v int32) *GetExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentResponse) SetBody(v *GetExperimentResponseBody) *GetExperimentResponse {
	s.Body = v
	return s
}

func (s *GetExperimentResponse) Validate() error {
	return dara.Validate(s)
}
