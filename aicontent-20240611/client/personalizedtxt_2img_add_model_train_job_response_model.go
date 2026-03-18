// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedtxt2imgAddModelTrainJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *Personalizedtxt2imgAddModelTrainJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *Personalizedtxt2imgAddModelTrainJobResponse
	GetStatusCode() *int32
	SetBody(v *Personalizedtxt2imgAddModelTrainJobResponseBody) *Personalizedtxt2imgAddModelTrainJobResponse
	GetBody() *Personalizedtxt2imgAddModelTrainJobResponseBody
}

type Personalizedtxt2imgAddModelTrainJobResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Personalizedtxt2imgAddModelTrainJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Personalizedtxt2imgAddModelTrainJobResponse) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgAddModelTrainJobResponse) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddModelTrainJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *Personalizedtxt2imgAddModelTrainJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *Personalizedtxt2imgAddModelTrainJobResponse) GetBody() *Personalizedtxt2imgAddModelTrainJobResponseBody {
	return s.Body
}

func (s *Personalizedtxt2imgAddModelTrainJobResponse) SetHeaders(v map[string]*string) *Personalizedtxt2imgAddModelTrainJobResponse {
	s.Headers = v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponse) SetStatusCode(v int32) *Personalizedtxt2imgAddModelTrainJobResponse {
	s.StatusCode = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponse) SetBody(v *Personalizedtxt2imgAddModelTrainJobResponseBody) *Personalizedtxt2imgAddModelTrainJobResponse {
	s.Body = v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
