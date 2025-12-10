// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExperimentMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExperimentMetaResponse
	GetStatusCode() *int32
	SetBody(v *GetExperimentMetaResponseBody) *GetExperimentMetaResponse
	GetBody() *GetExperimentMetaResponseBody
}

type GetExperimentMetaResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExperimentMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExperimentMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentMetaResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExperimentMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExperimentMetaResponse) GetBody() *GetExperimentMetaResponseBody {
	return s.Body
}

func (s *GetExperimentMetaResponse) SetHeaders(v map[string]*string) *GetExperimentMetaResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentMetaResponse) SetStatusCode(v int32) *GetExperimentMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentMetaResponse) SetBody(v *GetExperimentMetaResponseBody) *GetExperimentMetaResponse {
	s.Body = v
	return s
}

func (s *GetExperimentMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
