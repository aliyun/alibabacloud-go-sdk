// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedtxt2imgQueryModelTrainJobListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *Personalizedtxt2imgQueryModelTrainJobListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *Personalizedtxt2imgQueryModelTrainJobListResponse
	GetStatusCode() *int32
	SetBody(v *Personalizedtxt2imgQueryModelTrainJobListResponseBody) *Personalizedtxt2imgQueryModelTrainJobListResponse
	GetBody() *Personalizedtxt2imgQueryModelTrainJobListResponseBody
}

type Personalizedtxt2imgQueryModelTrainJobListResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Personalizedtxt2imgQueryModelTrainJobListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Personalizedtxt2imgQueryModelTrainJobListResponse) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgQueryModelTrainJobListResponse) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponse) GetBody() *Personalizedtxt2imgQueryModelTrainJobListResponseBody {
	return s.Body
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponse) SetHeaders(v map[string]*string) *Personalizedtxt2imgQueryModelTrainJobListResponse {
	s.Headers = v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponse) SetStatusCode(v int32) *Personalizedtxt2imgQueryModelTrainJobListResponse {
	s.StatusCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponse) SetBody(v *Personalizedtxt2imgQueryModelTrainJobListResponseBody) *Personalizedtxt2imgQueryModelTrainJobListResponse {
	s.Body = v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainJobListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
