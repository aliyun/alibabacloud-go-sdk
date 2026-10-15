// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDistributionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDistributionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDistributionResponse
	GetStatusCode() *int32
	SetBody(v *CreateDistributionResponseBody) *CreateDistributionResponse
	GetBody() *CreateDistributionResponseBody
}

type CreateDistributionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDistributionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDistributionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDistributionResponse) GoString() string {
	return s.String()
}

func (s *CreateDistributionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDistributionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDistributionResponse) GetBody() *CreateDistributionResponseBody {
	return s.Body
}

func (s *CreateDistributionResponse) SetHeaders(v map[string]*string) *CreateDistributionResponse {
	s.Headers = v
	return s
}

func (s *CreateDistributionResponse) SetStatusCode(v int32) *CreateDistributionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDistributionResponse) SetBody(v *CreateDistributionResponseBody) *CreateDistributionResponse {
	s.Body = v
	return s
}

func (s *CreateDistributionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
