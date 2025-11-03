// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAlertStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAlertStrategyResponse
	GetStatusCode() *int32
	SetBody(v *GetAlertStrategyResponseBody) *GetAlertStrategyResponse
	GetBody() *GetAlertStrategyResponseBody
}

type GetAlertStrategyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlertStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlertStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAlertStrategyResponse) GoString() string {
	return s.String()
}

func (s *GetAlertStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAlertStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAlertStrategyResponse) GetBody() *GetAlertStrategyResponseBody {
	return s.Body
}

func (s *GetAlertStrategyResponse) SetHeaders(v map[string]*string) *GetAlertStrategyResponse {
	s.Headers = v
	return s
}

func (s *GetAlertStrategyResponse) SetStatusCode(v int32) *GetAlertStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlertStrategyResponse) SetBody(v *GetAlertStrategyResponseBody) *GetAlertStrategyResponse {
	s.Body = v
	return s
}

func (s *GetAlertStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
