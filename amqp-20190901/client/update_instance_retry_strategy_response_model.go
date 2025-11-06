// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceRetryStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceRetryStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceRetryStrategyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceRetryStrategyResponseBody) *UpdateInstanceRetryStrategyResponse
	GetBody() *UpdateInstanceRetryStrategyResponseBody
}

type UpdateInstanceRetryStrategyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceRetryStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceRetryStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRetryStrategyResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRetryStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceRetryStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceRetryStrategyResponse) GetBody() *UpdateInstanceRetryStrategyResponseBody {
	return s.Body
}

func (s *UpdateInstanceRetryStrategyResponse) SetHeaders(v map[string]*string) *UpdateInstanceRetryStrategyResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceRetryStrategyResponse) SetStatusCode(v int32) *UpdateInstanceRetryStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceRetryStrategyResponse) SetBody(v *UpdateInstanceRetryStrategyResponseBody) *UpdateInstanceRetryStrategyResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceRetryStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
