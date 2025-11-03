// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAlertStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAlertStrategyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAlertStrategyResponseBody) *UpdateAlertStrategyResponse
	GetBody() *UpdateAlertStrategyResponseBody
}

type UpdateAlertStrategyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAlertStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAlertStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertStrategyResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlertStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAlertStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAlertStrategyResponse) GetBody() *UpdateAlertStrategyResponseBody {
	return s.Body
}

func (s *UpdateAlertStrategyResponse) SetHeaders(v map[string]*string) *UpdateAlertStrategyResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlertStrategyResponse) SetStatusCode(v int32) *UpdateAlertStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlertStrategyResponse) SetBody(v *UpdateAlertStrategyResponseBody) *UpdateAlertStrategyResponse {
	s.Body = v
	return s
}

func (s *UpdateAlertStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
