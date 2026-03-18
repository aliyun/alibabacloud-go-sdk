// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedtxt2imgQueryInferenceJobInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *Personalizedtxt2imgQueryInferenceJobInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *Personalizedtxt2imgQueryInferenceJobInfoResponse
	GetStatusCode() *int32
	SetBody(v *Personalizedtxt2imgQueryInferenceJobInfoResponseBody) *Personalizedtxt2imgQueryInferenceJobInfoResponse
	GetBody() *Personalizedtxt2imgQueryInferenceJobInfoResponseBody
}

type Personalizedtxt2imgQueryInferenceJobInfoResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Personalizedtxt2imgQueryInferenceJobInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Personalizedtxt2imgQueryInferenceJobInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgQueryInferenceJobInfoResponse) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponse) GetBody() *Personalizedtxt2imgQueryInferenceJobInfoResponseBody {
	return s.Body
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponse) SetHeaders(v map[string]*string) *Personalizedtxt2imgQueryInferenceJobInfoResponse {
	s.Headers = v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponse) SetStatusCode(v int32) *Personalizedtxt2imgQueryInferenceJobInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponse) SetBody(v *Personalizedtxt2imgQueryInferenceJobInfoResponseBody) *Personalizedtxt2imgQueryInferenceJobInfoResponse {
	s.Body = v
	return s
}

func (s *Personalizedtxt2imgQueryInferenceJobInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
