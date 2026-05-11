// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceRolloutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceRolloutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceRolloutResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceRolloutResponseBody) *CreateServiceRolloutResponse
	GetBody() *CreateServiceRolloutResponseBody
}

type CreateServiceRolloutResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceRolloutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceRolloutResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceRolloutResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceRolloutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceRolloutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceRolloutResponse) GetBody() *CreateServiceRolloutResponseBody {
	return s.Body
}

func (s *CreateServiceRolloutResponse) SetHeaders(v map[string]*string) *CreateServiceRolloutResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceRolloutResponse) SetStatusCode(v int32) *CreateServiceRolloutResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceRolloutResponse) SetBody(v *CreateServiceRolloutResponseBody) *CreateServiceRolloutResponse {
	s.Body = v
	return s
}

func (s *CreateServiceRolloutResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
