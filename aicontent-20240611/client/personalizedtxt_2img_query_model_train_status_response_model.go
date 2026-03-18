// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedtxt2imgQueryModelTrainStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *Personalizedtxt2imgQueryModelTrainStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *Personalizedtxt2imgQueryModelTrainStatusResponse
	GetStatusCode() *int32
	SetBody(v *Personalizedtxt2imgQueryModelTrainStatusResponseBody) *Personalizedtxt2imgQueryModelTrainStatusResponse
	GetBody() *Personalizedtxt2imgQueryModelTrainStatusResponseBody
}

type Personalizedtxt2imgQueryModelTrainStatusResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Personalizedtxt2imgQueryModelTrainStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Personalizedtxt2imgQueryModelTrainStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgQueryModelTrainStatusResponse) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponse) GetBody() *Personalizedtxt2imgQueryModelTrainStatusResponseBody {
	return s.Body
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponse) SetHeaders(v map[string]*string) *Personalizedtxt2imgQueryModelTrainStatusResponse {
	s.Headers = v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponse) SetStatusCode(v int32) *Personalizedtxt2imgQueryModelTrainStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponse) SetBody(v *Personalizedtxt2imgQueryModelTrainStatusResponseBody) *Personalizedtxt2imgQueryModelTrainStatusResponse {
	s.Body = v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
