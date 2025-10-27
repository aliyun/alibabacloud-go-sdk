// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOpaStrategyNewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOpaStrategyNewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOpaStrategyNewResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOpaStrategyNewResponseBody) *UpdateOpaStrategyNewResponse
	GetBody() *UpdateOpaStrategyNewResponseBody
}

type UpdateOpaStrategyNewResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOpaStrategyNewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOpaStrategyNewResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpaStrategyNewResponse) GoString() string {
	return s.String()
}

func (s *UpdateOpaStrategyNewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOpaStrategyNewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOpaStrategyNewResponse) GetBody() *UpdateOpaStrategyNewResponseBody {
	return s.Body
}

func (s *UpdateOpaStrategyNewResponse) SetHeaders(v map[string]*string) *UpdateOpaStrategyNewResponse {
	s.Headers = v
	return s
}

func (s *UpdateOpaStrategyNewResponse) SetStatusCode(v int32) *UpdateOpaStrategyNewResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOpaStrategyNewResponse) SetBody(v *UpdateOpaStrategyNewResponseBody) *UpdateOpaStrategyNewResponse {
	s.Body = v
	return s
}

func (s *UpdateOpaStrategyNewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
