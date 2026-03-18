// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedtxt2imgAddInferenceJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *Personalizedtxt2imgAddInferenceJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *Personalizedtxt2imgAddInferenceJobResponse
	GetStatusCode() *int32
	SetBody(v *Personalizedtxt2imgAddInferenceJobResponseBody) *Personalizedtxt2imgAddInferenceJobResponse
	GetBody() *Personalizedtxt2imgAddInferenceJobResponseBody
}

type Personalizedtxt2imgAddInferenceJobResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Personalizedtxt2imgAddInferenceJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Personalizedtxt2imgAddInferenceJobResponse) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgAddInferenceJobResponse) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddInferenceJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *Personalizedtxt2imgAddInferenceJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *Personalizedtxt2imgAddInferenceJobResponse) GetBody() *Personalizedtxt2imgAddInferenceJobResponseBody {
	return s.Body
}

func (s *Personalizedtxt2imgAddInferenceJobResponse) SetHeaders(v map[string]*string) *Personalizedtxt2imgAddInferenceJobResponse {
	s.Headers = v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponse) SetStatusCode(v int32) *Personalizedtxt2imgAddInferenceJobResponse {
	s.StatusCode = &v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponse) SetBody(v *Personalizedtxt2imgAddInferenceJobResponseBody) *Personalizedtxt2imgAddInferenceJobResponse {
	s.Body = v
	return s
}

func (s *Personalizedtxt2imgAddInferenceJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
