// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateForwardStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateForwardStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateForwardStrategyResponse
	GetStatusCode() *int32
	SetBody(v *CreateForwardStrategyResponseBody) *CreateForwardStrategyResponse
	GetBody() *CreateForwardStrategyResponseBody
}

type CreateForwardStrategyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateForwardStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateForwardStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardStrategyResponse) GoString() string {
	return s.String()
}

func (s *CreateForwardStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateForwardStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateForwardStrategyResponse) GetBody() *CreateForwardStrategyResponseBody {
	return s.Body
}

func (s *CreateForwardStrategyResponse) SetHeaders(v map[string]*string) *CreateForwardStrategyResponse {
	s.Headers = v
	return s
}

func (s *CreateForwardStrategyResponse) SetStatusCode(v int32) *CreateForwardStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateForwardStrategyResponse) SetBody(v *CreateForwardStrategyResponseBody) *CreateForwardStrategyResponse {
	s.Body = v
	return s
}

func (s *CreateForwardStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
