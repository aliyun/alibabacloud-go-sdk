// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAlertStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAlertStrategyResponse
	GetStatusCode() *int32
	SetBody(v *CreateAlertStrategyResponseBody) *CreateAlertStrategyResponse
	GetBody() *CreateAlertStrategyResponseBody
}

type CreateAlertStrategyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAlertStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAlertStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertStrategyResponse) GoString() string {
	return s.String()
}

func (s *CreateAlertStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAlertStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAlertStrategyResponse) GetBody() *CreateAlertStrategyResponseBody {
	return s.Body
}

func (s *CreateAlertStrategyResponse) SetHeaders(v map[string]*string) *CreateAlertStrategyResponse {
	s.Headers = v
	return s
}

func (s *CreateAlertStrategyResponse) SetStatusCode(v int32) *CreateAlertStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlertStrategyResponse) SetBody(v *CreateAlertStrategyResponseBody) *CreateAlertStrategyResponse {
	s.Body = v
	return s
}

func (s *CreateAlertStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
