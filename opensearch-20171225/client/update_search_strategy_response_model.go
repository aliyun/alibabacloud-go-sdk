// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSearchStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSearchStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSearchStrategyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSearchStrategyResponseBody) *UpdateSearchStrategyResponse
	GetBody() *UpdateSearchStrategyResponseBody
}

type UpdateSearchStrategyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSearchStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSearchStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSearchStrategyResponse) GoString() string {
	return s.String()
}

func (s *UpdateSearchStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSearchStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSearchStrategyResponse) GetBody() *UpdateSearchStrategyResponseBody {
	return s.Body
}

func (s *UpdateSearchStrategyResponse) SetHeaders(v map[string]*string) *UpdateSearchStrategyResponse {
	s.Headers = v
	return s
}

func (s *UpdateSearchStrategyResponse) SetStatusCode(v int32) *UpdateSearchStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSearchStrategyResponse) SetBody(v *UpdateSearchStrategyResponseBody) *UpdateSearchStrategyResponse {
	s.Body = v
	return s
}

func (s *UpdateSearchStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
