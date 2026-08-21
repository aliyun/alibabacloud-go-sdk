// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetForwardStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetForwardStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetForwardStrategyResponse
	GetStatusCode() *int32
	SetBody(v *GetForwardStrategyResponseBody) *GetForwardStrategyResponse
	GetBody() *GetForwardStrategyResponseBody
}

type GetForwardStrategyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetForwardStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetForwardStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetForwardStrategyResponse) GoString() string {
	return s.String()
}

func (s *GetForwardStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetForwardStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetForwardStrategyResponse) GetBody() *GetForwardStrategyResponseBody {
	return s.Body
}

func (s *GetForwardStrategyResponse) SetHeaders(v map[string]*string) *GetForwardStrategyResponse {
	s.Headers = v
	return s
}

func (s *GetForwardStrategyResponse) SetStatusCode(v int32) *GetForwardStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetForwardStrategyResponse) SetBody(v *GetForwardStrategyResponseBody) *GetForwardStrategyResponse {
	s.Body = v
	return s
}

func (s *GetForwardStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
