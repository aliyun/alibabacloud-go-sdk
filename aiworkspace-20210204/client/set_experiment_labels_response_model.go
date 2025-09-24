// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetExperimentLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetExperimentLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetExperimentLabelsResponse
	GetStatusCode() *int32
	SetBody(v *SetExperimentLabelsResponseBody) *SetExperimentLabelsResponse
	GetBody() *SetExperimentLabelsResponseBody
}

type SetExperimentLabelsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetExperimentLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetExperimentLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s SetExperimentLabelsResponse) GoString() string {
	return s.String()
}

func (s *SetExperimentLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetExperimentLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetExperimentLabelsResponse) GetBody() *SetExperimentLabelsResponseBody {
	return s.Body
}

func (s *SetExperimentLabelsResponse) SetHeaders(v map[string]*string) *SetExperimentLabelsResponse {
	s.Headers = v
	return s
}

func (s *SetExperimentLabelsResponse) SetStatusCode(v int32) *SetExperimentLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *SetExperimentLabelsResponse) SetBody(v *SetExperimentLabelsResponseBody) *SetExperimentLabelsResponse {
	s.Body = v
	return s
}

func (s *SetExperimentLabelsResponse) Validate() error {
	return dara.Validate(s)
}
