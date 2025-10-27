// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpaStrategyNewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOpaStrategyNewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOpaStrategyNewResponse
	GetStatusCode() *int32
	SetBody(v *CreateOpaStrategyNewResponseBody) *CreateOpaStrategyNewResponse
	GetBody() *CreateOpaStrategyNewResponseBody
}

type CreateOpaStrategyNewResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOpaStrategyNewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOpaStrategyNewResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOpaStrategyNewResponse) GoString() string {
	return s.String()
}

func (s *CreateOpaStrategyNewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOpaStrategyNewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOpaStrategyNewResponse) GetBody() *CreateOpaStrategyNewResponseBody {
	return s.Body
}

func (s *CreateOpaStrategyNewResponse) SetHeaders(v map[string]*string) *CreateOpaStrategyNewResponse {
	s.Headers = v
	return s
}

func (s *CreateOpaStrategyNewResponse) SetStatusCode(v int32) *CreateOpaStrategyNewResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOpaStrategyNewResponse) SetBody(v *CreateOpaStrategyNewResponseBody) *CreateOpaStrategyNewResponse {
	s.Body = v
	return s
}

func (s *CreateOpaStrategyNewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
